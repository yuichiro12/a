package main

import (
	"math/rand"
	"time"

	pb "github.com/yuichiro12/a/graph"
)

func main() {
	for {
		rand.Seed(time.Now().UnixNano())
		maxVertices := rand.Intn(50)
		var vList []*pb.Vertex
		for i := 0; i < maxVertices; i++ {
			v := &pb.Vertex{
				Id:     int64(i),
				Weight: rand.Float64(),
			}
			vList = append(vList, v)
		}

		maxArcs := rand.Intn(2500)
		var aList []*pb.Arc
		for i := 0; i < maxArcs; i++ {
			a := &pb.Arc{
				Id:     int64(i),
				Weight: rand.Float64(),
				From:   vList[rand.Intn(len(vList))],
				To:     vList[rand.Intn(len(vList))],
			}
			aList = append(aList, a)
		}

		d := &pb.DirectedWeightedMultiGraph{
			Vertex: vList,
			Arc:    aList,
		}

	}
}
