package weather

import s "simulation/shared"

type Wind struct {
	Direction float64
	Speed     float64
	CoordS    s.Coord
}
