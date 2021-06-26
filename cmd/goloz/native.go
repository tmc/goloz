// +build !js

package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials"
)

func dialRemoteServer(cfg RunConfig) (*grpc.ClientConn, error) {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second)
	fmt.Println("dialing", cfg.ServerAddr)
	dialOpts := []grpc.DialOption{
		grpc.FailOnNonTempDialError(true),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.DefaultConfig,
		}),
	}
	if cfg.Insecure {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	} else {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))
	}

	return grpc.DialContext(ctx, cfg.ServerAddr, dialOpts...)
}
