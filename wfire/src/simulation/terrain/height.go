package terrain

import (
	//"fmt"
	"strconv"  
	//"Strings"
)

type geo_coord struct {
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
