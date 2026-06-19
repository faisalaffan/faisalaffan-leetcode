package main

// LeetCode #2366: Minimum Replacements to Sort the Array
// https://leetcode.com/problems/minimum-replacements-to-sort-the-array/
// Difficulty: Hard
//
// You are given a 0-indexed integer array nums. In one operation you can replace
// any element with any number of elements that sum to it. Return the minimum
// number of operations to make the array non-decreasing.
//
// Approach: Greedy from right to left. Maintain a bound = last element.
// For each nums[i] > bound, we must split nums[i] into pieces each <= bound.
// The minimum pieces needed = ceil(nums[i] / bound). We add (pieces-1) operations.
// Then the new bound for the left side is floor(nums[i] / pieces) which is the
// largest possible value that keeps things non-decreasing.

import (
	"fmt"
)

func minimumReplacements(nums []int) int64 {
	n := len(nums)
	var ans int64
	bound := nums[n-1]

	for i := n - 2; i >= 0; i-- {
		if nums[i] > bound {
			// Minimum number of pieces each <= bound
			parts := (nums[i] + bound - 1) / bound
			ans += int64(parts - 1)
			// Largest possible value for the leftmost piece
			bound = nums[i] / parts
		} else {
			bound = nums[i]
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(minimumReplacements([]int{3, 9, 3}))
	// Example 2
	fmt.Println(minimumReplacements([]int{1, 2, 3, 4, 5}))
	// Edge: strictly decreasing
	fmt.Println(minimumReplacements([]int{5, 4, 3, 2, 1}))
	// Edge: single element
	fmt.Println(minimumReplacements([]int{1}))
	// Edge: large gap
	fmt.Println(minimumReplacements([]int{12, 9, 7, 6, 17, 19, 21}))
}
