package interactions

import (
	"math"
	"simulation/fire"
	s "simulation/shared"
	"simulation/weather"
)

// FlameImport fire.Flame
type FlameImport fire.Flame

// WindImport weather.Wind
type WindImport weather.Wind

// DetermineShape implements flame and wind interaction
func (f *FlameImport) DetermineShape(w weather.Wind) {
	pos := f.Coord
	c, p := fire.Perimeter(20, w.Speed, w.Direction)

	for _, point := range c {
		newPoint := s.NewCoordXY(pos, point[0], point[1])
		f.Circle = append(f.Circle, newPoint)
	}
	for _, point := range p {
		newPoint := s.NewCoordXY(pos, point[0], point[1])
		f.Parabola = append(f.Parabola, newPoint)
	}
}

// UpdateWindTemperature will update the wind temperature located nearby a set of flames
func (w *WindImport) UpdateWindTemperature(fire []fire.Flame) {

	factor := 0.0
	for _, flm := range fire {
		factor += flm.Temperature
		flm.Height = s.Sigmoid(factor*0.01) * 100 // 100ºC seems to be the air temperature near fire fronts
	}
}

// UpdateWindTemperature will update the wind temperature located nearby a set of flames
// Fourier’s law determines that temperature varies inversely to distance
// if angle is negative, that flame will have no impact in the added temperature
func UpdateWindTemperature(f FlameImport, wm map[s.Coord]weather.Wind) (dTemperature float64) {

	lat1 := f.Coord.Lat
	lon1 := f.Coord.Lon
	for _, w := range wm {
		wLat := w.Coord.Lat
		wLon := w.Coord.Lon
		distance := s.Abs(s.Haversine(lat1, lon1, wLat, wLon))
		angle := math.Cos(s.Angle(lat1, lon1, wLat, wLon))
		if angle > 0 {
			dTemperature = f.Temperature * angle / distance
		}
	}
	return dTemperature
}
