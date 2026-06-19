package main

// LeetCode #452: Minimum Number of Arrows to Burst Balloons
// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	if len(points) == 0 {
		return 0
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][1] < points[j][1]
	})

	arrows := 1
	end := points[0][1]

	for _, p := range points[1:] {
		if p[0] > end {
			arrows++
			end = p[1]
		}
	}
	return arrows
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", findMinArrowShots([][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", findMinArrowShots([][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}))
	// Expected: 2
}
