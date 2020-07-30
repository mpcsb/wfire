package fire

import (
	"simulation/fuel"
	s "simulation/shared"
	"simulation/weather"
)

type Flame struct {
	Coord            s.Coord
	Height           float64
	Direction        float64 // this should be a polygon shape close like an (circumpherence to ellipsoid) under a force
	Length           float64
	FlameTemperature float64
	circle           []s.Coord // perimeter composed of a list of 2d coordinates: parabola + circumpherence
	parabola         []s.Coord
}

type Fire struct {
	Flame_lst     []Flame
	IgnitionPoint s.Coord
}

func Sigmoid(x float64) (sig float64) {
	sig = x / (1 + s.Abs(x))
	return sig
}

func (f *Flame) DetermineShape(w weather.Wind) {
	pos := f.Coord
	c, p := Perimeter(20, w.Speed, w.Direction)

	for _, point := range c {
		new_point := NewCoord_XY(pos, point[0], point[1])
		f.circle = append(f.circle, new_point)
	}
	for _, point := range p {
		new_point := NewCoord_XY(pos, point[0], point[1])
		f.parabola = append(f.parabola, new_point)
	}
}

func FirePotential(tree fuel.Tree_data) {

}
