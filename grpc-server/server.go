package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/yuichiro12/a/graph"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var port = "127.0.0.1:5000"

type server struct{}

func (s server) Save(ctx context.Context, req *pb.SaveRequest) (*pb.SaveResponse, error) {
	time.Sleep(1 * time.Second)
	b, err := json.MarshalIndent(req.Digraph, "", "  ")
	if err != nil {
		return new(pb.SaveResponse), err
	}
	fmt.Printf("%v\n", string(b))
	return &pb.SaveResponse{
		IsSuccess: true,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	srv := new(server)
	pb.RegisterDirectedWeightedMultiGraphServiceServer(s, srv)
	log.Fatal(s.Serve(lis))
}
