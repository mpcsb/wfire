package shared

import (
	"math"
)

// func Distance(p1 point, p2 point) float64 {
// 	d := 0.0
// 	return d
// }

// func CoordToDistance(c1 Coord) (float64, float64) {
// 	x := 0.0
// 	y := 0.0
// 	return x, y
// }

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
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


func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}


func binterp(f map[[2]float64]float64,  x1, x2, y1, y2 float64, p Coord) float64 {
	x , y := p.Lat, p.Lon
 
	R1 := f[[2]float64{x1, y1}] + (x-x1)/(x2-x1)*(f[[2]float64{x2, y1}] - f[[2]float64{x1, y1}])
	R2 := f[[2]float64{x1, y2}] + (x-x1)/(x2-x1)*(f[[2]float64{x2, y2}] - f[[2]float64{x1, y2}]) 
	alt := R2 + (y-y2)/(y2-y1)*(R1-R2)
	return alt
}