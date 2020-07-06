package main

import (
	"fmt"
	"simulation/terrain"
)

func main() {
	p1 := terrain.Coord{Lat: 38.123, Lon: 12.455, Alt: 0.0}
	p2 := terrain.Coord{Lat: 38.245, Lon: 12.889, Alt: 0.0}

	t := terrain.GenerateTerrain(p1, p2)
	fmt.Println(t.Length, t.Width)

}
