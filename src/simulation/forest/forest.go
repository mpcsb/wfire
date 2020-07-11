package forest

import (
	"simulation/fuel"
	"simulation/shared"
	"simulation/terrain"
)

type forest struct {
	tree_lst []fuel.Tree_data
}

func loadTreeDimensions() map[string][5]float64 {
	var treeDimensions map[string][5]float64
	treeDimensions["pine"] = [5]float64{12.0, 1.0, 0.03, 3.5, 0.75}
	return treeDimensions
}

func ForestGeneration(p1 shared.Coord, p2 shared.Coord) forest {
	terrain_points := terrain.GenerateTerrain(p1, p2)
	treeDimensions := loadTreeDimensions()

	var f forest
	for i := range terrain_points.Coords {
		coordinate := terrain_points.Coords[i]
		t := fuel.CreateTree(1, coordinate, "pine", treeDimensions)
		f.tree_lst = append(f.tree_lst, t)
	}
	return f
}
