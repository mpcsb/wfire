package terrain

import (
	"os"
	"fmt"
	"strconv"  
	"compress/gzip" 
	"io/ioutil"
	"encoding/binary"
)

type coord struct {
    lat float64 
	lon float64
	alt int
}

func Abs(i int) int {
    if i >= 0{
		return i
	}else{
		return -1 * i
	}
}

func str_change(str string, replacement string, index int) string {
	return str[:index] + string(replacement) + str[index+1:] 
}


func FilenameGen(lat int, lon int) string {
	fname := " 00 000" // file has 7 chars
	
	//latitude 
	s_lat := strconv.Itoa(lat)
	if lat >= 0 {
		fname = str_change(fname, "N", 0)
		if (lat >= 10){
			fname = str_change(fname, string(s_lat[0]), 1)  
			fname = str_change(fname, string(s_lat[1]), 2)
			}  
		if (lat >= 0 && lat < 10){  
			fname = str_change(fname, "0", 1)  
			fname = str_change(fname, string(s_lat[0]), 2)  
		}
	}else{
		fname = str_change(fname, "S", 0) 
			if Abs(lat) >= 10{  
				fname = str_change(fname, string(s_lat[1]), 1) 
				fname = str_change(fname, string(s_lat[2]), 2)
			}
			if(Abs(lat) >= 0 && Abs(lat) < 10){ 
				fname = str_change(fname, "0", 1)  
				fname = str_change(fname, string(s_lat[1]), 2)  
			}
		}
	//longitude
	s_lon := strconv.Itoa(lon) 
	if lon >= 0 {
		fname = str_change(fname, "E", 3)
		if (lon >= 100){
			fname = str_change(fname, string(s_lon[0]), 4)  
			fname = str_change(fname, string(s_lon[1]), 5)
			fname = str_change(fname, string(s_lon[2]), 6)
			}  
		if (lon >= 10){
			fname = str_change(fname, string(s_lon[0]), 5)  
			fname = str_change(fname, string(s_lon[1]), 6)
			}  
		if (lon >= 0 && lon < 10){  
			fname = str_change(fname, "0", 5)  
			fname = str_change(fname, string(s_lon[0]), 6)  
		}
	}else{
		fname = str_change(fname, "W", 3) 
		if Abs(lon) >= 100{ 
			fname = str_change(fname, string(s_lon[1]), 4)  
			fname = str_change(fname, string(s_lon[2]), 5) 
			fname = str_change(fname, string(s_lon[3]), 6)
		}
			if Abs(lon) >= 10{  
				fname = str_change(fname, string(s_lon[1]), 5) 
				fname = str_change(fname, string(s_lon[2]), 6)
			}
			if(Abs(lon) >= 0 && Abs(lon) < 10){ 
				fname = str_change(fname, "0", 5)  
				fname = str_change(fname, string(s_lon[1]), 6)  
			}
		}
		fname = fname + ".hgt.gz" 
    return fname
}


func check(e error) {
    if e != nil {
        panic(e)
    }
}

/* 
func DecompressHGT(){
	filename := "C:\\Users\\Miguel\\Documents\\repos\\wfire\\wfire\\src\\test_data\\N38W009.hgt.gz"
	f, _ := os.Open(filename)
	reader, _ := gzip.NewReader(f)
    // Empty byte slice.
    result := make([]byte, 1000)

    // Read in data.
	count, _ := reader.Read(result)
	fmt.Println(count)
	//fmt.Println(string(result))
 
    err := ioutil.WriteFile("C:\\Users\\Miguel\\Documents\\repos\\wfire\\wfire\\src\\test_data\\dat1", result, 0644)
    check(err)
}

func ConvertBigEndian(){
	filename := "C:\\Users\\Miguel\\Documents\\repos\\wfire\\wfire\\src\\test_data\\N38W009.hgt"
	f, _ := os.Open(filename)
	reader, _ := gzip.NewReader(f)
	result := make([]byte, 2000)	
	
	count, _ := reader.Read(result)
	fmt.Println(count)
	converted := binary.BigEndian.Uint16(result)
	fmt.Println(converted)
 

} */

/*

def get_idx_in_file(self, pos):
	""" From a position (lat, lng) as float. Get the index of the elevation value inside the HGT file
	:param tuple pos: (lat, lng) of the position
	:return: tuple (index on the latitude from the top, index on the longitude from the left, index in the file)
	:rtype: (int, int, int)
	:raises Exception: if the point could not be found in the parsed HGT file
	"""
	if not self.is_inside(pos):
		raise Exception('point {} is not inside HGT file {}'.format(pos, self.filename))

	lat_idx = (self.sample_lat - 1) - int(round((pos[0] - self.bottom_left_center[0]) / self.square_height))
	lng_idx = int(round((pos[1] - self.bottom_left_center[1]) / self.square_width))
	idx = lat_idx * self.sample_lng + lng_idx
	return lat_idx, lng_idx, idx

*/