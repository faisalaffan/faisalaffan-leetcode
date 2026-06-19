package main

// LeetCode #1401: Circle and Rectangle Overlapping
// https://leetcode.com/problems/circle-and-rectangle-overlapping/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(checkOverlap(1, 0, 0, 1, -1, 3, 1)) // true

	// Test case 2
	fmt.Println(checkOverlap(1, 1, 1, -3, -3, 3, 3)) // true

	// Test case 3
	fmt.Println(checkOverlap(1, 0, 0, -1, 0, 0, 1)) // false

	// Test case 4
	fmt.Println(checkOverlap(1, 1, 1, -3, -3, 3, -1)) // true
}

// Time: O(1)
// Space: O(1)
func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
	// Find the closest point on the rectangle to the circle center
	closestX := max(x1, min(x2, xCenter))
	closestY := max(y1, min(y2, yCenter))

	// Calculate distance from circle center to closest point
	dx := xCenter - closestX
	dy := yCenter - closestY

	return dx*dx+dy*dy <= radius*radius
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
