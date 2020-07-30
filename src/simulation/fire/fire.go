package fire

import (
	"simulation/fuel"
	s "simulation/shared"
)

// Potential function returns the probability of fuel catching fire
// It will take fuel distance, temperature, relative humidity ...
func Potential(tree fuel.Tree_data) {

}

// MergeFronts receives a collection of flames and determines how they are merged
// This is applied to a flame in fuel coordinate, and should receive flame list from neighbours
func (f *Flame) MergeFronts(fire []Flame) {
	factor := 0.0
	for i, flm := range fire {
		factor += flm.Height + float64(i)*0.1
		//factor += 1.0
	}

	f.Height = s.Sigmoid(factor)
	f.UpdateTemperature()
}
