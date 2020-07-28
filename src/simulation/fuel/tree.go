package fuel

import "simulation/shared"

// type xy_coor struct {
// 	x, y float64
// }

type Tree_static struct {
	Height                 float64 //https://en.wikipedia.org/wiki/Tree_measurement#Tree_databases
	Diameter_breast_height float64
	Bark_thickness         float64
	Crown_radius           float64
	Flammability           float64 //https://fireandemergency.nz/at-home/flammability-of-plant-species/
}

type Tree_dynamic struct {
	Moisture float64
	State    string // tree, burning, ember, burnt
	Biomass  float64
}

type Tree_data struct {
	ID      int
	species string
	Coords  shared.Coord 
	Static  Tree_static
	Dynamic Tree_dynamic

	North_facing string
	Neighbours   []int
}

func CreateTree(id int, p shared.Coord, species string, tree_db map[string][5]float64) Tree_data {
	t := Tree_data{ID: id}
	t.species = species
	t.Coords = shared.Coord{Lat:p.Lat, Lon:p.Lon, Alt:p.Alt}

	t.InitStatic(tree_db)
	t.InitBiomass()
	return t
}

func (t *Tree_data) InitStatic(tree_db map[string][5]float64) {
	dims := tree_db[t.species]

	t.Static.Height = dims[0]
	t.Static.Diameter_breast_height = dims[1]
	t.Static.Bark_thickness = dims[2]
	t.Static.Crown_radius = dims[3]
	t.Static.Flammability = dims[4]
}

func (t *Tree_data) InitBiomass() {
	t.Dynamic.Biomass = t.Static.Height * t.Static.Diameter_breast_height * t.Static.Crown_radius
	t.Dynamic.State = "tree"
}

func (t *Tree_data) UpdateMoisture(temperature float64) {
	temperature_diff := temperature - 25.0 // 25ºC to be defined
	diff := 0.0
	if temperature_diff > 0 {
		diff = 0.01 * temperature_diff
	} else {
		diff = 0
	}
	t.Dynamic.Moisture = t.Dynamic.Moisture - diff
}


// type tree interface {
// 	CreateTree()
// 	UpdateMoisture()
// }
