package fuel

import (
	"math"
	"simulation/fire"
	s "simulation/shared"
)

type Tree_static struct {
	Height               float64 //https://en.wikipedia.org/wiki/Tree_measurement#Tree_databases
	DiameterBreastHeight float64
	BarkThickness        float64
	CrownRadius          float64
	Flammability         float64 //https://fireandemergency.nz/at-home/flammability-of-plant-species/
	SparsenessFactor     float64
}

type Tree_dynamic struct {
	Moisture float64
	State    string // tree, burning, ember, burnt
	Trunk    float64
	Canopy   float64
	Flame    fire.Flame
}

// Tree_data holds static and dynamic information about each tree
type Tree_data struct {
	ID      int
	species string
	Coords  s.Coord
	Static  Tree_static
	Dynamic Tree_dynamic

	Aspect     string
	Neighbours []int
}

// CreateTree generates a tree based on tree_db details
func CreateTree(id int, p s.Coord, species string, treeDB map[string][5]float64) Tree_data {
	t := Tree_data{ID: id}
	t.species = species
	t.Coords = s.Coord{Lat: p.Lat, Lon: p.Lon, Alt: p.Alt}

	t.initStatic(treeDB)
	t.initBiomass()
	t.Dynamic.State = "tree"
	return t
}

func (t *Tree_data) initStatic(treeDB map[string][5]float64) {
	dims := treeDB[t.species]

	t.Static.Height = dims[0]
	t.Static.DiameterBreastHeight = dims[1]
	t.Static.BarkThickness = dims[2] * 100 // conversion to meter
	t.Static.CrownRadius = dims[3]
	t.Static.Flammability = dims[4]
	t.Static.SparsenessFactor = 0.2
	// t.Static.SparsenessFactor = dims[5]
}

// InitBiomass implements estimates of biomass belonging to:
// trunk: basic volumetric estimate
// Canopy: half of tree height * crown radius * how sparse the tree is
func (t *Tree_data) initBiomass() {
	t.Dynamic.Trunk = t.Static.Height * math.Pi * math.Pow(t.Static.DiameterBreastHeight, 2) / 4.0
	t.Dynamic.Canopy = (0.5 * t.Static.Height) * (math.Pi * math.Pow(t.Static.CrownRadius, 2) / 4.0) * t.Static.SparsenessFactor

}

// UpdateMoisture controls the water content depending on the temperature of the surrounding temperature
func (t *Tree_data) UpdateMoisture(temperature float64) {
	temperatureDiff := temperature - 25.0 // 25ºC. This would be an equilibirum point where no water transfers occur
	diff := 0.0
	if temperatureDiff > 0 {
		diff = 0.01 * temperatureDiff
	} else {
		diff = 0
	}
	t.Dynamic.Moisture = t.Dynamic.Moisture - (diff / t.Static.BarkThickness)
}
