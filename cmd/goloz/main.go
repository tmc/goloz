package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmc/goloz"
	pb "github.com/tmc/goloz/proto/goloz/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	var flagConnect = flag.String("connect", "golozd-1.tmc.dev:443", "server address")
	// var flagConnectWs = flag.String("connect-ws", "ws://golozd-1.tmc.dev:443/ws", "server address")
	var flagUserName = flag.String("username", "", "username")
	var flagInsecure = flag.Bool("insecure", false, "if specified, allow insecure traffic")
	var flagLocalOnly = flag.Bool("local", false, "if true, only run in local mode")
	flag.Parse()

	runClient(RunConfig{
		ServerAddr:   *flagConnect,
		UserIdentity: resolveUserIdentity(*flagUserName),
		Insecure:     *flagInsecure,
		LocalOnly:    *flagLocalOnly,
	})
}

func runClient(cfg RunConfig) {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("goloz")
	ebiten.SetInitFocused(false)
	ebiten.SetWindowPosition(0, 0)

	ctx := context.Background()
	var syncClient pb.GameServerService_SyncClient

	// If in remote mode, create a connection to the server.
	if !cfg.LocalOnly {
		conn, err := dialRemoteServer(cfg)
		if err != nil {
			log.Fatal(err)
		}
		if conn != nil {
			defer conn.Close()
			syncClient, err = establishServerSync(ctx, cfg, conn)
		}
	}
	// Create the Game.
	g, err := goloz.NewGame(ctx, syncClient)
	if err != nil {
		log.Fatal(err)
	}

	if !cfg.LocalOnly {
		go g.RunNetworkSync(ctx, cfg.UserIdentity)
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func resolveUserIdentity(explicitUsername string) string {
	if explicitUsername != "" {
		return explicitUsername
	}
	hostname, _ := os.Hostname()
	pid := os.Getpid()
	if pid == -1 {
		pid = rand.Intn(1000)
	}
	return fmt.Sprintf("%v:%v", hostname, pid)
}

func establishServerSync(ctx context.Context, cfg RunConfig, conn *grpc.ClientConn) (pb.GameServerService_SyncClient, error) {
	fmt.Println("syncing as", cfg.UserIdentity)
	client := pb.NewGameServerServiceClient(conn)
	ctx = metadata.AppendToOutgoingContext(ctx,
		"id", cfg.UserIdentity,
	)
	return client.Sync(ctx)
}
