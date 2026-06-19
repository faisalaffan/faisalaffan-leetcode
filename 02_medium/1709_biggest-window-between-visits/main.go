package main

// LeetCode #1709: Biggest Window Between Visits
// https://leetcode.com/problems/biggest-window-between-visits/
// Difficulty: Medium [Paid] (SQL)
// This is a SQL problem. In Go we simulate the logic.
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

func biggestWindow(visits []int) int {
	if len(visits) == 0 {
		return 0
	}
	sort.Ints(visits)
	maxGap := 0
	for i := 1; i < len(visits); i++ {
		gap := visits[i] - visits[i-1]
		if gap > maxGap {
			maxGap = gap
		}
	}
	return maxGap
}

func main() {
	fmt.Println(biggestWindow([]int{1, 3, 7, 10})) // Expected: 4 (between 3 and 7)
	fmt.Println(biggestWindow([]int{1, 2, 3, 4}))  // Expected: 1
	fmt.Println(biggestWindow([]int{5}))            // Expected: 0
}
