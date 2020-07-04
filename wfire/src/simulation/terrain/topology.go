package terrain

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/im7mortal/UTM"
)

type Coord struct {
	lat, lon, alt float64
}

type terrain struct {
	Coords []Coord
	width  float64
	length float64
}

func extractCoordinates() {
	python_exec := "C:/Users/Miguel/Anaconda3/envs/quest/pythonw.exe"
	filePath, _ := filepath.Abs("../../simulation/terrain/HGT/HGT_parser.py")
	params := " -lat1 38.123 -lon1 12.455 -lat2 38.145 -lon2 12.489"
	command := filePath + params

	fmt.Println(command)
	cmd := exec.Command(python_exec, command)
	err := cmd.Run()
	fmt.Println(err)
}

func rawTerrain() [][]float64 {

	filePath, _ := filepath.Abs("../../simulation/terrain/HGT/coords.csv")
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

func genDimensions(p1 Coord, p2 Coord) (float64, float64) {
	x1, y1, _, _, _ := UTM.FromLatLon(p1.lat, p1.lon, true)
	x2, y2, _, _, _ := UTM.FromLatLon(p2.lat, p2.lon, true)
	return x2 - x1, y2 - y1
}

func GenerateTerrain(p1 Coord, p2 Coord) terrain {
	extractCoordinates()
	coord_lst := rawTerrain()

	t := terrain{}
	for _, v := range coord_lst {
		t.Coords = append(t.Coords, Coord{lat: v[0], lon: v[1], alt: v[2]})
	}
	t.width, t.length = genDimensions(p1, p2)
	return t
}
