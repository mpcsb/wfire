package terrain

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"simulation/shared"
	"sort"
	"strconv"
)

type Coord_label struct {
	Coord  shared.Coord // lat, lon, alt
	Slope  float64
	Aspect float64 // type of structure in terrain: origin openstreet maps
}

type Terrain struct {
	Coord_Type    []Coord_label
	Width, Length float64
	Coord2Alt     map[shared.Coord2]float64
	Coord2Slope   map[shared.Coord2]float64
	Coord2Aspect  map[shared.Coord2]float64
	SetLon        []float64
	SetLat        []float64

	MinHeight float64
	MaxHeight float64
}

func CallPythonScripts(p1, p2 shared.Coord, sample_size int, task string) {

	// python_exec := "../../../bin/python3.8" //to py venv WRONG
	// python_exec := "/usr/bin/python3.8"
	python_exec := "/home/miguel/anaconda3/bin/python3.7"
	filePath := ""

	if task == "altitude" {
		filePath, _ = filepath.Abs("../../simulation/terrain/HGT_parser.py") //relative to cmd folder
	}

	if task == "structures" {
		filePath, _ = filepath.Abs("../../simulation/terrain/generate_objects.py") //relative to cmd folder
	}

	cmd := exec.Command(python_exec, filePath,
		fmt.Sprintf("%f", p1.Lat),
		fmt.Sprintf("%f", p1.Lon),
		fmt.Sprintf("%f", p2.Lat),
		fmt.Sprintf("%f", p2.Lon),
		fmt.Sprintf("%d", sample_size))

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String())
	}
}

func rawTerrain() [][]float64 {
	filePath, _ := filepath.Abs("../../simulation/terrain/temp/coords.csv")
	f, _ := os.Open(filePath)

	terrain := [][]float64{} // list of geo coordinates
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		floatRecord := make([]float64, 5)
		for i, s := range record {
			if v, err := strconv.ParseFloat(s, 64); err == nil {
				floatRecord[i] = v
			}
		}
		terrain = append(terrain, floatRecord)
	}
	return terrain
}

func (t *Terrain) GenDimensions(p1 shared.Coord, p2 shared.Coord) {
	t.Width = shared.Haversine(p1.Lat, p2.Lon, p2.Lat, p2.Lon)
	t.Length = shared.Haversine(p1.Lat, p1.Lon, p1.Lat, p2.Lon)
}

func (t *Terrain) LatLon2Alt() {
	coord2alt := make(map[shared.Coord2]float64, len(t.Coord_Type))
	t.Coord2Alt = make(map[shared.Coord2]float64, len(t.Coord_Type))

	for i := range t.Coord_Type {
		coord := shared.Coord2{Lat: t.Coord_Type[i].Coord.Lat, Lon: t.Coord_Type[i].Coord.Lon}
		coord2alt[coord] = t.Coord_Type[i].Coord.Alt
	}

	for k, v := range coord2alt {
		t.Coord2Alt[k] = v
	}
}

func (t *Terrain) LatLon2Slope() {
	coord2Slope := make(map[shared.Coord2]float64, len(t.Coord_Type))
	t.Coord2Slope = make(map[shared.Coord2]float64, len(t.Coord_Type))

	for i := range t.Coord_Type {
		coord := shared.Coord2{Lat: t.Coord_Type[i].Coord.Lat, Lon: t.Coord_Type[i].Coord.Lon}
		coord2Slope[coord] = t.Coord_Type[i].Slope
	}
	for k, v := range coord2Slope {
		t.Coord2Slope[k] = v
	}
}

func (t *Terrain) LatLon2Aspect() {
	coord2aspect := make(map[shared.Coord2]float64, len(t.Coord_Type))
	t.Coord2Aspect = make(map[shared.Coord2]float64, len(t.Coord_Type))

	for i := range t.Coord_Type {
		coord := shared.Coord2{Lat: t.Coord_Type[i].Coord.Lat, Lon: t.Coord_Type[i].Coord.Lon}
		coord2aspect[coord] = t.Coord_Type[i].Aspect
	}

	for k, v := range coord2aspect {
		t.Coord2Aspect[k] = v
	}
}

func (t *Terrain) Uniques() {
	set_lat := make(map[float64]bool)
	set_lon := make(map[float64]bool)

	for _, v := range t.Coord_Type {
		lat, lon := v.Coord.Lat, v.Coord.Lon
		set_lat[lat] = true
		set_lon[lon] = true
	}

	var LatKeys, LonKeys []float64
	for k := range set_lat {
		LatKeys = append(LatKeys, k)
	}
	sort.Float64s(LatKeys)

	for k := range set_lon {
		LonKeys = append(LonKeys, k)
	}
	sort.Float64s(LonKeys)

	t.SetLat = LatKeys
	t.SetLon = LonKeys
}

