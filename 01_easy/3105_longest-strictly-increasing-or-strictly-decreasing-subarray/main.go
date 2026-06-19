package main

// LeetCode #3105: Longest Strictly Increasing or Strictly Decreasing Subarray
// https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: longestMonotonicSubarray
	fmt.Println(LongestStrictlyIncreasingOrStrictlyDecreasingSubarray([]int{1, 4, 3, 3, 2})) // 2
	fmt.Println(LongestStrictlyIncreasingOrStrictlyDecreasingSubarray([]int{3, 3, 3, 3}))  // 1
	fmt.Println(LongestStrictlyIncreasingOrStrictlyDecreasingSubarray([]int{3, 2, 1}))     // 3
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: longestMonotonicSubarray
func LongestStrictlyIncreasingOrStrictlyDecreasingSubarray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	inc := 1
	dec := 1
	maxLen := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			inc++
			dec = 1
		} else if nums[i] < nums[i-1] {
			dec++
			inc = 1
		} else {
			inc = 1
			dec = 1
		}
		if inc > maxLen {
			maxLen = inc
		}
		if dec > maxLen {
			maxLen = dec
		}
	}
	return maxLen
}
