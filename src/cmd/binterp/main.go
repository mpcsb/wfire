package main

import "fmt"

type point struct {
	x float64
	y float64
}

func f(x, y float64) float64 {
	return 22*x + 3*y + 2*x*x
}

func binterp(x1, x2, y1, y2 float64, p point) float64 {
	x := p.x
	y := p.y

	R1 := f(x1, y1) + (x-x1)/(x2-x1)*(f(x2, y1)-f(x1, y1))
	R2 := f(x1, y2) + (x-x1)/(x2-x1)*(f(x2, y2)-f(x1, y2))

	val := R2 + (y-y2)/(y2-y1)*(R1-R2)
	return val
}

func main() {
	X := 34.234
	Y := 9.434
	P := point{x: X, y: Y}

	x1 := 32.123
	y1 := 9.1234
	x2 := 38.456
	y2 := 9.3456

	a := binterp(x1, x2, y1, y2, P)
	fmt.Println(a, f(X, Y))
}
