package terrain

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"simulation/shared"

	"github.com/im7mortal/UTM"
)

type Coord_label struct{
	coord shared.Coord // lat, lon, alt
	label string // type of structure in terrain: origin openstreet maps
}

type Terrain struct {
	Coord_Type []Coord_label
	Width  float64
	Length float64
}

func CallPythonScripts(p1, p2 shared.Coord, task string) { 

	python_exec := "../../../bin/python3.8" //to py venv WRONG
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
		fmt.Sprintf("%f", p2.Lon))

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

		floatRecord := make([]float64, 3)
		for i, s := range record {
			if v, err := strconv.ParseFloat(s, 64); err == nil {
				floatRecord[i] = v
			}
		}
		terrain = append(terrain, floatRecord)
	}
	return terrain
}

func genDimensions(p1 shared.Coord, p2 shared.Coord) (float64, float64) {
	x1, y1, _, _, _ := UTM.FromLatLon(p1.Lat, p1.Lon, true)
	x2, y2, _, _, _ := UTM.FromLatLon(p2.Lat, p2.Lon, true)
	return x2 - x1, y2 - y1
}


func GenerateTerrain(p1, p2 shared.Coord) Terrain {
	CallPythonScripts(p1, p2, "altitude")
	coord_lst := rawTerrain()

	t := Terrain{}
	
	for _, v := range coord_lst {
		l := "undetermined" // TODO get label from coord 2 label map
		t.Coord_Type = append(t.Coord_Type, 
			Coord_label{shared.Coord{Lat: v[0], Lon: v[1], Alt: v[2]}, l})
	}

	t.Width, t.Length = genDimensions(p1, p2)
	
	CallPythonScripts(p1, p2, "structures")
	return t
}

 