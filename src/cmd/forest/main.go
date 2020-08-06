package main

import (
	"log"
	"simulation/forest"
	"simulation/shared"
	"time"
)

func main() {
	p1 := shared.Coord{Lat: 38.848577, Lon: -9.410535, Alt: 0.0} //38.944104, -9.400535
	p2 := shared.Coord{Lat: 38.924104, Lon: -9.354442, Alt: 0.0} //38.908577, -9.364442

	f := forest.ForestGeneration(p1, p2, 100, 1.0)
	t0 := time.Now()
	f.GetNeighbours(1.3)
	t1 := time.Since(t0)
	log.Printf("serial took %s", t1)

	t2 := time.Now()
	f.DistributedNeighbour(1.3)

	for i := 0; i < 36; i++ {
		f.RecordFrame()
	}

	f.Plot_forest()
	t3 := time.Since(t2)
	log.Printf("parallel took %s", t3)
}
