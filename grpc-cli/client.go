package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"

	pb "github.com/yuichiro12/a/graph"
)

var serverAddr = "127.0.0.1:5000"

func main() {
	conn, err := grpc.Dial(serverAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	c := pb.NewDirectedWeightedMultiGraphServiceClient(conn)
	ctx := context.Background()
	for {
		req := &pb.SaveRequest{
			Digraph: generateGraph(),
		}
		resp, err := c.Save(ctx, req)
		if err != nil {
			log.Fatal(err)
		}
		if resp.IsSuccess {
			fmt.Println("success")
		}
	}
}

func generateGraph() *pb.DirectedWeightedMultiGraph {
	rand.Seed(time.Now().UnixNano())
	numVertices := rand.Intn(5) + 1
	var vList []*pb.Vertex
	for i := 0; i < numVertices; i++ {
		v := &pb.Vertex{
			Id:     int64(i),
			Weight: rand.Float64(),
		}
		vList = append(vList, v)
	}

	numArcs := rand.Intn(numVertices * numVertices)
	var aList []*pb.Arc
	for i := 0; i < numArcs; i++ {
		a := &pb.Arc{
			Id:     int64(i),
			Weight: rand.Float64(),
			From:   vList[rand.Intn(len(vList))],
			To:     vList[rand.Intn(len(vList))],
		}
		aList = append(aList, a)
	}
	return &pb.DirectedWeightedMultiGraph{
		Vertex: vList,
		Arc:    aList,
	}
}
