package forest

import (
	"simulation/fuel"
	"simulation/shared"
	"simulation/terrain"

	"gonum.org/v1/gonum/spatial/vptree"
)

type forest struct {
	tree_lst []fuel.Tree_data
	coords []vptree.Comparable{} 
}

func loadTreeDimensions() map[string][5]float64 {
	var treeDimensions map[string][5]float64
	treeDimensions["pine"] = [5]float64{12.0, 1.0, 0.03, 3.5, 0.75}
	return treeDimensions
}

type Coord_label2 struct {
	Lat, Lon float64
	Label     string
}

func (p Coord_label2) Distance(c vptree.Comparable) float64 {
	q := c.(Coord_label2)
	return shared.Haversine(p.Lat, p.Lon, q.Lat, q.Lon)
}




func ForestGeneration(p1, p2 shared.Coord) forest {
	terrain_points := terrain.GenerateTerrain(p1, p2, 200)
	treeDimensions := loadTreeDimensions()

	// needs to be fractioned in smaller areas
	latitudes := shared.Linspace(p1.Lat, p2.Lat, t.Width)
	longitudes := shared.Linspace(p1.Lat, p2.Lat, t.Length) 

	for lat in range latitudes{
		for lon in range longitudes{
			lat := shared.Binterp()
			t := fuel.CreateTree(1, shared.Coord{Lat:lat, Loon:lon, Alt:alt}, "pine", treeDimensions)
			f.tree_lst = append(f.tree_lst, t)
			coord_type := terrain_points.Coord_Type[i]
			label := coord_type.Label
		}
	}



	
	var f forest
	for i := range terrain_points.Coord_Type {
		coord_type := terrain_points.Coord_Type[i]
		label := coord_type.Label
		if label == "tree"{
			t := fuel.CreateTree(1, coord_type.Coord, "pine", treeDimensions)
			f.tree_lst = append(f.tree_lst, t)
		}
	}
	return f
}
