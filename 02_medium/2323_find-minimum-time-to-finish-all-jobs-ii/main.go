package main

// LeetCode #2323: Find Minimum Time to Finish All Jobs II
// https://leetcode.com/problems/find-minimum-time-to-finish-all-jobs-ii/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimumTime(jobs []int, workers []int) int {
	sort.Ints(jobs)
	sort.Ints(workers)
	maxDays := 0

	for i := 0; i < len(jobs); i++ {
		days := (jobs[i] + workers[i] - 1) / workers[i] // ceil division
		if days > maxDays {
			maxDays = days
		}
	}
	return maxDays
}

func main() {
	// Test case 1
	fmt.Println(minimumTime([]int{5, 2, 4}, []int{1, 7, 5}))
	// Expected: 2

	// Test case 2
	fmt.Println(minimumTime([]int{3, 18, 30}, []int{3, 15, 5}))
	// Expected: 6
}
