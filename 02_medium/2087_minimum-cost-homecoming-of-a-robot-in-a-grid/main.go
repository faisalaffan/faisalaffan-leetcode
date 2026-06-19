package main

// LeetCode #2087: Minimum Cost Homecoming of a Robot in a Grid
// https://leetcode.com/problems/minimum-cost-homecoming-of-a-robot-in-a-grid/
// Difficulty: Medium
// Time: O(m + n) | Space: O(1)

import "fmt"

func minCost(startPos []int, homePos []int, rowCosts []int, colCosts []int) int {
	cost := 0
	r1, c1 := startPos[0], startPos[1]
	r2, c2 := homePos[0], homePos[1]

	// Move rows
	step := 1
	if r1 > r2 {
		step = -1
	}
	for r := r1 + step; r != r2+step; r += step {
		cost += rowCosts[r]
	}

	// Move columns
	step = 1
	if c1 > c2 {
		step = -1
	}
	for c := c1 + step; c != c2+step; c += step {
		cost += colCosts[c]
	}

	return cost
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minCost([]int{1, 0}, []int{2, 3}, []int{5, 4, 3}, []int{8, 2, 6, 7}))
	// Expected: 18

	// Test case 2
	fmt.Println("Test 2:", minCost([]int{0, 0}, []int{0, 0}, []int{5}, []int{5}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", minCost([]int{2, 2}, []int{0, 0}, []int{1, 2, 3}, []int{4, 5, 6}))
	// Expected: 15 (row 1 + row 0 + col 1 + col 0 = 2+1+5+4 = 12... let me recalculate)
	// Moving from row 2 to 0: rowCosts[1] + rowCosts[0] = 2 + 1 = 3
	// Moving from col 2 to 0: colCosts[1] + colCosts[0] = 5 + 4 = 9
	// Total = 12
}
