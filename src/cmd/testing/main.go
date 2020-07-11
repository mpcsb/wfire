package main

import (
	"fmt"
	"math/rand"

	"simulation/shared"
	"simulation/terrain"
)

func main() {
	p1 := shared.Coord{Lat: 38.123, Lon: 12.755, Alt: 0.0}
	p2 := shared.Coord{Lat: 38.245, Lon: 12.889, Alt: 0.0}

	t := terrain.GenerateTerrain(p1, p2)
	fmt.Println(t.Length, t.Width)

	r := rand.New(rand.NewSource(99))
	u := 1
	for u < 10 {
		w
		u += 1
		fmt.Println(r.Float64())
	}
}
