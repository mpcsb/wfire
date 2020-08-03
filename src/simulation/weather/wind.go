package weather

import (
	s "simulation/shared"
	"simulation/terrain"
)

// Wind contains coordinates and information regarding speed and direction and also temperature
type Wind struct {
	Coord       s.Coord
	Direction   float64
	Speed       float64
	Temperature float64
}

// func (w *Wind) UpdateAdjacentWind() {

// }

// WindGeneration should have some minimum threshold on the amount of detail. It's not realistic to have more than 100 m of detail
// this will be replaced by a ML model that ingests WindNinja models and returns wind maps
func WindGeneration(t terrain.Terrain) map[s.Coord]Wind {
	var wm map[s.Coord]Wind
	for _, v := range t.Coord_Type {
		w := Wind{Coord: v.Coord, Direction: 35.0, Speed: 3.0, Temperature: 25.0}
		wm[v.Coord] = w
	}
	return wm
}
