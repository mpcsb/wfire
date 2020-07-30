package fire

import (
	s "simulation/shared"
	"simulation/weather"
)

// Flame contains flame related information such as coordinates and size
type Flame struct {
	Coord            s.Coord
	Height           float64
	Direction        float64 // this should be a polygon shape close like an (circumpherence to ellipsoid) under a force
	Radius           float64
	FlameTemperature float64
	circle           []s.Coord // perimeter composed of a list of 2d coordinates: parabola + circumpherence
	parabola         []s.Coord
}

// DetermineShape implements flame and wind interaction
func (f *Flame) DetermineShape(w weather.Wind) {
	pos := f.Coord
	c, p := Perimeter(20, w.Speed, w.Direction)

	for _, point := range c {
		newPoint := NewCoord_XY(pos, point[0], point[1])
		f.circle = append(f.circle, newPoint)
	}
	for _, point := range p {
		newPoint := NewCoord_XY(pos, point[0], point[1])
		f.parabola = append(f.parabola, newPoint)
	}
}

// UpdateTemperature implements the relation between flame size and temperature
// 1 meter = 800ºc; 50 meter=1200ºc
func (f *Flame) UpdateTemperature() {
	m := 8.0 // 400ºC/50 m
	f.FlameTemperature = 800 + m*f.Height
}
