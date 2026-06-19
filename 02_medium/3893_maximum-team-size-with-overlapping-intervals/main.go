package main

// LeetCode #3893: Maximum Team Size with Overlapping Intervals
// https://leetcode.com/problems/maximum-team-size-with-overlapping-intervals/
// Difficulty: Medium [Paid]
// Time: O(N log N) | Space: O(N)
// Approach: For each employee, count overlapping intervals using binary search
// on sorted start and end times.

import (
	"fmt"
	"sort"
)

func MaximumTeamSizeWithOverlappingIntervals(startTime []int, endTime []int) int {
	n := len(startTime)
	st := make([]int, n)
	et := make([]int, n)
	copy(st, startTime)
	copy(et, endTime)
	sort.Ints(st)
	sort.Ints(et)

	ans := 0
	for i := 0; i < n; i++ {
		start := startTime[i]
		end := endTime[i]

		// Count employees whose start <= end
		startsBeforeEnd := sort.SearchInts(st, end+1)
		// Count employees whose end < start
		endsBeforeStart := sort.SearchInts(et, start)

		overlap := startsBeforeEnd - endsBeforeStart
		if overlap > ans {
			ans = overlap
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MaximumTeamSizeWithOverlappingIntervals([]int{1, 2, 3}, []int{4, 5, 6})) // Expected: 3

	// Example 2
	fmt.Println(MaximumTeamSizeWithOverlappingIntervals([]int{2, 5, 8}, []int{3, 7, 9})) // Expected: 1

	// Example 3
	fmt.Println(MaximumTeamSizeWithOverlappingIntervals([]int{3, 4, 6}, []int{8, 5, 7})) // Expected: 3
}
