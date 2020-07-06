package main

import (
	"simulation/terrain"
)

func main() {
	p1 := terrain.Coord{Lat: 38.123, Lon: 12.455, Alt: 0.0}
	p2 := terrain.Coord{Lat: 38.145, Lon: 12.489, Alt: 0.0}

	terrain.ExtractCoordinates(p1, p2)
	// t := terrain.GenerateTerrain(p1, p2)
	// fmt.Println(t.Length, t.Width)

}
