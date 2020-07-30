package fire

import (
	s "simulation/shared"
)

// https://www.geeksforgeeks.org/how-to-check-if-a-given-point-lies-inside-a-polygon/
// Given three colinear points p, q, r, the function checks if
// point q lies on line segment 'pr'
func onSegment(p, q, r s.Coord) bool {
	if q.Lat <= s.Max(p.Lat, r.Lat) && q.Lat >= s.Min(p.Lat, r.Lat) &&
		q.Lon <= s.Max(p.Lon, r.Lon) && q.Lon >= s.Min(p.Lon, r.Lon) {
		return true
	}
	return false
}

// To find orientation of ordered triplet (p, q, r).
// The function returns following values
// 0 --> p, q and r are colinear
// 1 --> Clockwise
// 2 --> Counterclockwise
func orientation(p, q, r s.Coord) int {
	val := (q.Lon-p.Lon)*(r.Lat-q.Lat) -
		(q.Lat-p.Lat)*(r.Lon-q.Lon)

	if val == 0 {
		return 0
	}
	if val > 0 {
		return 1
	} else {
		return 2
	}
}

func doIntersect(p1, q1, p2, q2 s.Coord) bool {
	// Find the four orientations needed for general and
	// special cases
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)

	// General case
	if o1 != o2 && o3 != o4 {
		return true
	}

	// Special Cases
	// p1, q1 and p2 are colinear and p2 lies on segment p1q1
	if o1 == 0 && onSegment(p1, p2, q1) {
		return true
	}

	// p1, q1 and p2 are colinear and q2 lies on segment p1q1
	if o2 == 0 && onSegment(p1, q2, q1) {
		return true
	}

	// p2, q2 and p1 are colinear and p1 lies on segment p2q2
	if o3 == 0 && onSegment(p2, p1, q2) {
		return true
	}

	// p2, q2 and q1 are colinear and q1 lies on segment p2q2
	if o4 == 0 && onSegment(p2, q1, q2) {
		return true
	}

	return false // Doesn't fall in any of the above cases
}

// IsInside determines if a point p is within a given polygon
func (f Flame) IsInside(contour []s.Coord, p s.Coord) bool {
	// There must be at least 3 vertices in polygon[]
	n := len(contour)

	// Create a point for line segment from p to infinite
	extreme := s.Coord{Lat: 90, Lon: p.Lon}

	// Count intersections of the above line with sides of polygon
	count := 0
	for i := range contour {
		next := (i + 1) % n
		// Check if the line segment from 'p' to 'extreme' intersects
		// with the line segment from 'polygon[i]' to 'polygon[next]'
		if doIntersect(contour[i], contour[next], p, extreme) {
			// If the point 'p' is colinear with line segment 'i-next',
			// then check if it lies on segment. If it lies, return true,
			// otherwise false
			if orientation(contour[i], p, contour[next]) == 0 {
				return onSegment(contour[i], p, contour[next])
			}
			count++
		}
		i = next
	}

	// Return true if count is odd, false otherwise
	if count%2 == 1 {
		return true
	} else {
		return false
	}
}
