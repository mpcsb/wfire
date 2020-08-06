package main

import (
	"fmt"
	"simulation/forest"
	"simulation/shared"
)

func main() {
	p1 := shared.Coord{Lat: 38.793613, Lon: -9.453429, Alt: 0.0}
	p2 := shared.Coord{Lat: 38.813257, Lon: -9.434138, Alt: 0.0}

	// t := terrain.GenerateTerrain(p1, p2, 500)
	// terrain.CallPythonScripts(p1, p2, 0, "structures")

	samples := 10000
	dist := 1.0
	f := forest.ForestGeneration(p1, p2, samples, dist)
	if false {
		fmt.Println(f)
	}

}
