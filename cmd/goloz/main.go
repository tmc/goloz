package main

import (
	"context"
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/tmc/goloz"
	pb "github.com/tmc/goloz/proto/goloz/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

func main() {
	var flagConnect = flag.String("connect", "goloz-gameserver-kblm3ew5ta-uc.a.run.app:443", "server address")
	var flagUserName = flag.String("username", "", "username")
	var flagInsecure = flag.Bool("insecure", false, "username")
	flag.Parse()

	runClient(*flagConnect, userIdentity(*flagUserName), *flagInsecure)
}

func runClient(serverAddr string, userIdentity string, insecure bool) {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("goloz")

	fmt.Println("dialing", serverAddr)

	dialOpts := []grpc.DialOption{
		grpc.WithBlock(),
	}
	if insecure {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	} else {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(serverAddr, dialOpts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	fmt.Println("connected")
	ctx := context.Background()
	client := pb.NewGameServerServiceClient(conn)
	ctx = metadata.AppendToOutgoingContext(ctx,
		"id", userIdentity,
	)
	fmt.Println("syncing as", userIdentity)
	syncClient, err := client.Sync(ctx)
	if err != nil {
		log.Println(err)
	}
	// Contact the server and print out its response.
	g := goloz.NewGame(ctx, syncClient)

	go g.RunNetworkSync(ctx, userIdentity)
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}

func userIdentity(explicitUsername string) string {
	if explicitUsername != "" {
		return explicitUsername
	}
	hostname, _ := os.Hostname()
	pid := os.Getpid()
	return fmt.Sprintf("%v:%v", hostname, pid)
}
