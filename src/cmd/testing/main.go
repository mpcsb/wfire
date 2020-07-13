package main

import (
	"fmt" 

	"simulation/shared"
	"simulation/terrain"
)

func main() {
	p1 := shared.Coord{Lat: 38.709054, Lon:-9.482466, Alt: 0.0}
	p2 := shared.Coord{Lat: 38.806044, Lon: -9.386461, Alt: 0.0}  

	// fmt.Println(p1.Lat, p2.Lon)
	t := terrain.GenerateTerrain(p1, p2)
	fmt.Println(t.Length, t.Width)
 
}
