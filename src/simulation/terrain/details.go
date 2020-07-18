package terrain

import(
	"encoding/csv"
	"fmt"
	"io"
	"os" 
	"path/filepath"
	"strconv"
	// "simulation/shared"

	// "gonum.org/v1/gonum/spatial/vptree"
)

type Coord_label2 struct {
	Lat, Lon float64
	Label     string
}


// func (p Coord_label2) Distance(c vptree.Comparable) float64 {
// 	q := c.(Coord_label2)
// 	return shared.Haversine(p.Lat, p.Lon, q.Lat, q.Lon)
// }


func rawObjects(tag string) ([]Coord_label2 , error){
	filePath, _ := filepath.Abs("../../simulation/terrain/temp/" + tag + "_coordinates.csv")

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
	if len(objects) == 0{
		return objects, fmt.Errorf(tag + ": No objects in terrain")
	} 
	return objects, nil
}


func ObjectCoordinates()[]Coord_label2{
	var object_lst []Coord_label2
	tags := []string{"structures", "natural", "water", "landuse", "building", "highway", "railway"}
	for _, tag := range tags{
		objects, err := rawObjects(tag)
		if err != nil{
			continue
		}
		for i := range(objects){
			object_lst = append(object_lst, objects[i])
		} 
	}
	return object_lst
}

// func ProcessTaggedCoords(){
// 	objCoords := ObjectCoordinates()
// 	curatedObjects := []Coord_label2

// 	for _, v := range objCoords{
// 		lat, lon, tag := v.Lat, v.Lon, v.Label
// 	}

// }

// func (ter Terrain) GenerateObjets(dist float64, tag string) (err error){  

// 	tag_points := rawObjects(tag) 
// 	var coordinates_tag = []vptree.Comparable{} 

// 	if len(tag_points) == 0{
// 		return fmt.Errorf(tag + ": No objects in terrain")
// 	} 

// 	for i, _ := range tag_points {
// 		labelled_coord := Coord_label2{Lat: tag_points[i].Lat, Lon:tag_points[i].Lon, Label: tag_points[i].Label}
// 		coordinates_tag = append(coordinates_tag, labelled_coord)
// 	}

// 	t, _ := vptree.New(coordinates_tag, 1, nil)
 
// 	coord_lst := ter.Coord_Type

// 	for i, v := range coord_lst {  // sample size 
		
// 		v_coord := v.Coord
// 		var keep vptree.Keeper
// 		keep = vptree.NewNKeeper(1) // 8 adjacent points in lattice
// 		tree := Coord_label2{Lat:v_coord.Lat, Lon:v_coord.Lon, Label:""} 
// 		t.NearestSet(keep, tree)

// 		for _, c := range keep.(*vptree.NKeeper).Heap {
// 			p := c.Comparable.(Coord_label2) 
// 			if p.Distance(tree) < dist{
// 				coord_lst[i].Label = p.Label
// 				// fmt.Println(p.Label, p.Distance(tree))
// 			}
// 		}  
// 	}

// 	return 
// }




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