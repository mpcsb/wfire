package fire

import (
	"simulation/fuel"
	s "simulation/shared"
)

func Sigmoid(x float64) (sig float64) {
	sig = x / (1 + s.Abs(x))
	return sig
}

func FirePotential(tree fuel.Tree_data) {

}
