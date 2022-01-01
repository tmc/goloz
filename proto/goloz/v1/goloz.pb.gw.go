// Code generated by protoc-gen-grpc-gateway. DO NOT EDIT.
// source: proto/goloz/v1/goloz.proto

/*
Package v1 is a reverse proxy.

It translates gRPC into RESTful JSON APIs.
*/
package v1

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/grpc-ecosystem/grpc-gateway/v2/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// Suppress "imported and not used" errors
var _ codes.Code
var _ io.Reader
var _ status.Status
var _ = runtime.String
var _ = utilities.NewDoubleArray
var _ = metadata.Join

func request_GameServerService_Sync_0(ctx context.Context, marshaler runtime.Marshaler, client GameServerServiceClient, req *http.Request, pathParams map[string]string) (GameServerService_SyncClient, runtime.ServerMetadata, error) {
	var metadata runtime.ServerMetadata
	stream, err := client.Sync(ctx)
	if err != nil {
		grpclog.Infof("Failed to start streaming: %v", err)
		return nil, metadata, err
	}
	dec := marshaler.NewDecoder(req.Body)
	handleSend := func() error {
		var protoReq SyncRequest
		err := dec.Decode(&protoReq)
		if err == io.EOF {
			return err
		}
		if err != nil {
			grpclog.Infof("Failed to decode request: %v", err)
			return err
		}
		if err := stream.Send(&protoReq); err != nil {
			grpclog.Infof("Failed to send request: %v", err)
			return err
		}
		return nil
	}
	go func() {
		for {
			if err := handleSend(); err != nil {
				break
			}
		}
		if err := stream.CloseSend(); err != nil {
			grpclog.Infof("Failed to terminate client stream: %v", err)
		}
	}()
	header, err := stream.Header()
	if err != nil {
		grpclog.Infof("Failed to get header from client: %v", err)
		return nil, metadata, err
	}
	metadata.HeaderMD = header
	return stream, metadata, nil
}

// RegisterGameServerServiceHandlerServer registers the http handlers for service GameServerService to "mux".
// UnaryRPC     :call GameServerServiceServer directly.
// StreamingRPC :currently unsupported pending https://github.com/grpc/grpc-go/issues/906.
// Note that using this registration option will cause many gRPC library features to stop working. Consider using RegisterGameServerServiceHandlerFromEndpoint instead.
func RegisterGameServerServiceHandlerServer(ctx context.Context, mux *runtime.ServeMux, server GameServerServiceServer) error {

	mux.Handle("POST", pattern_GameServerService_Sync_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		err := status.Error(codes.Unimplemented, "streaming calls are not yet supported in the in-process transport")
		_, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
		return
	})

	return nil
}

// RegisterGameServerServiceHandlerFromEndpoint is same as RegisterGameServerServiceHandler but
// automatically dials to "endpoint" and closes the connection when "ctx" gets done.
func RegisterGameServerServiceHandlerFromEndpoint(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error) {
	conn, err := grpc.Dial(endpoint, opts...)
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				grpclog.Infof("Failed to close conn to %s: %v", endpoint, cerr)
			}
		}()
	}()

	return RegisterGameServerServiceHandler(ctx, mux, conn)
}

// RegisterGameServerServiceHandler registers the http handlers for service GameServerService to "mux".
// The handlers forward requests to the grpc endpoint over "conn".
func RegisterGameServerServiceHandler(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error {
	return RegisterGameServerServiceHandlerClient(ctx, mux, NewGameServerServiceClient(conn))
}

// RegisterGameServerServiceHandlerClient registers the http handlers for service GameServerService
// to "mux". The handlers forward requests to the grpc endpoint over the given implementation of "GameServerServiceClient".
// Note: the gRPC framework executes interceptors within the gRPC handler. If the passed in "GameServerServiceClient"
// doesn't go through the normal gRPC flow (creating a gRPC client etc.) then it will be up to the passed in
// "GameServerServiceClient" to call the correct interceptors.
func RegisterGameServerServiceHandlerClient(ctx context.Context, mux *runtime.ServeMux, client GameServerServiceClient) error {

	mux.Handle("POST", pattern_GameServerService_Sync_0, func(w http.ResponseWriter, req *http.Request, pathParams map[string]string) {
		ctx, cancel := context.WithCancel(req.Context())
		defer cancel()
		inboundMarshaler, outboundMarshaler := runtime.MarshalerForRequest(mux, req)
		rctx, err := runtime.AnnotateContext(ctx, mux, req, "/goloz.v1.GameServerService/Sync", runtime.WithHTTPPathPattern("/goloz.v1.GameServerService/Sync"))
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}
		resp, md, err := request_GameServerService_Sync_0(rctx, inboundMarshaler, client, req, pathParams)
		ctx = runtime.NewServerMetadataContext(ctx, md)
		if err != nil {
			runtime.HTTPError(ctx, mux, outboundMarshaler, w, req, err)
			return
		}

		forward_GameServerService_Sync_0(ctx, mux, outboundMarshaler, w, req, func() (proto.Message, error) { return resp.Recv() }, mux.GetForwardResponseOptions()...)

	})

	return nil
}

var (
	pattern_GameServerService_Sync_0 = runtime.MustPattern(runtime.NewPattern(1, []int{2, 0, 2, 1}, []string{"goloz.v1.GameServerService", "Sync"}, ""))
)

var (
	forward_GameServerService_Sync_0 = runtime.ForwardResponseStream
)
