package fire

import (
	"math"

	s "simulation/shared"
)

// Linear interpolation
func lerp(a, b int, i float64) float64 {
	return float64(a) + i*float64(b-a)
}

func equidistant(a, b, n int) (points []float64) {
	for i := 0; i <= n; i++ {
		c := float64(1.0 / (n * i))
		points = append(points, lerp(a, b, c))
	}
	return points
}

// NewCoordXY receive a coordinate and deltas in X and Y and returns the corresponding coordinate
func NewCoordXY(p s.Coord, dx, dy float64) s.Coord {
	const r = 6371000
	newLatitude := p.Lat + (dy/r)*(180/math.Pi)
	newLongitude := p.Lon + (dx/r)*(180/math.Pi)/math.Cos(s.Radians(p.Lat))
	newCoord := s.Coord{Lat: newLatitude, Lon: newLongitude, Alt: p.Alt}
	return newCoord
}

func circle(nPoints int) (perimeter [][]float64) {
	var p []float64
	for _, d := range equidistant(0, 360, nPoints) {
		p[0] = math.Cos(s.Radians(d))
		p[1] = math.Sin(s.Radians(d))
		perimeter = append(perimeter, p)
	}
	return perimeter
}

func parabola(nPoints int, speed, alpha float64) (perimeter [][]float64) {
	var parabola [][]float64
	for _, p := range equidistant(-2.0, 2.0, nPoints) {
		parabola = append(parabola, []float64{float64(p), speed * (float64(math.Pow(2.0, 2) - math.Pow(p, 2.0)))})
	}

	// reflect and compress parabola to match unit circle
	for _, p := range parabola {
		x := p[0]*math.Cos(s.Radians(-alpha))*0.5 - p[1]*math.Sin(s.Radians(-alpha))*0.5
		y := p[0]*math.Sin(s.Radians(-alpha))*0.5 + p[1]*math.Cos(s.Radians(-alpha))*0.5

		perimeter = append(perimeter, []float64{x, y})
	}
	return perimeter
}

// Perimeter returns the contours from the circle containing the tree and the parabola generated from the wind
func Perimeter(nPoints int, speed, alpha float64) (c, p [][]float64) {
	c = circle(nPoints)
	p = parabola(nPoints, speed, alpha)

	return c, p
}
