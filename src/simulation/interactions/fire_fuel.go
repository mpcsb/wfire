package interactions

import (
	"math/rand"
	"simulation/fire"

	"simulation/fuel"
	s "simulation/shared"
	"simulation/weather"
)

// TreeImport fuel.Tree_data
type TreeImport fuel.Tree_data

// Burn starts a fire in a tree: this is a 1m flame
func (t *TreeImport) Burn(w weather.Wind) {
	t.Dynamic.Flame = fire.Flame{Coord: t.Coords, Height: 1.0, Direction: w.Direction, Radius: 1.0, Temperature: 800.0}
	t.Dynamic.State = "burning"
}

// Burning emits embers and burns adjacent trees
func (t *TreeImport) Burning(w weather.Wind) (emberPosition s.Coord) {
	if rand.Float64() > 0.05 {
		emberPosition = t.emitEmber(w)
	}
	return emberPosition
}

// EmitEmber is called when a tree is burning
func (t *TreeImport) emitEmber(w weather.Wind) s.Coord {
	rand.Seed(1999)
	stdDev := w.Speed * 25
	mean := 0.0
	// sampling from Half-Normal dist
	dx := s.Abs(rand.NormFloat64()*stdDev + mean)
	dy := s.Abs(rand.NormFloat64()*stdDev + mean)
	landingPoint := s.NewCoordXY(t.Coords, dx, dy)
	return landingPoint
	// newLocation := TreeCoord{ID: 0, Lat: newPoint.Lat, Lon: newPoint.Lon, Alt: t.Coords.Alt}

	// var keep vptree.Keeper
	// keep = vptree.NewDistKeeper(2.0)
	// if len(keep) > 0 {
	// 	VP.NearestSet(keep, newLocation)
	// 	for _, neighbour_tree := range keep.(*vptree.DistKeeper).Heap {
	// 		tree := neighbour_tree.Comparable.(TreeCoord)

	// 		treeToBurn := Tree_lst[tree.ID]
	// 		treeToBurn.Burn(w)
	// 	}
	// }
}
