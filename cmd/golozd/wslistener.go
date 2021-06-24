package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"

	"google.golang.org/grpc"
	"nhooyr.io/websocket"
)

type wsListener struct {
	stop chan struct{}
	errc chan error
	conn chan net.Conn
	h    *http.Server

	grpcServer *grpc.Server
}

func ListenWS(lis net.Listener, grpcServer *grpc.Server) (net.Listener, error) {
	srv := wsListener{
		stop: make(chan struct{}),
		errc: make(chan error, 1),
		conn: make(chan net.Conn),

		grpcServer: grpcServer,
	}
	// TODO: support HTTPS
	srv.h = &http.Server{
		Handler: srv,
	}

	go func() {
		defer close(srv.errc)
		srv.errc <- srv.h.Serve(lis)
	}()
	return srv, nil
}

func (w wsListener) ServeHTTP(wr http.ResponseWriter, r *http.Request) {
	fmt.Println("http")
	if r.ProtoMajor == 2 && !strings.HasPrefix(
		r.Header.Get("Connection"), "Upgrade") {
		fmt.Println("grpc")
		w.grpcServer.ServeHTTP(wr, r)
		return
	} else {
		fmt.Println("ws")
		w.ServeWebsocket(wr, r)
	}
}

func (w wsListener) ServeWebsocket(wr http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(wr, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})

	if err != nil {
		log.Println(err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "fail")

	ctx := r.Context()
	select {
	case <-w.stop:
		return
	default:
		w.conn <- websocket.NetConn(ctx, c, websocket.MessageBinary)
		select {
		// wait until wsListener is closed or when request is over
		case <-w.stop:
		case <-r.Context().Done():
		}
	}
	c.Close(websocket.StatusNormalClosure, "ok")
}

func (w wsListener) Accept() (net.Conn, error) {
	select {
	case <-w.stop:
		return nil, fmt.Errorf("server stopped")
	case err := <-w.errc:
		_ = w.Close()
		return nil, err
	case c := <-w.conn:
		return c, nil
	}
}

func (w wsListener) Close() error {
	select {
	case <-w.stop:
	default:
		close(w.stop)
	}
	if w.h != nil {
		return w.h.Close()
	}

	return nil
}

func (w wsListener) Addr() net.Addr {
	return net.Addr(nil)
}
