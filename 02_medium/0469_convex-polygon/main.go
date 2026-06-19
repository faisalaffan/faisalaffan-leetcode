package main

// LeetCode #469: Convex Polygon
// https://leetcode.com/problems/convex-polygon/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(ConvexPolygon([][]int{{0, 0}, {0, 1}, {1, 1}, {1, 0}}))
	fmt.Println(ConvexPolygon([][]int{{0, 0}, {0, 10}, {10, 10}, {10, 0}, {5, 5}}))
}

func ConvexPolygon(polygon [][]int) bool {
	n := len(polygon)
	if n < 3 {
		return false
	}

	var prevCross int
	first := true

	for i := 0; i < n; i++ {
		a, b, c := polygon[i], polygon[(i+1)%n], polygon[(i+2)%n]
		cross := (b[0]-a[0])*(c[1]-b[1]) - (b[1]-a[1])*(c[0]-b[0])
		if cross != 0 {
			if first {
				prevCross = cross
				first = false
			} else if (cross > 0 && prevCross < 0) || (cross < 0 && prevCross > 0) {
				return false
			}
		}
	}
	return true
}
