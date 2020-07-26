package main

import ( 
	"simulation/shared"
	"simulation/forest"
)

func main() {
	p1 := shared.Coord{Lat: 38.793613, Lon: -9.383429, Alt: 0.0}
	p2 := shared.Coord{Lat: 38.813257, Lon: -9.354138, Alt: 0.0} 
 
	f := forest.ForestGeneration(p1, p2, 100) 

	for i:=0; i<36; i++{
		f.RecordFrame()
	}
	f.Plot_forest()
}