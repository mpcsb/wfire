package main

import ( 
	//"fmt"
	
	"simulation/terrain"
)
 


func main() { 
	/*
	fmt.Println(terrain.FilenameGen(-2,-6))
	fmt.Println(terrain.FilenameGen(-22,-66))
	fmt.Println(terrain.FilenameGen(2,6))
	fmt.Println(terrain.FilenameGen(22,66))
	fmt.Println(terrain.FilenameGen(22,-166))
	fmt.Println(terrain.FilenameGen(22,166)) 
	
	terrain.ConvertBigEndian()*/
	terrain.Generate_map()
}
