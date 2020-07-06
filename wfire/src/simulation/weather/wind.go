package weather

import "simulation/shared"

type wind struct {
	direction float64
	speed     float64
	coords    shared.Coord
}
