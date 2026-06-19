package main

// LeetCode #1779: Find Nearest Point That Has the Same X or Y Coordinate
// https://leetcode.com/problems/find-nearest-point-that-has-the-same-x-or-y-coordinate/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func NearestValidPoint(x int, y int, points [][]int) int {
	minDist := -1
	bestIdx := -1
	for i, p := range points {
		if p[0] == x || p[1] == y {
			dist := abs(p[0]-x) + abs(p[1]-y)
			if minDist == -1 || dist < minDist {
				minDist = dist
				bestIdx = i
			}
		}
	}
	return bestIdx
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	fmt.Println(NearestValidPoint(3, 4, [][]int{{1, 2}, {3, 1}, {2, 4}, {2, 3}, {4, 4}}))
	fmt.Println(NearestValidPoint(3, 4, [][]int{{3, 4}}))
	fmt.Println(NearestValidPoint(3, 4, [][]int{{2, 3}}))
}
