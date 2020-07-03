package terrain

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type cart_coord struct {
	x   float64
	y   float64
	alt float64
}

type geo_coord struct {
	lat float64
	lon float64
	alt float64
}

type terrain struct {
	coords []geo_coord
	width  float64
	length float64
}

func extractCoordinates() {
	python_exec := "C:/Users/Miguel/Anaconda3/envs/quest/pythonw.exe"
	filePath, _ := filepath.Abs("../../simulation/terrain/HGT/HGT_parser.py")
	params := " -lat1 38.123 -lon1 12.455 -lat2 38.145 -lon2 12.489"
	command := filePath + params
	cmd := exec.Command(python_exec, command)
	err := cmd.Run()
	fmt.Println(err)
}

func rawTerrain() [][]float64 {

	filePath, _ := filepath.Abs("../../simulation/terrain/HGT/coords.csv")
	//fmt.Println("file is:", filePath)
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

func genDimensions() (float64, float64) {
	fmt.Println("TODO")
	return 0.0, 0.2
}

func GenerateTerrain() terrain {

	//extractCoordinates()
	coord_lst := rawTerrain()

	t := terrain{}
	for _, v := range coord_lst {
		t.coords = append(t.coords, geo_coord{lat: v[0], lon: v[1], alt: v[2]})
	}
	t.width, t.length = genDimensions()

	return t
}
