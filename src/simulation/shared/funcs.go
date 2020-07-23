package shared

import (
	// "fmt"
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

 

func Min(a, b float64) float64 {
    if a < b {
        return a
    }
    return b
}

func Max(a, b float64) float64 {
    if a > b {
        return a
    }
    return b
}


func Linspace(lower, upper, num float64) []float64 {
 
	lower = Min(lower, upper)
	upper = Max(lower, upper) 
	num_int := int(num) 

	var array []float64
	for i := 1; i <= num_int; i++ {
		array = append(array, lower + float64(i) * (upper - lower)/ float64(num_int))
	} 
	return array
}

func adjacent_points()[][2]int{
    adjacent_coords := [][2]int{{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{-1,0},{-1,1}}
    return adjacent_coords
}



