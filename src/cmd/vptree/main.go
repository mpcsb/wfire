package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"

	"gonum.org/v1/gonum/spatial/vptree"
)

func main() {
	r := rand.New(rand.NewSource(99))

	var stations = []vptree.Comparable{
		place{name: "00", lat: 51.0 + r.Float64(), lon: r.Float64()},
	}
	for i := 0; i < 1000000; i++ {
		new_place := place{name: fmt.Sprintf("%d", i), lat: 51.0 + r.Float64(), lon: r.Float64()}
		stations = append(stations, new_place)
	}
	fmt.Println(len(stations))
	// Construct a vp tree of train station locations
	// to identify accessible public transport for the
	// elderly.
	t, err := vptree.New(stations, 5, nil)
	if err != nil {
		log.Fatal(err)
	}

	for index, q := range stations {
		// Residence.
		// q := place{lat: 51.501476, lon: 0.140634}

		var keep vptree.Keeper

		// // Find all stations within 0.75 of the residence.
		keep = vptree.NewDistKeeper(0.25)
		t.NearestSet(keep, q)

		if len(keep.(*vptree.DistKeeper).Heap) == 1 {
			continue
		}

		if index%50000 == 0 {
			fmt.Println(index)
		}
		// fmt.Println(fmt.Sprintf("%f %f", q.lat, q.lon))
		// for _, c := range keep.(*vptree.DistKeeper).Heap {
		// 	p := c.Comparable.(place)
		// 	fmt.Printf("%s: %0.3f km\n", p.name, p.Distance(q))
		// }
		// fmt.Println()
	}
	fmt.Println("complete")
	// Find the five closest stations to the residence.
	// keep = vptree.NewNKeeper(5)
	// t.NearestSet(keep, q)

	// fmt.Println(`5 closest stations to 51.501476N 0.140634W.`)
	// for _, c := range keep.(*vptree.NKeeper).Heap {
	// 	p := c.Comparable.(place)
	// 	fmt.Printf("%s: %0.3f km\n", p.name, p.Distance(q))
	// }
} 

// place is a vptree.Comparable implementations.
type place struct {
	name     string
	lat, lon float64
}

// Distance returns the distance between the receiver and c.
func (p place) Distance(c vptree.Comparable) float64 {
	q := c.(place)
	return haversine(p.lat, p.lon, q.lat, q.lon)
}

// haversine returns the distance between two geographic coordinates.
func haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const r = 6371 // km
	sdLat := math.Sin(radians(lat2-lat1) / 2)
	sdLon := math.Sin(radians(lon2-lon1) / 2)
	a := sdLat*sdLat + math.Cos(radians(lat1))*math.Cos(radians(lat2))*sdLon*sdLon
	d := 2 * r * math.Asin(math.Sqrt(a))
	return d // km
}

func radians(d float64) float64 {
	return d * math.Pi / 180
}
