package main

// LeetCode #2033: Minimum Operations to Make a Uni-Value Grid
// https://leetcode.com/problems/minimum-operations-to-make-a-uni-value-grid/
// Difficulty: Medium
// Time: O(m*n log(m*n)) | Space: O(m*n)

import (
	"fmt"
	"sort"
)

func minOperations(grid [][]int, x int) int {
	m, n := len(grid), len(grid[0])
	vals := make([]int, 0, m*n)

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			vals = append(vals, grid[i][j])
		}
	}

	// Check all values have same remainder mod x
	rem := vals[0] % x
	for _, v := range vals {
		if v%x != rem {
			return -1
		}
	}

	sort.Ints(vals)
	median := vals[len(vals)/2]

	ops := 0
	for _, v := range vals {
		diff := v - median
		if diff < 0 {
			diff = -diff
		}
		ops += diff / x
	}
	return ops
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minOperations([][]int{{2, 4}, {6, 8}}, 2))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", minOperations([][]int{{1, 5}, {2, 3}}, 1))
	// Expected: 5

	// Test case 3
	fmt.Println("Test 3:", minOperations([][]int{{1, 2}, {3, 4}}, 2))
	// Expected: -1
}
