package main

// LeetCode #2931: Maximum Spending After Buying Items
// https://leetcode.com/problems/maximum-spending-after-buying-items/
// Difficulty: Hard
//
// Approach: Sort all values ascending.
// By the rearrangement inequality, to maximize sum(day * value),
// the smallest values should be bought on the earliest days.
// Since each row is sorted ascending, merging all values into one
// sorted list and pairing with day number gives the optimal result.

import (
	"fmt"
	"sort"
)

func maxSpending(values [][]int) int64 {
	m := len(values)
	if m == 0 {
		return 0
	}
	n := len(values[0])

	// Flatten all values into one slice
	flat := make([]int, 0, m*n)
	for _, row := range values {
		flat = append(flat, row...)
	}
	sort.Ints(flat)

	var ans int64
	for day, val := range flat {
		ans += int64(day+1) * int64(val)
	}
	return ans
}

func main() {
	// Example: values=[[8,5,2],[6,4,1],[9,7,3]] -> 285
	// Sorted: [1,2,3,4,5,6,7,8,9]
	// Spending: 1*1 + 2*2 + 3*3 + 4*4 + 5*5 + 6*6 + 7*7 + 8*8 + 9*9 = 285
	fmt.Println(maxSpending([][]int{{8, 5, 2}, {6, 4, 1}, {9, 7, 3}}))

	// Single row
	fmt.Println(maxSpending([][]int{{1, 2, 3}}))

	// Simple case
	fmt.Println(maxSpending([][]int{{10, 20}, {5, 15}}))
}
