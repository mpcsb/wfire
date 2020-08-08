package main

import (
	"fmt"
	"simulation/forest"
	"simulation/shared"
	"time"
)

func main() {
	// p1 := shared.Coord{Lat: 38.848577, Lon: -9.410535, Alt: 0.0}  //38.944104, -9.400535
	// p2 := shared.Coord{Lat: 38.924104, Lon: -9.354442, Alt: 0.0} //38.908577, -9.364442
	p1 := shared.Coord{Lat: 38.848577, Lon: -9.410535, Alt: 0.0} //38.944104, -9.400535
	p2 := shared.Coord{Lat: 38.854104, Lon: -9.404442, Alt: 0.0} //38.908577, -9.364442

	f := forest.Generation(p1, p2, 100, 10)

	// fmt.Println("Number of trees:", len(f.Tree_lst))
	// t0 := time.Now()
	// f.GetNeighbours(4.0)
	// t1 := time.Since(t0)
	// fmt.Println("serial took", t1)
	// ntree := f.Tree_lst[100].Neighbours
	// fmt.Println("tree 100 adjacent trees:", ntree)

	t2 := time.Now()
	f.DistributedGetNeighbours(4.0)
	t26 := time.Since(t2)
	fmt.Println("parallel took", t26)
	ntree := f.Tree_lst[100].Neighbours
	ntree1 := f.Tree_lst[104].Neighbours

	fmt.Println("tree 100 adjacent trees:", ntree, ntree1)

	for i := 0; i < 36; i++ {
		f.RecordFrame()
	}

	f.Plot_forest()
	t3 := time.Since(t2)
	fmt.Println("total tasks took", t3)
}
