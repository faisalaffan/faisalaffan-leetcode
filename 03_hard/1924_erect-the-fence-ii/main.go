package main

// LeetCode #1924: Erect the Fence II (Minimum Enclosing Circle)
// https://leetcode.com/problems/erect-the-fence-ii/
// Difficulty: Hard [Paid]
//
// Given a set of points (trees) on a 2D plane, find the minimum-radius circle
// that encloses all points. Return [center_x, center_y, radius].
//
// Uses Welzl's randomized algorithm for minimum enclosing circle,
// expected O(n) time.

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	// Test case 1: two points
	fmt.Println(erectTheFenceIi([][]int{{0, 0}, {2, 0}}))
	// Test case 2: three points forming a triangle
	fmt.Println(erectTheFenceIi([][]int{{0, 0}, {1, 0}, {0, 1}}))
	// Test case 3: single point
	fmt.Println(erectTheFenceIi([][]int{{5, 5}}))
	// Test case 4: colinear points
	fmt.Println(erectTheFenceIi([][]int{{0, 0}, {2, 0}, {4, 0}}))
	// Test case 5: square
	fmt.Println(erectTheFenceIi([][]int{{0, 0}, {2, 0}, {2, 2}, {0, 2}}))
	// Test case 6: random points
	fmt.Println(erectTheFenceIi([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
}

// Point represents a 2D point
type Point struct {
	x, y float64
}

// Circle represents a circle with center (x, y) and radius r
type Circle struct {
	x, y, r float64
}

func erectTheFenceIi(trees [][]int) []float64 {
	n := len(trees)
	if n == 0 {
		return []float64{0, 0, 0}
	}

	// Convert to float64 points
	pts := make([]Point, n)
	for i, t := range trees {
		pts[i] = Point{float64(t[0]), float64(t[1])}
	}

	// Shuffle for expected linear time
	rand.Shuffle(n, func(i, j int) {
		pts[i], pts[j] = pts[j], pts[i]
	})

	c := welzl(pts, nil, 0)
	return []float64{c.x, c.y, c.r}
}

func welzl(pts []Point, boundary []Point, idx int) Circle {
	if idx == len(pts) || len(boundary) == 3 {
		return trivialCircle(boundary)
	}

	c := welzl(pts, boundary, idx+1)
	if inside(c, pts[idx]) {
		return c
	}

	// pts[idx] must be on the boundary
	newBoundary := append(boundary, pts[idx])
	return welzl(pts, newBoundary, idx+1)
}

func trivialCircle(pts []Point) Circle {
	switch len(pts) {
	case 0:
		return Circle{0, 0, 0}
	case 1:
		return Circle{pts[0].x, pts[0].y, 0}
	case 2:
		return circleFromTwo(pts[0], pts[1])
	case 3:
		// Try circle through all three points
		c := circleFromThree(pts[0], pts[1], pts[2])
		// If it's valid (not NaN), return it
		if !math.IsNaN(c.r) && c.r >= 0 {
			return c
		}
		// Otherwise, return smallest circle from any two
		c12 := circleFromTwo(pts[0], pts[1])
		c13 := circleFromTwo(pts[0], pts[2])
		c23 := circleFromTwo(pts[1], pts[2])
		best := c12
		if c13.r < best.r {
			best = c13
		}
		if c23.r < best.r {
			best = c23
		}
		return best
	}
	return Circle{0, 0, 0}
}

func circleFromTwo(a, b Point) Circle {
	cx := (a.x + b.x) / 2
	cy := (a.y + b.y) / 2
	r := dist(a, b) / 2
	return Circle{cx, cy, r}
}

func circleFromThree(a, b, c Point) Circle {
	// Compute circumcenter of triangle abc
	// Using formula from https://en.wikipedia.org/wiki/Circumcircle
	d := 2 * (a.x*(b.y-c.y) + b.x*(c.y-a.y) + c.x*(a.y-b.y))
	if math.Abs(d) < 1e-12 {
		return Circle{0, 0, -1} // colinear, invalid
	}

	ux := ((a.x*a.x+a.y*a.y)*(b.y-c.y) + (b.x*b.x+b.y*b.y)*(c.y-a.y) + (c.x*c.x+c.y*c.y)*(a.y-b.y)) / d
	uy := ((a.x*a.x+a.y*a.y)*(c.x-b.x) + (b.x*b.x+b.y*b.y)*(a.x-c.x) + (c.x*c.x+c.y*c.y)*(b.x-a.x)) / d

	r := dist(Point{ux, uy}, a)
	return Circle{ux, uy, r}
}

func inside(c Circle, p Point) bool {
	dx := p.x - c.x
	dy := p.y - c.y
	return dx*dx+dy*dy <= c.r*c.r+1e-9
}

func dist(a, b Point) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	return math.Sqrt(dx*dx + dy*dy)
}
