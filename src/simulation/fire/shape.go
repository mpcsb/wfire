package fire

import (
	"math"

	s "simulation/shared"
)

func Lerp(a, b int, i float64) float64 {
	return float64(a) + i*float64(b-a)
}

func Equidistant(a, b, n int) (points []float64) {
	for i := 0; i <= n; i++ {
		c := float64(1.0 / (n * i))
		points = append(points, Lerp(a, b, c))
	}
	return points
}

func Radians(d float64) float64 {
	return d * math.Pi / 180
}

func NewCoord_XY(p s.Coord, dx, dy float64) s.Coord {
	const r = 6371000
	new_latitude := p.Lat + (dy/r)*(180/math.Pi)
	new_longitude := p.Lon + (dx/r)*(180/math.Pi)/math.Cos(Radians(p.Lat))
	new_coord := s.Coord{Lat: new_latitude, Lon: new_longitude, Alt: p.Alt}
	return new_coord
}

func circle(n_points int) (perimeter [][]float64) {
	var p []float64
	for _, d := range Equidistant(0, 360, n_points) {
		p[0] = math.Cos(Radians(d))
		p[1] = math.Sin(Radians(d))
		perimeter = append(perimeter, p)
	}
	return perimeter
}

func parabola(n_points int, speed, alpha float64) (perimeter [][]float64) {
	var parabola [][]float64
	for _, p := range Equidistant(-2.0, 2.0, n_points) {
		parabola = append(parabola, []float64{float64(p), speed * (float64(math.Pow(2.0, 2) - math.Pow(p, 2.0)))})
	}

	// reflect and compress parabola to match unit circle
	for _, p := range parabola {
		x := p[0]*math.Cos(Radians(-alpha))*0.5 - p[1]*math.Sin(Radians(-alpha))*0.5
		y := p[0]*math.Sin(Radians(-alpha))*0.5 + p[1]*math.Cos(Radians(-alpha))*0.5

		perimeter = append(perimeter, []float64{x, y})
	}
	return perimeter
}

func Perimeter(n_points int, speed, alpha float64) (c, p [][]float64) {
	c = circle(n_points)
	p = parabola(n_points, speed, alpha)

	return c, p
}
