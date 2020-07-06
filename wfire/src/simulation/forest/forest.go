package forest

import (
	"simulation/flora"
	"simulation/terrain"
)

type forest struct {
	tree_lst []flora.Tree_data
}

func ForestGeneration(p1 terrain.Coord, p2 terrain.Coord) forest {
	var treeDimensions map[string][5]float64
	treeDimensions["pine"] = [5]float64{12.0, 1.0, 0.03, 3.5, 0.75}

	terrain_points := terrain.GenerateTerrain(p1, p2)

	var f forest
	for i := range terrain_points.Coords {
		coordinate := terrain_points.Coords[i]
		t := flora.CreateTree(1, coordinate, "pine", treeDimensions)
		f.tree_lst = append(f.tree_lst, t)
	}
	return f
}
