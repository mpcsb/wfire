package terrain

import(
	"encoding/csv"
	"fmt"
	"io"
	"os"
	// "os/exec"
	"path/filepath"
	"strconv"
	"simulation/shared"

	"gonum.org/v1/gonum/spatial/vptree"
)

type Coord_label2 struct {
	Lat, Lon float64
	Label     string
}


func (p Coord_label2) Distance(c vptree.Comparable) float64 {
	q := c.(Coord_label2)
	return shared.Haversine(p.Lat, p.Lon, q.Lat, q.Lon)
}


func rawObjects(tag string) []Coord_label2 {
	filePath, _ := filepath.Abs("../../simulation/terrain/temp/" + tag + "_coordinates.csv")
	fmt.Println("gen objects",filePath)
	f, _ := os.Open(filePath)

	objects := []Coord_label2{} // list of geo coordinates
	r := csv.NewReader(f)
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		
		Lat, _ := strconv.ParseFloat(record[0], 64)
		Lon, _ := strconv.ParseFloat(record[1], 64)
		Label := record[2]

		var line Coord_label2
		line.Lat = Lat
		line.Lon = Lon
		line.Label = Label 

		objects = append(objects, line)
	} 
	return objects
}


func GenerateObjets(ter Terrain, tag string){ 

	tag_points := rawObjects(tag) 
	var coordinates_tag = []vptree.Comparable{} 
	
	for i, _ := range tag_points {
		labelled_coord := Coord_label2{Lat: tag_points[i].Lat, Lon:tag_points[i].Lon, Label: tag_points[i].Label}
		coordinates_tag = append(coordinates_tag, labelled_coord)
	}

	t, _ := vptree.New(coordinates_tag, 1, nil)
 
	coord_lst := ter.Coord_Type
	for _, v := range coord_lst {
		
		v_coord := v.Coord
		var keep vptree.Keeper
		keep = vptree.NewNKeeper(8) // 8 adjacent points in lattice
		tree := Coord_label2{Lat:v_coord.Lat, Lon:v_coord.Lon, Label:""} 
		t.NearestSet(keep, tree)

		for _, c := range keep.(*vptree.NKeeper).Heap {
			p := c.Comparable.(Coord_label2)
			fmt.Println(p.Label, p.Distance(tree))
			// fmt.Println(p.Label, p.Distance(tree), tree.Lat, tree.Lon, p.Lat, v_coord.Lat)
		} 
		fmt.Println()
	}
}


func adjacent_points(p shared.Coord){
	adjacent_coords := [][2]int{{-1,-1},{-1,0},{-1,1},{0,-1},{0,1},{1,-1},{-1,0},{-1,1}}
	fmt.Println(adjacent_coords)
}

// def adjacent_points(la, lo, h):
//     i_lat = latitudes.index(la)
//     i_lon = longitudes.index(lo)

//     adjacent_coords = [(i,j) for i in [-1, 0, 1] for j in [-1, 0, 1]]
//     adjacent = list()
//     for i, j in adjacent_coords:
//         if i == 0 and j == 0: continue
//         try:
//             latitude, longitude = latitudes[i_lat + i], longitudes[i_lon + j]
//             adjacent.append((latitude, longitude, altitude_dict[latitude, longitude]))
//         except:
//             pass
//     return adjacent