func (t *Terrain) Adjacent(p shared.Coord) (x1 float64, x2 float64, y1 float64, y2 float64) {

	var i int
	for i_lat, v := range t.SetLat {
		if v >= p.Lat {
			i = i_lat
			break
		}
	}

	var j int
	for i_lon, v := range t.SetLon {
		if v >= p.Lon {
			j = i_lon
			break
		}
	}

	// latitude
	if i == 0 {
		x1 = t.SetLat[0]
		x2 = t.SetLat[1]
	} else if i == len(t.SetLat) {
		x1 = t.SetLat[len(t.SetLat)-1]
		x2 = t.SetLat[len(t.SetLat)]
	} else {
		x1, x2 = t.SetLat[i-1], t.SetLat[i]
	}
	// longitude
	if j == 0 {
		y1 = t.SetLon[0]
		y2 = t.SetLon[1]
	} else if j == len(t.SetLon) {
		y1 = t.SetLon[len(t.SetLon)-1]
		y2 = t.SetLon[len(t.SetLon)]
	} else {
		y1, y2 = t.SetLon[j-1], t.SetLon[j]
	}
	return x1, x2, y1, y2
}

// https://en.wikipedia.org/wiki/Bilinear_interpolation
// a linear interpolation is enough given the relative error of the SRTM measurements
// linear vs cubic interp seems justified for 'regular' topographies
func (t Terrain) Binterp(target shared.Coord) (float64, float64, float64) {
	// x := target.Lat
	// y := target.Lon

	x1, x2, y1, y2 := t.Adjacent(target)

	// R1 := t.Coord2Alt[shared.Coord2{Lat:x1, Lon:y1}] + (x-x1)/(x2-x1)*(t.Coord2Alt[shared.Coord2{Lat:x2, Lon:y1}]-t.Coord2Alt[shared.Coord2{Lat:x1, Lon:y1}])
	// R2 := t.Coord2Alt[shared.Coord2{Lat:x1, Lon:y2}] + (x-x1)/(x2-x1)*(t.Coord2Alt[shared.Coord2{Lat:x2, Lon:y2}]-t.Coord2Alt[shared.Coord2{Lat:x1, Lon:y2}])
	// altitude := R2 + (y-y2)/(y2-y1)*(R1-R2)

	altitude := (t.Coord2Alt[shared.Coord2{Lat: x1, Lon: y1}] +
		t.Coord2Alt[shared.Coord2{Lat: x1, Lon: y2}] +
		t.Coord2Alt[shared.Coord2{Lat: x2, Lon: y1}] +
		t.Coord2Alt[shared.Coord2{Lat: x2, Lon: y2}]) * 0.25

	slope := (t.Coord2Slope[shared.Coord2{Lat: x1, Lon: y1}] +
		t.Coord2Slope[shared.Coord2{Lat: x1, Lon: y2}] +
		t.Coord2Slope[shared.Coord2{Lat: x2, Lon: y1}] +
		t.Coord2Slope[shared.Coord2{Lat: x2, Lon: y2}]) * 0.25

	aspect := (t.Coord2Alt[shared.Coord2{Lat: x1, Lon: y1}] +
		t.Coord2Aspect[shared.Coord2{Lat: x1, Lon: y2}] +
		t.Coord2Aspect[shared.Coord2{Lat: x2, Lon: y1}] +
		t.Coord2Aspect[shared.Coord2{Lat: x2, Lon: y2}]) * 0.25

	aspect = shared.Radians(aspect)
	return altitude, slope, aspect
}

func (t *Terrain) HeightExtremes() {
	var heights []float64

	for k := range t.Coord2Alt {
		heights = append(heights, t.Coord2Alt[k])
	}

	min, max := 10000.0, -10000.0
	for _, h := range heights {
		if h > max {
			max = h
		}
		if h < min {
			min = h
		}
	}

	t.MinHeight = min
	t.MaxHeight = max

	fmt.Println(t.MinHeight, t.MaxHeight)
}

func GenerateTerrain(p1, p2 shared.Coord, sample_size int) (t Terrain) {

	CallPythonScripts(p1, p2, sample_size, "altitude")
	coord_lst := rawTerrain()

	for _, v := range coord_lst {
		t.Coord_Type = append(t.Coord_Type,
			Coord_label{shared.Coord{Lat: v[0], Lon: v[1], Alt: v[2]}, v[3], v[4]}) // Slope: v[3], Aspect:v[4]
	}

	t.GenDimensions(p1, p2)
	t.Uniques()
	t.LatLon2Alt()
	t.HeightExtremes()
	return t
}
