package flora

import "simulation/terrain"

// type Coord struct {
// 	lat, lon, alt float64
// }

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
	coords  terrain.Coord
	x_y     xy_coor
	static  tree_static
	dynamic tree_dynamic

	north_facing string
	neighbours   []int
}

//var treeDimensions map[string][]float64

func CreateTree(id int, p terrain.Coord, species string, tree_db map[string][5]float64) Tree_data {
	t := Tree_data{ID: id}
	t.species = species
	initStatic(t, tree_db)
	initBiomass(t)
	return t
}

func initStatic(t Tree_data, tree_db map[string][5]float64) {
	dims := tree_db[t.species]

	t.static.height = dims[0]
	t.static.diameter_breast_height = dims[1]
	t.static.bark_thickness = dims[2]
	t.static.crown_radius = dims[3]
	t.static.flammability = dims[4]
}

func initBiomass(t Tree_data) {
	t.dynamic.biomass = t.static.height * t.static.diameter_breast_height * t.static.crown_radius
}

func updateMoisture(t Tree_data, temp float64) {
	temp_diff := temp - 25.0 // 25ºC to be defined
	diff := 0.0
	if temp_diff > 0 {
		diff = 0.01 * temp_diff
	} else {
		diff = 0
	}
	t.dynamic.moisture = t.dynamic.moisture - diff
}

type tree interface {
	CreateTree()
	updateMoisture()
}
