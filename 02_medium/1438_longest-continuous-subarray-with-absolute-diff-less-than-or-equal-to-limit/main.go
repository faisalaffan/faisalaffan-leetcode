package main

// LeetCode #1438: Longest Continuous Subarray With Absolute Diff Less Than or Equal to Limit
// https://leetcode.com/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/
// Difficulty: Medium

import "fmt"

func main() {
	// Test case 1
	fmt.Println(longestSubarray([]int{8, 2, 4, 7}, 4)) // 2

	// Test case 2
	fmt.Println(longestSubarray([]int{10, 1, 2, 4, 7, 2}, 5)) // 4

	// Test case 3
	fmt.Println(longestSubarray([]int{4, 2, 2, 2, 4, 4, 2, 2}, 0)) // 3

	// Test case 4
	fmt.Println(longestSubarray([]int{1, 5, 6, 7, 8, 10, 6, 5, 6}, 4)) // 5
}

// Time: O(n) where n = len(nums)
// Space: O(n) for deques
func longestSubarray(nums []int, limit int) int {
	// Monotonic deques for tracking min and max in current window
	minDeque := make([]int, 0) // increasing
	maxDeque := make([]int, 0) // decreasing

	left := 0
	maxLen := 0

	for right := 0; right < len(nums); right++ {
		// Maintain min deque (increasing)
		for len(minDeque) > 0 && minDeque[len(minDeque)-1] > nums[right] {
			minDeque = minDeque[:len(minDeque)-1]
		}
		minDeque = append(minDeque, nums[right])

		// Maintain max deque (decreasing)
		for len(maxDeque) > 0 && maxDeque[len(maxDeque)-1] < nums[right] {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, nums[right])

		// Shrink window if diff > limit
		for maxDeque[0]-minDeque[0] > limit {
			if nums[left] == minDeque[0] {
				minDeque = minDeque[1:]
			}
			if nums[left] == maxDeque[0] {
				maxDeque = maxDeque[1:]
			}
			left++
		}

		length := right - left + 1
		if length > maxLen {
			maxLen = length
		}
	}

	return maxLen
}
