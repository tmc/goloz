// +build tools

package goloz

import (
	_ "github.com/bufbuild/buf/cmd/buf"
	_ "github.com/evanw/esbuild/cmd/esbuild"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
