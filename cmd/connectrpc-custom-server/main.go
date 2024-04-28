package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	_ "net/http/pprof"

	"connectrpc.com/connect"
	"github.com/sudorandom/go-grpc-bench/gen"
	"github.com/sudorandom/go-grpc-bench/gen/genconnect"
	"go.akshayshah.org/connectproto"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

var _ genconnect.FlexServiceHandler = (*server)(nil)

type server struct {
	genconnect.UnimplementedFlexServiceHandler
}

// NormalRPC implements genconnect.FlexServiceHandler.
func (s *server) NormalRPC(ctx context.Context, req *connect.Request[gen.FlexRequest]) (*connect.Response[gen.FlexReply], error) {
	return connect.NewResponse(&gen.FlexReply{Msg: req.Msg.Msg}), nil
}

func main() {
	var withVT bool
	flag.BoolVar(&withVT, "withvt", false, "enable vtprotobuf optimizations")
	flag.Parse()

	log.Printf("Starting connectrpc withVT=%t", withVT)

	go func() {
		log.Println(http.ListenAndServe("localhost:6669", nil))
	}()

	mux := http.NewServeMux()
	opts := []connect.HandlerOption{}
	if withVT {
		opts = append(opts, connectproto.WithBinaryVT())
	}
	mux.Handle(genconnect.NewFlexServiceHandler(&server{}, opts...))

	srv := http.Server{
		Addr:    ":6660",
		Handler: h2c.NewHandler(mux, &http2.Server{}),
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("listen failed: %s", err)
	}
}
