package main

// LeetCode #1459: Rectangles Area
// https://leetcode.com/problems/rectangles-area/
// Difficulty: Medium (listed here as Hard) [Paid]
//
// Given points (x, y) in a 2D plane, find all rectangles formed by these points.
// A rectangle has sides parallel to the axes. Return the area and the two
// opposite corner points for each rectangle.

import (
	"fmt"
	"sort"
)

// Point represents a 2D coordinate.
type Point struct {
	X, Y int
}

// Rectangle represents a rectangle with two opposite corners.
type Rectangle struct {
	Area     int
	P1, P2   Point
}

// findRectangles finds all axis-aligned rectangles from given points.
func findRectangles(points []Point) []Rectangle {
	// Build a set of points for O(1) lookup
	pointSet := make(map[Point]bool)
	for _, p := range points {
		pointSet[p] = true
	}

	var rects []Rectangle
	seen := make(map[string]bool)

	// For each pair of points, check if they can be opposite corners
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			p1, p2 := points[i], points[j]

			// They must be diagonal corners (different x AND different y)
			if p1.X == p2.X || p1.Y == p2.Y {
				continue
			}

			// The other two corners must exist
			c3 := Point{p1.X, p2.Y}
			c4 := Point{p2.X, p1.Y}

			if pointSet[c3] && pointSet[c4] {
				// Ensure we don't add the same rectangle twice
				// by normalizing: smaller x first
				rectP1, rectP2 := p1, p2
				if rectP1.X > rectP2.X || (rectP1.X == rectP2.X && rectP1.Y > rectP2.Y) {
					rectP1, rectP2 = rectP2, rectP1
				}
				key := fmt.Sprintf("%d,%d-%d,%d", rectP1.X, rectP1.Y, rectP2.X, rectP2.Y)
				if seen[key] {
					continue
				}
				seen[key] = true

				area := abs(p2.X-p1.X) * abs(p2.Y-p1.Y)
				if area > 0 {
					rects = append(rects, Rectangle{area, rectP1, rectP2})
				}
			}
		}
	}

	// Sort by area descending, then by p1.x, p1.y, p2.x, p2.y
	sort.Slice(rects, func(i, j int) bool {
		if rects[i].Area != rects[j].Area {
			return rects[i].Area > rects[j].Area
		}
		if rects[i].P1.X != rects[j].P1.X {
			return rects[i].P1.X < rects[j].P1.X
		}
		if rects[i].P1.Y != rects[j].P1.Y {
			return rects[i].P1.Y < rects[j].P1.Y
		}
		if rects[i].P2.X != rects[j].P2.X {
			return rects[i].P2.X < rects[j].P2.X
		}
		return rects[i].P2.Y < rects[j].P2.Y
	})

	return rects
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	// Test case 1: Simple rectangle
	points := []Point{
		{0, 0}, {0, 2}, {2, 0}, {2, 2},
	}
	rects := findRectangles(points)
	fmt.Println("Test 1 - Rectangles:")
	for _, r := range rects {
		fmt.Printf("  Area=%d, (%d,%d)-(%d,%d)\n", r.Area, r.P1.X, r.P1.Y, r.P2.X, r.P2.Y)
	}

	// Test case 2: Two overlapping rectangles
	points2 := []Point{
		{0, 0}, {0, 2}, {2, 0}, {2, 2},
		{1, 0}, {1, 2},
	}
	rects2 := findRectangles(points2)
	fmt.Println("\nTest 2 - Rectangles (overlapping):")
	for _, r := range rects2 {
		fmt.Printf("  Area=%d, (%d,%d)-(%d,%d)\n", r.Area, r.P1.X, r.P1.Y, r.P2.X, r.P2.Y)
	}

	// Test case 3: No rectangles (collinear)
	points3 := []Point{
		{0, 0}, {1, 1}, {2, 2},
	}
	rects3 := findRectangles(points3)
	fmt.Printf("\nTest 3 - Rectangles (collinear): %d (expected 0)\n", len(rects3))

	// Test case 4: Multiple rectangles with different areas
	points4 := []Point{
		{0, 0}, {0, 3}, {3, 0}, {3, 3},
		{1, 1}, {1, 2}, {2, 1}, {2, 2},
	}
	rects4 := findRectangles(points4)
	fmt.Println("\nTest 4 - Rectangles (nested):")
	for _, r := range rects4 {
		fmt.Printf("  Area=%d, (%d,%d)-(%d,%d)\n", r.Area, r.P1.X, r.P1.Y, r.P2.X, r.P2.Y)
	}

	// Test case 5: Single point
	points5 := []Point{{5, 5}}
	rects5 := findRectangles(points5)
	fmt.Printf("\nTest 5 - Rectangles (single point): %d (expected 0)\n", len(rects5))
}
