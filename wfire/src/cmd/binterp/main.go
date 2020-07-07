package main

import "fmt"

type point struct {
	x float64
	y float64
}

func f(x, y float64) float64 {
	return 55*x + y
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
	X := 38.234
	Y := 9.434
	P := point{x: X, y: Y}

	x1 := 38.123
	y1 := 9.1234
	x2 := 38.456
	y2 := 9.3456

	a := binterp(x1, x2, y1, y2, P)
	b := binterp(y1, y2, x1, x2, P)
	
	c := binterp(y1, x2, y1, x2, P)
	d := binterp(x1, y2, y1, x2, P)

	fmt.Println(f(x1, y1))
	fmt.Println(f(x2, y2))
	fmt.Println(f(x1, y2))
	fmt.Println(f(x2, y1))

	fmt.Println(a, b, c, d)
	fmt.Println((a + b) * 0.5)
	fmt.Println(f(X, Y))
}
