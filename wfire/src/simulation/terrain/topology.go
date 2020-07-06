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
	Lat, Lon, Alt float64
}

type Terrain struct {
	Coords []Coord
	Width  float64
	Length float64
}

func ExtractCoordinates(p1 Coord, p2 Coord) {

	python_exec := "C:/Users/Miguel/Anaconda3/envs/ffire/python.exe"
	filePath, _ := filepath.Abs("../../simulation/terrain/HGT/HGT_parser.py")
	// params := "-lat1 38.123 -lon1 12.455 -lat2 38.145 -lon2 12.489"
	// args := []string{"-lat1 38.123", "-lon1 12.455", "-lat2 38.145", "-lon2 12.489"}
	cmd := exec.Command(python_exec, filePath, "38.123", "12.455", "38.145", "12.489")
	fmt.Println(cmd.Args)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

	// stdout, err := cmd.Output()
	// if err != nil {
	// 	fmt.Println("Terminal error:", err.Error())
	// 	return
	// }
	// fmt.Println(string(stdout))
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
	x1, y1, _, _, _ := UTM.FromLatLon(p1.Lat, p1.Lon, true)
	x2, y2, _, _, _ := UTM.FromLatLon(p2.Lat, p2.Lon, true)
	return x2 - x1, y2 - y1
}

func GenerateTerrain(p1 Coord, p2 Coord) Terrain {
	ExtractCoordinates(p1, p2)
	coord_lst := rawTerrain()

	t := Terrain{}
	for _, v := range coord_lst {
		t.Coords = append(t.Coords, Coord{Lat: v[0], Lon: v[1], Alt: v[2]})
	}
	t.Width, t.Length = genDimensions(p1, p2)
	return t
}
