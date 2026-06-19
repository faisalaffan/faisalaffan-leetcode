package main

// LeetCode #1037: Valid Boomerang
// https://leetcode.com/problems/valid-boomerang/
// Difficulty: Easy
// Time: O(1) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 3}, {3, 2}})) // true
	fmt.Println(isBoomerang([][]int{{1, 1}, {2, 2}, {3, 3}})) // false
	fmt.Println(isBoomerang([][]int{{0, 0}, {0, 2}, {0, 1}})) // false
}

// LeetCode submission: isBoomerang
func isBoomerang(points [][]int) bool {
	x1, y1 := points[0][0], points[0][1]
	x2, y2 := points[1][0], points[1][1]
	x3, y3 := points[2][0], points[2][1]
	// Check area of triangle: (x2-x1)*(y3-y1) != (x3-x1)*(y2-y1)
	return (x2-x1)*(y3-y1) != (x3-x1)*(y2-y1)
}
