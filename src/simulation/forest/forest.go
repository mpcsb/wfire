package forest

import (
	"math"

	"math/rand"
	"simulation/fuel"
	"simulation/shared"
	"simulation/terrain"

	"gonum.org/v1/gonum/spatial/vptree"
)

type Forest struct {
	Tree_lst    []fuel.Tree_data
	Tree_Coords []vptree.Comparable
	// VP           vptree.Tree
	Frame        int
	Sample_trees []int
}

type TreeCoord struct {
	ID  int
	Lat float64
	Lon float64
	Alt float64
}

func loadTreeDimensions() map[string][5]float64 {
	treeDimensions := make(map[string][5]float64)
	treeDimensions["pine"] = [5]float64{12.0, 1.0, 0.03, 3.5, 0.75}
	treeDimensions["eucalyptus"] = [5]float64{12.0, 1.0, 0.01, 3.5, 0.75}
	return treeDimensions
}

// Distance merge TreeCoord with fuel.tree_data
func (p TreeCoord) Distance(c vptree.Comparable) float64 {
	q := c.(TreeCoord)
	Dxy := shared.Haversine(p.Lat, p.Lon, q.Lat, q.Lon)
	Dz := q.Alt - p.Alt
	return math.Sqrt(math.Pow(Dxy, 2) + math.Pow(Dz, 2))
}

// Generation implements terrain and generates trees in a metered grid, forming a forest
func Generation(p1, p2 shared.Coord, samples int, dist float64) (f Forest) {
	rand.Seed(1999)
	t := terrain.GenerateTerrain(p1, p2, samples) // terrain should not be controlled by samples, but instead by SRTM resolution

	treeDimensions := loadTreeDimensions()

	// We can divide the range of latitudes or coordinates by width and length
	// Width and length are the northing and easting which is equivalent
	latitudes := shared.Linspace(p1.Lat, p2.Lat, t.Width)
	longitudes := shared.Linspace(p1.Lon, p2.Lon, t.Length)

	id := 0
	for i_lat, lat := range latitudes {
		if i_lat < 2 {
			continue
		}
		if i_lat > len(latitudes)-3 {
			continue
		}

		for i_lon, lon := range longitudes {
			if i_lon < 2 {
				continue
			}
			if i_lon > len(longitudes)-3 {
				continue
			}

			// forest density: 0.1 is somewhat light, but realistic
			if rand.Float64() > 0.1 {
				continue
			} else {
				// binterp interpolates alt, slope, aspect given lat and lon
				alt, _, _ := t.Binterp(shared.Coord{Lat: lat, Lon: lon, Alt: 0.0})
				if alt < t.MinHeight {
					alt = t.MinHeight
				}
				if alt > t.MaxHeight {
					alt = t.MaxHeight
				}
				tree := fuel.CreateTree(id, shared.Coord{Lat: lat, Lon: lon, Alt: alt}, "pine", treeDimensions)

				f.Tree_lst = append(f.Tree_lst, tree)

				new_tree := TreeCoord{ID: id, Lat: lat, Lon: lon, Alt: alt}
				f.Tree_Coords = append(f.Tree_Coords, new_tree)
				id += 1
			}
		}
	}

	f.DetermineSample(10000)

	f.RecordFrame()

	return f
}

func (f *Forest) DetermineSample(size int) {
	selected := make(map[int]bool, size)

	for tries := 0; len(selected) < size; tries++ {
		i := rand.Intn(len(f.Tree_lst))

		if selected[i] {
			continue
		} else {
			f.Sample_trees = append(f.Sample_trees, i)
			selected[i] = true
		}
	}

}
