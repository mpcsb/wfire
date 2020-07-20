package terrain

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"sort"

	"simulation/shared"
	// "gonum.org/v1/gonum/spatial/vptree"

	// "github.com/im7mortal/UTM"
)

type Coord_label struct{
	Coord shared.Coord // lat, lon, alt
	Label string // type of structure in terrain: origin openstreet maps
}
 

type Terrain struct {
	Coord_Type []Coord_label
	Width  float64
	Length float64
	Coord2Alt map[shared.Coord2]float64
	SetLon []float64
	SetLat []float64
}


func CallPythonScripts(p1, p2 shared.Coord, sample_size int, task string) { 

	// python_exec := "../../../bin/python3.8" //to py venv WRONG
	// python_exec := "/usr/bin/python3.8"
	python_exec := "/home/miguel/anaconda3/bin/python3.7"
	filePath := ""

	if task == "altitude"{
		filePath, _ = filepath.Abs("../../simulation/terrain/HGT_parser.py") //relative to cmd folder
	}

	if task == "structures"{
		filePath, _ = filepath.Abs("../../simulation/terrain/generate_objects.py") //relative to cmd folder
	}

	cmd := exec.Command(python_exec, filePath,
		fmt.Sprintf("%f", p1.Lat),
		fmt.Sprintf("%f", p1.Lon),
		fmt.Sprintf("%f", p2.Lat),
		fmt.Sprintf("%f", p2.Lon),
		fmt.Sprintf("%d", sample_size))

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
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


func (t Terrain) GenDimensions(p1 shared.Coord, p2 shared.Coord){ 
	t.Width = shared.Haversine(p1.Lat, p2.Lon, p2.Lat, p2.Lon) 
	t.Length = shared.Haversine(p1.Lat, p1.Lon, p1.Lat, p2.Lon)
}


func (t Terrain) LatLon2Alt()  {
	coord2alt := make(map[shared.Coord2]float64, len(t.Coord_Type))
	for i := range t.Coord_Type{
		coord := shared.Coord2{Lat: t.Coord_Type[i].Coord.Lat, Lon: t.Coord_Type[i].Coord.Lon}
		coord2alt[coord] = t.Coord_Type[i].Coord.Alt
	}
	t.Coord2Alt = coord2alt
}


func (t Terrain) Uniques() {
	set_lat := make(map[float64]bool)
	set_lon := make(map[float64]bool)
	for _, v := range t.Coord_Type{
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


func (t Terrain) Adjacent(p shared.Coord) (float64, float64, float64, float64){
	lat := p.Lat
	lon := p.Lon

	var i int
	var j int
	for i_lat := range t.SetLat{
		if t.SetLat[i_lat] == lat{
			i = i_lat
			break
		}
	}

	for i_lon := range t.SetLon{
		if t.SetLon[i_lon] == lon{
			j = i_lon
			break
		}
	}

	if i + 1 <= len(t.SetLon) && j + 1 <= len(t.SetLat){
		return t.SetLat[i], t.SetLat[i + 1], t.SetLon[j], t.SetLon[j + 1]
	} 
	return 0, 0, 0, 0 
}


// https://en.wikipedia.org/wiki/Bilinear_interpolation
// a linear interpolation is enough given the relative error of the SRTM measurements
// linear vs cubic interp seems justified for 'regular' topographies
func (t Terrain) Binterp(target shared.Coord) float64 {
	x := target.Lat
	y := target.Lon

	x1, x2, y1, y2 := t.Adjacent(target)

	R1 := t.Coord2Alt[shared.Coord2{Lat:x1, Lon:y1}] + (x-x1)/(x2-x1)*(t.Coord2Alt[shared.Coord2{Lat:x2, Lon:y1}]-t.Coord2Alt[shared.Coord2{Lat:x1, Lon:y1}])
	R2 := t.Coord2Alt[shared.Coord2{Lat:x1, Lon:y2}] + (x-x1)/(x2-x1)*(t.Coord2Alt[shared.Coord2{Lat:x2, Lon:y2}]-t.Coord2Alt[shared.Coord2{Lat:x1, Lon:y2}])

	val := R2 + (y-y2)/(y2-y1)*(R1-R2)
	return val
}


func GenerateTerrain(p1, p2 shared.Coord, sample_size int) Terrain {

	CallPythonScripts(p1, p2, sample_size, "altitude") 
	coord_lst := rawTerrain() 
	
	var t Terrain
	for _, v := range coord_lst {  
		t.Coord_Type = append(t.Coord_Type, 
			Coord_label{shared.Coord{Lat: v[0], Lon: v[1], Alt: v[2]}, ""})// TODO get label from coord 2 label map
	}  
	// t.GenDimensions(p1, p2) 
 
	t.Length = shared.Haversine(p1.Lat, p1.Lon, p1.Lat, p2.Lon)
	t.Width = shared.Haversine(p1.Lat, p2.Lon, p2.Lat, p2.Lon)

	t.LatLon2Alt()
	t.Uniques()

	return t
} 