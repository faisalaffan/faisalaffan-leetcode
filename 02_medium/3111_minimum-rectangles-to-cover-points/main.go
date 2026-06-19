package main

// LeetCode #3111: Minimum Rectangles to Cover Points
// https://leetcode.com/problems/minimum-rectangles-to-cover-points/
// Difficulty: Medium
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func minRectanglesToCoverPoints(points [][]int, w int) int {
	xs := make([]int, len(points))
	for i, p := range points {
		xs[i] = p[0]
	}
	sort.Ints(xs)

	ans := 0
	i := 0
	for i < len(xs) {
		ans++
		end := xs[i] + w
		for i < len(xs) && xs[i] <= end {
			i++
		}
	}
	return ans
}

func main() {
	fmt.Println(minRectanglesToCoverPoints([][]int{{2, 1}, {1, 0}, {1, 4}, {1, 8}, {3, 5}, {4, 6}}, 1)) // Expected: 2
	fmt.Println(minRectanglesToCoverPoints([][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}}, 2)) // Expected: 2
}
