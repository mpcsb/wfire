package terrain


import (
	"fmt"
	"math"
	"testing"
	"simulation/shared"
)


func TerrainForTest()(ter Terrain){
	p1 := shared.Coord{Lat: 38.773613, Lon: -9.363429, Alt: 0.0}
	p2 := shared.Coord{Lat: 38.813257, Lon: -9.324138, Alt: 0.0} 
	ter = GenerateTerrain(p1, p2, 100)
	return ter
}

func TestInterp(t *testing.T) { 
	ter := TerrainForTest()

	X := 38.793613
	Y := -9.343429
	P := shared.Coord{Lat: X, Lon: Y, Alt:0.0}
  
	interpolated, _, _ := ter.Binterp(P) 
	if math.IsNaN(interpolated){
		t.Errorf("got NaN  %f", interpolated)
	} 
}

func TestAdjacent(t *testing.T) {
	ter := TerrainForTest()

	X := 38.793613
	Y := -9.343429
	P := shared.Coord{Lat: X, Lon: Y, Alt:0.0}
  
	lat0, lat1, lon0, lon1 := ter.Adjacent(P) 
	out := []float64{lat0, lat1, lon0, lon1} 

	for _, n := range out{
		if math.IsNaN(n){
			t.Errorf("got NaN")
		}
	}

	fmt.Println(lat0, X, lat1, lon0, Y, lon1)
	if (lat0 >= lat1)  {
		t.Errorf("Latitude: values are not coherent %f & %f ", lat0, lat1)
	}
	if (lat0 > X) || (lat1 < X) {
		t.Errorf("Latitude: %f is not between %f and %f", X, lat0, lat1)
	}

	if (lon0 >= lon1){
		t.Errorf("Longitude: values are not coherent %f & %f ", lon0, lon1)
	}
	if (lon0 > Y) || (lon1 < Y){
		t.Errorf("Longitude: %f is not between %f and %f", Y, lon0, lon1)
	}

}

func TestCoord2Alt(t *testing.T) {
	ter := TerrainForTest()
	X := 38.773613 
	Y := -9.363429 

	val := 0.0
	if h, ok := ter.Coord2Alt[shared.Coord2{Lat:X, Lon:Y}]; ok {
		val = h
	} else{
		t.Errorf("No key in map, %f", h)
	}

	fmt.Println(val)
	if val == 0 {
		t.Errorf("Height should not be zero: %f", val)
	}	
}

func TestGenDimensions(t *testing.T) {
	p1 := shared.Coord{Lat: 38.773613, Lon: -9.363429, Alt: 0.0}
	p2 := shared.Coord{Lat: 38.813257, Lon: -9.324138, Alt: 0.0} 
	ter := GenerateTerrain(p1, p2, 100)

	ter.GenDimensions(p1, p2) 
	
	if ter.Length <= 0{
		t.Errorf("t.Length should not be zero: %f", ter.Length)
	}
	if ter.Width <= 0{
		t.Errorf("t.Width should not be zero: %f", ter.Width)
	}
}