package forest

import (
	"fmt"

	"simulation/fuel"
	"simulation/shared"
	"simulation/terrain"
	"math/rand"

	"gonum.org/v1/gonum/spatial/vptree"
)

type Forest struct {
	Tree_lst []fuel.Tree_data
	Coords []vptree.Comparable
}

type Coord_label2 struct {
	Lat, Lon float64
	Label     string
}

func loadTreeDimensions() map[string][5]float64 {
	treeDimensions := make(map[string][5]float64)
	treeDimensions["pine"] = [5]float64{12.0, 1.0, 0.03, 3.5, 0.75}
	treeDimensions["eucalyptus"] = [5]float64{12.0, 1.0, 0.01, 3.5, 0.75}
	return treeDimensions
}



func (p Coord_label2) Distance(c vptree.Comparable) float64 {
	q := c.(Coord_label2)
	return shared.Haversine(p.Lat, p.Lon, q.Lat, q.Lon)
}




func ForestGeneration(p1, p2 shared.Coord, samples int) Forest {
	rand.Seed(1999) 
	t := terrain.GenerateTerrain(p1, p2, samples)
	fmt.Println("Terrain completed")
	treeDimensions := loadTreeDimensions()  

	// We can divide the range of latitudes or coordinates by width and length
	// Width and length are the northing and easting which is equivalent
	t.Length = shared.Haversine(p1.Lat, p1.Lon, p1.Lat, p2.Lon)
	t.Width = shared.Haversine(p1.Lat, p2.Lon, p2.Lat, p2.Lon)
	latitudes := shared.Linspace(p1.Lat, p2.Lat, t.Width)
	longitudes := shared.Linspace(p1.Lon, p2.Lon, t.Length) 
	// fmt.Println(t.Width, t.Length)

	var f Forest
	// fmt.Println(shared.Haversine(p1.Lat, p1.Lon, p1.Lat, p2.Lon), shared.Haversine(p1.Lat, p2.Lon, p2.Lat, p2.Lon))
	// fmt.Println(len(longitudes), len(latitudes))
	for _, lat := range latitudes{
		for _, lon := range longitudes{ 
			if rand.Float64() > 0.01 {
				continue
			} else{ 
				// binterp interpolates altitude given lat and lon
				alt := t.Binterp(shared.Coord{Lat:lat, Lon:lon, Alt:0.0})
				tree := fuel.CreateTree(1, shared.Coord{Lat:lat, Lon:lon, Alt:alt}, "pine", treeDimensions)
				f.Tree_lst = append(f.Tree_lst, tree)
				
				// label := coord_type.Label
 
			}
		}
		// fmt.Println(len(f.Tree_lst ))
	} 
	return f
}
