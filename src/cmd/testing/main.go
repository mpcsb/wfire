package main

import (
	"fmt"
	"simulation/shared"
	"simulation/terrain"
)

func main() {
	p1 := shared.Coord{Lat: 38.609054, Lon:-9.52466, Alt: 0.0}
	p2 := shared.Coord{Lat: 38.906044, Lon: -9.286461, Alt: 0.0}  
 
	t := terrain.GenerateTerrain(p1, p2, 30)
	fmt.Println(t.Width)
	terrain.GenerateObjets(t, "highway")
	terrain.GenerateObjets(t, "water")
 
}
