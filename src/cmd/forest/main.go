package main

import (
	// "fmt"
	
	"simulation/shared"
	"simulation/forest"
)

func main() {
	p1 := shared.Coord{Lat: 38.773613, Lon: -9.363429, Alt: 0.0}
	p2 := shared.Coord{Lat: 38.813257, Lon: -9.324138, Alt: 0.0} 
 
	f := forest.ForestGeneration(p1, p2, 200)
	f.GetNeighbours(16.0)
}