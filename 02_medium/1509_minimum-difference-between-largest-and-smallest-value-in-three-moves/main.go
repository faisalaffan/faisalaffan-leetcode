package main

// LeetCode #1509: Minimum Difference Between Largest and Smallest Value in Three Moves
// https://leetcode.com/problems/minimum-difference-between-largest-and-smallest-value-in-three-moves/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinDifference([]int{5, 3, 2, 4}))
	fmt.Println(MinDifference([]int{1, 5, 0, 10, 14}))
	fmt.Println(MinDifference([]int{3, 100, 20}))
}

func MinDifference(nums []int) int {
	// Time: O(N log N), Space: O(1) if ignoring sort space
	if len(nums) <= 4 {
		return 0
	}

	sort.Ints(nums)
	n := len(nums)

	// After 3 moves, we can change up to 3 values.
	// The minimum difference will be between some combination
	// of removing 0-3 from left and 3-0 from right.
	minDiff := nums[n-1] - nums[0]
	for i := 0; i <= 3; i++ {
		diff := nums[n-1-(3-i)] - nums[i]
		if diff < minDiff {
			minDiff = diff
		}
	}

	return minDiff
}
