package shared

import (
	"math"
)

func Distance(p1 point, p2 point) float64 {
	d := 0.0
	return d
}

func CoordToDistance(c1 Coord) (float64, float64) {
	x := 0.0
	y := 0.0
	return x, y
}

func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const r = 6371000 // m
	sdLat := math.Sin(radians(lat2-lat1) / 2)
	sdLon := math.Sin(radians(lon2-lon1) / 2)
	a := sdLat*sdLat + math.Cos(radians(lat1))*math.Cos(radians(lat2))*sdLon*sdLon
	d := 2 * r * math.Asin(math.Sqrt(a))
	return d //  m
}

func radians(d float64) float64 {
	return d * math.Pi / 180
}
