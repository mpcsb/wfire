package main

import ( 
	"simulation/shared"
	"simulation/forest"
)

func main() {
	p1 := shared.Coord{Lat: 38.773613, Lon: -9.393429, Alt: 0.0}
	p2 := shared.Coord{Lat: 38.873257, Lon: -9.294138, Alt: 0.0} 
 
	f := forest.ForestGeneration(p1, p2, 400) 

	for i:=0; i<36; i++{
		f.RecordFrame()
	}
	f.Plot_forest()
}