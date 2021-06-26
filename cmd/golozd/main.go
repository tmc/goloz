package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"

	"github.com/soheilhy/cmux"
	pb "github.com/tmc/goloz/proto/goloz/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	// port is the port that the server listens on.
	port = "50001"
)

func main() {
	flag.Parse()
	runServer()
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

func runServer() {
	if p := os.Getenv("PORT"); p != "" {
		port = p
	}
	fmt.Println("listening on :" + port)
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	ctx := context.Background()
	s := grpc.NewServer()
	srv, err := newServer()
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterGameServerServiceServer(s, srv)

	m := cmux.New(lis)
	grpcL := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpL := m.Match(cmux.Any())

	go func() {
		for {
			if err := srv.FanOutUpdates(ctx); err != nil {
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()
	go func() {
		l, err := ListenWS(httpL)
		if err != nil {
			panic(err)
		}
		if err := s.Serve(l); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	go func() {
		if err := s.Serve(grpcL); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	if err := m.Serve(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
