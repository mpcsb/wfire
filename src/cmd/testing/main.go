package main

import (
	"fmt"
	
	"simulation/shared"
	"simulation/terrain"
)

func main() {
	p1 := shared.Coord{Lat: 38.763613, Lon: -9.453429, Alt: 0.0}
	p2 := shared.Coord{Lat: 38.813257, Lon: -9.334138, Alt: 0.0} 
 
	t := terrain.GenerateTerrain(p1, p2, 500)
	terrain.CallPythonScripts(p1, p2, 0, "structures")  

	fmt.Println(t.Length)
	// err := t.GenerateObjets(30, "water") 
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = t.GenerateObjets(50, "natural") 
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = t.GenerateObjets(50, "landuse") 
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = t.GenerateObjets(30, "building") 
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = t.GenerateObjets(20, "highway")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// err = t.GenerateObjets(20, "railway")
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
