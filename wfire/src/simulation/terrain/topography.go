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

// type Coord struct {
// 	Lat, Lon, Alt float64
// }

type Terrain struct {
	Coords []shared.Coord
	Width  float64
	Length float64
}

func ExtractCoordinates(p1 shared.Coord, p2 shared.Coord) {
	fmt.Println("Extracting coordinates")

	python_exec := "C:/Users/Miguel/Anaconda3/envs/ffire/python.exe"
	filePath, _ := filepath.Abs("../../simulation/terrain/HGT_parser.py")
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

func GenerateTerrain(p1 shared.Coord, p2 shared.Coord) Terrain {
	ExtractCoordinates(p1, p2)
	coord_lst := rawTerrain()

	t := Terrain{}
	for _, v := range coord_lst {
		t.Coords = append(t.Coords, shared.Coord{Lat: v[0], Lon: v[1], Alt: v[2]})
	}
	t.Width, t.Length = genDimensions(p1, p2)
	return t
}
