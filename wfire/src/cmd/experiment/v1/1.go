package main

import (
	"fmt"
	"encoding/binary"
	"os"
)

func main() {
    buf, _ := os.Open("C:\\Users\\Miguel\\Documents\\repos\\wfire\\wfire\\src\\test_data\\N38W009.hgt")
    var dataOut int
    err := binary.Read(buf, binary.BigEndian, &dataOut)
    if err != nil {
        fmt.Println("failed to Read:", err)
        return 
    }
}