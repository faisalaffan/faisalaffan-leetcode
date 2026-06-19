package main

// LeetCode #2401: Longest Nice Subarray
// https://leetcode.com/problems/longest-nice-subarray/
// Difficulty: Medium
// Time: O(n) | Space: O(1)
// Sliding window: maintain OR of current window. If new element conflicts, shrink left.

import "fmt"

func main() {
	fmt.Println(longestNiceSubarray([]int{1, 3, 8, 48, 10})) // 3
	fmt.Println(longestNiceSubarray([]int{3, 1, 5, 11, 13}))  // 1
}

func longestNiceSubarray(nums []int) int {
	left, orMask, ans := 0, 0, 0
	for right, v := range nums {
		for orMask&v != 0 {
			orMask ^= nums[left]
			left++
		}
		orMask |= v
		if right-left+1 > ans {
			ans = right - left + 1
		}
	}
	return ans
}
