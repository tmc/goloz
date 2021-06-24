// +build !js

package main

import (
	"crypto/tls"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func dialRemoteServer(cfg RunConfig) (*grpc.ClientConn, error) {
	fmt.Println("dialing", cfg.ServerAddr)
	dialOpts := []grpc.DialOption{
		grpc.WithBlock(),
	}
	if cfg.Insecure {
		dialOpts = append(dialOpts, grpc.WithInsecure())
	} else {
		creds := credentials.NewTLS(&tls.Config{
			InsecureSkipVerify: true,
		})
		dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds))
	}

	return grpc.Dial(cfg.ServerAddr, dialOpts...)
}
