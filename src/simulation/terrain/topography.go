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
	"gonum.org/v1/gonum/spatial/vptree"

	"github.com/im7mortal/UTM"
)

type Coord_label struct{
	Coord shared.Coord // lat, lon, alt
	Label string // type of structure in terrain: origin openstreet maps
}

type Coord_label2 struct {
	Lat, Lon float64
	Label     string
}


type Terrain struct {
	Coord_Type []Coord_label
	Width  float64
	Length float64
}

func CallPythonScripts(p1, p2 shared.Coord, task string) { 

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
		fmt.Sprintf("%f", p2.Lon))

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
} 

func rawTerrain() [][]float64 {
	filePath, _ := filepath.Abs("../../simulation/terrain/temp/coords.csv")
	fmt.Println("gen terrain",filePath)
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


func GenerateObjects(tag string) []Coord_label2 {
	filePath, _ := filepath.Abs("../../simulation/terrain/temp/" + tag + "_coordinates.csv")
	fmt.Println("gen objects",filePath)
	f, _ := os.Open(filePath)

	objects := []Coord_label2{} // list of geo coordinates
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		
		Lat, _ := strconv.ParseFloat(record[0], 64)
		Lon, _ := strconv.ParseFloat(record[1], 64)
		Label := record[2]

		var line Coord_label2
		line.Lat = Lat
		line.Lon = Lon
		line.Label = Label 

		objects = append(objects, line)
	} 
	return objects
}


func genDimensions(p1 shared.Coord, p2 shared.Coord) (float64, float64) {
	x1, y1, _, _, _ := UTM.FromLatLon(p1.Lat, p1.Lon, true)
	x2, y2, _, _, _ := UTM.FromLatLon(p2.Lat, p2.Lon, true)
	return x2 - x1, y2 - y1
}

func (p Coord_label2) Distance(c vptree.Comparable) float64 {
	q := c.(Coord_label2)
	return shared.Haversine(p.Lat, p.Lon, q.Lat, q.Lon)
}

func GenerateTerrain(p1, p2 shared.Coord) Terrain {

	CallPythonScripts(p1, p2, "altitude")
	CallPythonScripts(p1, p2, "structures")

	coord_lst := rawTerrain() 
	landuse := GenerateObjects("landuse")
	fmt.Println(landuse[:10])

	var landuse_coords = []vptree.Comparable{} //Coord_label2{38.7928606, -9.4202621, "forest"}}
	for i, _ := range landuse {
		labelled_coord := Coord_label2{Lat: landuse[i].Lat, Lon:landuse[i].Lon, Label: landuse[i].Label}
		landuse_coords = append(landuse_coords, labelled_coord)
	}

	t, err := vptree.New(landuse_coords, 0, nil)
	fmt.Println(err)
	
	var ter Terrain
	for _, v := range coord_lst {
		
		var keep vptree.Keeper
		keep = vptree.NewNKeeper(1) // 8 adjacent points in lattice
		tree := Coord_label2{Lat:v[0], Lon:v[1], Label:""} 
		t.NearestSet(keep, tree)

		for _, c := range keep.(*vptree.NKeeper).Heap {
			p := c.Comparable.(Coord_label2)
			fmt.Println(p.Label, p.Distance(tree), tree.Lat, tree.Lon, v[0], v[1], keep)
		}
		fmt.Println()
		l := "undetermined" // TODO get label from coord 2 label map
		ter.Coord_Type = append(ter.Coord_Type, 
			Coord_label{shared.Coord{Lat: v[0], Lon: v[1], Alt: v[2]}, l})
	}



	ter.Width, ter.Length = genDimensions(p1, p2)
	return ter
}

 