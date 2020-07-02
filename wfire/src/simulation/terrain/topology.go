package terrain

import (
	"fmt"
	"os/exec"
	//"path/filepath"
)

type cart_coord struct {
	x float64
	y float64
	alt int
}

type geo_coord struct {
	lat float64
	lon float64
	alt int
}


type terrain struct{
	coords [] geo_coord
} 

func Generate_map(){
	python_exec := "C:/Users/Miguel/Anaconda3/envs/quest/pythonw.exe"
	command := "C:/Users/Miguel/Documents/repos/wfire/wfire/src/simulation/terrain/HGT/HGT_parser.py-lat1 38.123 -lon1 12.455 -lat2 38.145 -lon2 12.489"
	cmd := exec.Command(python_exec, command)
	err := cmd.Run()
	fmt.Println(err)
} 