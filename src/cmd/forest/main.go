package main

import (
	"simulation/forest"
	"simulation/shared"
)

func main() {
	p1 := shared.Coord{Lat: 38.908577, Lon: -9.400535, Alt: 0.0} //38.914104, -9.400535
	p2 := shared.Coord{Lat: 38.914104, Lon: -9.384442, Alt: 0.0} //38.908577, -9.384442

	f := forest.ForestGeneration(p1, p2, 100)

	for i := 0; i < 36; i++ {
		f.RecordFrame()
	}
	f.Plot_forest()
}
