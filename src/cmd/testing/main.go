package main

import (
	"simulation/shared"
	"simulation/terrain"
)

func main() {
	p1 := shared.Coord{Lat: 38.709054, Lon:-9.482466, Alt: 0.0}
	p2 := shared.Coord{Lat: 38.806044, Lon: -9.386461, Alt: 0.0}  
 
	t := terrain.GenerateTerrain(p1, p2)
	
	terrain.GenerateObjets(t, "highway")
 
}
