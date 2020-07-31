package shared

import (
	// "fmt"
	"math"
)

// Sigmoid function
func Sigmoid(x float64) (sig float64) {
	sig = x / (1 + Abs(x))
	return sig
}

// Haversine implements distance between two coordinates in meters
func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const r = 6371000 // m
	sdLat := math.Sin(Radians(lat2-lat1) / 2)
	sdLon := math.Sin(Radians(lon2-lon1) / 2)
	a := sdLat*sdLat + math.Cos(Radians(lat1))*math.Cos(Radians(lat2))*sdLon*sdLon
	d := 2 * r * math.Asin(math.Sqrt(a))
	return d //  m
}

func Angle(lat1, lon1, lat2, lon2 float64) float64 {
	dlon := lon1 - lon2
	return math.Atan2(math.Sin(Radians(dlon))*math.Cos(Radians(lat2)), math.Cos(Radians(lat1))*math.Sin(Radians(lat2))-math.Sin(Radians(lat1))*math.Cos(Radians(lat2))*math.Cos(Radians(dlon)))
}

func Radians(d float64) float64 {
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

func Abs(f float64) float64 {
	if f < 0 {
		return f * -1
	} else {
		return f
	}
}

func Linspace(lower, upper, num float64) []float64 {

	lower = Min(lower, upper)
	upper = Max(lower, upper)
	num_int := int(num)

	var array []float64
	for i := 1; i <= num_int; i++ {
		array = append(array, lower+float64(i)*(upper-lower)/float64(num_int))
	}
	return array
}

func adjacent_points() [][2]int {
	adjacent_coords := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {-1, 0}, {-1, 1}}
	return adjacent_coords
}
