package main

// LeetCode #1727: Largest Submatrix With Rearrangements
// https://leetcode.com/problems/largest-submatrix-with-rearrangements/
// Difficulty: Medium
// Time: O(m * n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func largestSubmatrix(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
	maxArea := 0

	heights := make([]int, n)

	for r := 0; r < m; r++ {
		// Update heights
		for c := 0; c < n; c++ {
			if matrix[r][c] == 1 {
				heights[c]++
			} else {
				heights[c] = 0
			}
		}

		// Sort heights for this row (to find max rectangle that can be formed
		// by rearranging columns)
		sorted := make([]int, n)
		copy(sorted, heights)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] > sorted[j]
		})

		// For each column, area = height * (col index) because it's sorted
		for c := 0; c < n; c++ {
			area := sorted[c] * (c + 1)
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func main() {
	fmt.Println(largestSubmatrix([][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}})) // Expected: 4
	fmt.Println(largestSubmatrix([][]int{{1, 0, 1, 0, 1}})) // Expected: 3
	fmt.Println(largestSubmatrix([][]int{{1, 1, 0}, {1, 0, 1}})) // Expected: 2
}
