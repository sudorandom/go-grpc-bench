package main

import (
	"context"
	"flag"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"

	vtgrpc "github.com/planetscale/vtprotobuf/codec/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"

	pb "github.com/sudorandom/go-grpc-bench/gen"
)

var _ pb.FlexServiceServer = (*server)(nil)

type server struct {
	pb.UnimplementedFlexServiceServer
}

// NormalRPC implements gen.FlexServiceServer.
func (s *server) NormalRPC(ctx context.Context, req *pb.FlexRequest) (*pb.FlexReply, error) {
	return &pb.FlexReply{Msg: req.Msg}, nil
}

func main() {
	var withVT bool
	flag.BoolVar(&withVT, "withvt", false, "enable vtprotobuf optimizations")
	flag.Parse()

	log.Printf("Starting grpc-go withVT=%t", withVT)

	if withVT {
		encoding.RegisterCodec(vtgrpc.Codec{})
	}

	go func() {
		log.Println(http.ListenAndServe("localhost:6669", nil))
	}()

	lis, err := net.Listen("tcp", ":6660")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterFlexServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("listen failed: %s", err)
	}
}
