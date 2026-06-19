package main

// LeetCode #2159: Order Two Columns Independently
// https://leetcode.com/problems/order-two-columns-independently/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

type Row struct {
	col1, col2 int
}

func orderColumns(rows [][]int) [][]int {
	// Sort by col1 ascending, col2 ascending
	sort.Slice(rows, func(i, j int) bool {
		if rows[i][0] != rows[j][0] {
			return rows[i][0] < rows[j][0]
		}
		return rows[i][1] < rows[j][1]
	})

	return rows
}

func main() {
	// Test case 1
	data1 := [][]int{{3, 1}, {1, 3}, {2, 2}}
	fmt.Println("Test 1:", orderColumns(data1))
	// Expected: [[1,3],[2,2],[3,1]]

	// Test case 2
	data2 := [][]int{{5, 5}, {1, 1}, {3, 3}}
	fmt.Println("Test 2:", orderColumns(data2))
	// Expected: [[1,1],[3,3],[5,5]]
}
