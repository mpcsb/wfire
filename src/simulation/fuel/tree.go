package fuel

import "simulation/shared"

type xy_coor struct {
	x, y float64
}

type tree_static struct {
	height                 float64 //https://en.wikipedia.org/wiki/Tree_measurement#Tree_databases
	diameter_breast_height float64
	bark_thickness         float64
	crown_radius           float64
	flammability           float64 //https://fireandemergency.nz/at-home/flammability-of-plant-species/
}

type tree_dynamic struct {
	moisture float64
	state    string // unburnt, burning, ember, burnt
	biomass  float64
}

type Tree_data struct {
	ID      int
	species string
	Coords  shared.Coord
	x_y     xy_coor
	static  tree_static
	dynamic tree_dynamic

	north_facing string
	Neighbours   []int
}

func CreateTree(id int, p shared.Coord, species string, tree_db map[string][5]float64) Tree_data {
	t := Tree_data{ID: id}
	t.species = species
	t.InitStatic(tree_db)
	t.InitBiomass()
	return t
}

func (t Tree_data) InitStatic(tree_db map[string][5]float64) {
	dims := tree_db[t.species]

	t.static.height = dims[0]
	t.static.diameter_breast_height = dims[1]
	t.static.bark_thickness = dims[2]
	t.static.crown_radius = dims[3]
	t.static.flammability = dims[4]
}

func (t Tree_data) InitBiomass() {
	t.dynamic.biomass = t.static.height * t.static.diameter_breast_height * t.static.crown_radius
}

func (t Tree_data) UpdateMoisture(temperature float64) {
	temperature_diff := temperature - 25.0 // 25ºC to be defined
	diff := 0.0
	if temperature_diff > 0 {
		diff = 0.01 * temperature_diff
	} else {
		diff = 0
	}
	t.dynamic.moisture = t.dynamic.moisture - diff
}


// type tree interface {
// 	CreateTree()
// 	UpdateMoisture()
// }
