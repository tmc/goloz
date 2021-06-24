package main

import (
	"context"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
	"nhooyr.io/websocket"
)

func dialer(s string, dt time.Duration) (net.Conn, error) {
	ctx := context.Background()
	wsConn, _, err := websocket.Dial(ctx, s, &websocket.DialOptions{})
	return websocket.NetConn(ctx, wsConn, websocket.MessageBinary), err
}

func dialRemoteServer(cfg RunConfig) (*grpc.ClientConn, error) {
	addr := "ws://" + cfg.ServerAddr + "/ws"
	fmt.Println("dialing", addr)
	dialOpts := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithDialer(dialer),
	}
	dialOpts = append(dialOpts, grpc.WithInsecure())
	return grpc.Dial(addr, dialOpts...)
}
