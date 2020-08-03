package fire

import (
	s "simulation/shared"
)

// Flame contains flame related information such as coordinates and size
type Flame struct {
	Coord       s.Coord
	Height      float64
	Direction   float64 // this should be a polygon shape close like an (circumpherence to ellipsoid) under a force
	Radius      float64
	Temperature float64
	Circle      []s.Coord // perimeter composed of a list of 2d coordinates: parabola + circumpherence
	Parabola    []s.Coord
}

// // DetermineShape implements flame and wind interaction
// func (f *Flame) DetermineShape(w weather.Wind) {
// 	pos := f.Coord
// 	c, p := Perimeter(20, w.Speed, w.Direction)

// 	for _, point := range c {
// 		newPoint := s.NewCoordXY(pos, point[0], point[1])
// 		f.Circle = append(f.Circle, newPoint)
// 	}
// 	for _, point := range p {
// 		newPoint := s.NewCoordXY(pos, point[0], point[1])
// 		f.Parabola = append(f.Parabola, newPoint)
// 	}
// }

// UpdateTemperature implements the relation between flame size and temperature
// 1 meter = 800ºc; 50 meter=1200ºc
func (f *Flame) UpdateTemperature() {
	m := 8.0 // 400ºC/50 m
	f.Temperature = 800 + m*f.Height
}

// // UpdateWindTemperature will update the wind temperature located nearby a set of flames
// // Fourier’s law determines that temperature varies inversely to distance
// // if angle is negative, that flame will have no impact in the added temperature
// func UpdateWindTemperature(f Flame, wm map[s.Coord]weather.Wind) (dTemperature float64) {

// 	lat1 := f.Coord.Lat
// 	lon1 := f.Coord.Lon
// 	for _, w := range wm {
// 		wLat := w.Coord.Lat
// 		wLon := w.Coord.Lon
// 		distance := s.Abs(s.Haversine(lat1, lon1, wLat, wLon))
// 		angle := math.Cos(s.Angle(lat1, lon1, wLat, wLon))
// 		if angle > 0 {
// 			dTemperature = f.Temperature * angle / distance
// 		}
// 	}
// 	return dTemperature
// }
