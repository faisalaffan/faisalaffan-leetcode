package main

// LeetCode #945: Minimum Increment to Make Array Unique
// https://leetcode.com/problems/minimum-increment-to-make-array-unique/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

// Time: O(n log n) | Space: O(1) ignoring sort space
func minIncrementForUnique(nums []int) int {
	sort.Ints(nums)
	moves := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			moves += nums[i-1] + 1 - nums[i]
			nums[i] = nums[i-1] + 1
		}
	}
	return moves
}

func main() {
	fmt.Println(minIncrementForUnique([]int{1, 2, 2}))
	fmt.Println(minIncrementForUnique([]int{3, 2, 1, 2, 1, 7}))
}
