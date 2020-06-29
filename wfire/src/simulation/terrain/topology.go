package terrain


type cart_coord struct {
	x float64
	y float64
	alt int
}

type geo_coord struct {
	lat float64
	lon float64
	alt int
}


type terrain struct{
	coords [] geo_coord
} 