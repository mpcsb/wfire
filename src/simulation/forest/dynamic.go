package forest

import (
	"os"
	"encoding/csv"
	"fmt"
	"os/exec"
	"path/filepath"
	"bytes"
)

type Record struct {
    Lat   float64
	Lon   float64
	Alt   float64
    State string
}



func WriteCsv(fname string, m map[int]Record){ 
	file, _ := os.Create(fname + ".csv")
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// define column headers
	headers := []string{
		"Lat",
		"Lon",
		"Alt",
		"State",
	}

	writer.Write(headers)  
	for key := range m {
		r := make([]string, 0, 1+len(headers)) 
		s := fmt.Sprintf("%f,%f,%f,%v", m[key].Lat, m[key].Lon, m[key].Alt, m[key].State)
		r = append(r, s)
		writer.Write(r)
	}
}
 

func (f *Forest) RecordFrame(){
	m := make(map[int]Record)
	for i, LstIdx := range f.Sample_trees{
		m[i] = Record{Lat:f.Tree_lst[LstIdx].Coords.Lat, 
					  Lon:f.Tree_lst[LstIdx].Coords.Lon, 
					  Alt:f.Tree_lst[LstIdx].Coords.Alt, 
					  State:f.Tree_lst[LstIdx].Dynamic.State}
	} 
	WriteCsv("../../simulation/viz/" + fmt.Sprintf("%d", f.Frame), m) 
 
	f.Frame += 1
}

func (f Forest) Plot_forest(){
	python_exec := "/home/miguel/anaconda3/bin/python3.7"
	filePath, _ := filepath.Abs("../../simulation/forest/plot_forest.py")
	cmd := exec.Command(python_exec, filePath)
 
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(fmt.Sprint(err) + ": " + stderr.String()) 
	}
}