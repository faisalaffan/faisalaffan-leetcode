package main

// LeetCode #223: Rectangle Area
// https://leetcode.com/problems/rectangle-area/
// Difficulty: Medium
// Time: O(1), Space: O(1)

import "fmt"

func computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2 int) int {
	areaA := (ax2 - ax1) * (ay2 - ay1)
	areaB := (bx2 - bx1) * (by2 - by1)

	overlapX := max(0, min(ax2, bx2)-max(ax1, bx1))
	overlapY := max(0, min(ay2, by2)-max(ay1, by1))

	return areaA + areaB - overlapX*overlapY
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(computeArea(-3, 0, 3, 4, 0, -1, 9, 2))
	fmt.Println(computeArea(-2, -2, 2, 2, -2, -2, 2, 2))
	fmt.Println(computeArea(0, 0, 0, 0, -1, -1, 1, 1))
}
