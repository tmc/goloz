package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmc/goloz"
	pb "github.com/tmc/goloz/proto/goloz/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	var (
		flagConnect = flag.String("connect", "golozd-1.tmc.dev:443", "server address")
		// var flagConnectWs = flag.String("connect-ws", "ws://golozd-1.tmc.dev:443/ws", "server address")
		flagUserName  = flag.String("username", "", "username")
		flagInsecure  = flag.Bool("insecure", false, "if specified, allow insecure traffic")
		flagLocalOnly = flag.Bool("local", false, "if true, only run in local mode")
		flagWindowIdx = flag.Int("w", 0, "if specified, picks tiled window position")
		flagMuted     = flag.Bool("muted", false, "if true, mutes audio output")
	)
	flag.Parse()

	runClient(RunConfig{
		ServerAddr: *flagConnect,
		Insecure:   *flagInsecure,
		LocalOnly:  *flagLocalOnly,

		WindowIdx: *flagWindowIdx,
	}, goloz.Settings{
		UserIdentity: resolveUserIdentity(*flagUserName),
		AudioMuted:   *flagMuted,
	})
}

func runClient(cfg RunConfig, settings goloz.Settings) {
	w, h := 640, 480
	ebiten.SetWindowSize(w, h)
	ebiten.SetWindowTitle("goloz")
	ebiten.SetInitFocused(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowPosition(50, cfg.WindowIdx*h)

	ctx := context.Background()

	var client pb.GameServerServiceClient
	// If in remote mode, create a connection to the server.
	if !cfg.LocalOnly {
		conn, err := dialRemoteServer(cfg)
		if err != nil {
			log.Fatal(err)
		}
		if conn != nil {
			defer func() {
				conn.Close()
				fmt.Println("calling conn close")
			}()
		}
		client = pb.NewGameServerServiceClient(conn)
	}
	// Create the Game.
	g, err := goloz.NewGame(ctx, settings, client)
	if err != nil {
		log.Fatal(err)
	}

	if !cfg.LocalOnly {
		go g.RunNetworkSync(ctx)
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
	rand.Seed(time.Now().Unix())
	if pid == -1 {
		pid = rand.Intn(1000)
	}
	return fmt.Sprintf("%v:%v", hostname, pid)
}

func establishServerSync(ctx context.Context, settings goloz.Settings, conn *grpc.ClientConn) (pb.GameServerService_SyncClient, error) {
	fmt.Println("syncing as", settings.UserIdentity)
	client := pb.NewGameServerServiceClient(conn)
	ctx = metadata.AppendToOutgoingContext(ctx,
		"id", settings.UserIdentity,
	)
	return client.Sync(ctx)
}
