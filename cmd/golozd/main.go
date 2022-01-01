package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/soheilhy/cmux"
	"github.com/tmc/goloz/apidocs"
	pb "github.com/tmc/goloz/proto/goloz/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	// port is the port that the server listens on.
	flagListen = flag.String("listen", ":50001", "listen address")
)

func main() {
	flag.Parse()
	ctx := context.Background()
	runServer(ctx, *flagListen)
}

type characterUpdate struct {
	key   string
	value *pb.Character
}

type server struct {
	characters map[string]*pb.Character

	updates chan *characterUpdate

	// TODO: add mutex
	clients map[string]pb.GameServerService_SyncServer

	// Embed the unimplemented game server type.
	pb.UnimplementedGameServerServiceServer
}

func newServer() (*server, error) {
	return &server{
		clients: make(map[string]pb.GameServerService_SyncServer),
		updates: make(chan *characterUpdate),
	}, nil
}

func (server *server) FanOutUpdates(ctx context.Context) error {
	for {
		update := <-server.updates

		// TODO: fix race
		for _, c := range server.clients {
			err := c.Send(&pb.SyncResponse{
				Characters: map[string]*pb.Character{
					update.key: update.value,
				},
			})
			if err != nil {
				log.Println(err)
				// TODO: cleanup dead clients
			}
		}
	}
	// return nil
}

func (server *server) Sync(stream pb.GameServerService_SyncServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return fmt.Errorf("issue getting metadata context")
	}

	if len(md.Get("id")) != 1 {
		return fmt.Errorf("missing client id in metadata")
	}
	id := md.Get("id")[0]
	fmt.Println("got new connection:", id)
	// TODO fix race
	server.clients[id] = stream

	for {
		m, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		server.updates <- &characterUpdate{key: id, value: m.Character}
	}
	// TODO: defer cleanup
}

func runServer(ctx context.Context, listenAddr string) {
	if p := os.Getenv("PORT"); p != "" {
		listenAddr = ":" + p
	}
	fmt.Println("listening on", listenAddr)
	lis, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	srv, err := newServer()
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterGameServerServiceServer(grpcServer, srv)

	mux := cmux.New(lis)

	// Match connections in order:
	// First grpc, then HTTP, and otherwise Go RPC/TCP.
	httpListener := mux.Match(cmux.HTTP1Fast())
	grpcListener := mux.Match(cmux.Any())

	gwMux := runtime.NewServeMux()
	if err := pb.RegisterGameServerServiceHandlerServer(ctx, gwMux, srv); err != nil {
		log.Fatalf("failed to register with gateway handler: %v", err)
	}
	httpMux := http.NewServeMux()
	//httpMux.Handle("/", gwMux)
	//httpMux.Handle("/apidocs/", apidocs.Handler())
	httpMux.Handle("/apidocs/", http.StripPrefix("apidocs", apidocs.Handler()))
	httpMux.Handle("/", apidocs.Handler())
	httpServer := &http.Server{
		Handler: httpMux,
	}

	// Main game state distribution loop.
	go func() {
		for {
			if err := srv.FanOutUpdates(ctx); err != nil {
				if err != nil {
					log.Fatalf("failed to start state server: %v", err)
				}
			}
		}
	}()

	group := errgroup.Group{}
	group.Go(func() error { return httpServer.Serve(httpListener) })
	group.Go(func() error { return grpcServer.Serve(grpcListener) })
	group.Go(func() error { return mux.Serve() })
	if err := group.Wait(); err != nil {
		log.Fatal(err)
	}
}
