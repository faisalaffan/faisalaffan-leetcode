package main

// LeetCode #795: Number of Subarrays with Bounded Maximum
// https://leetcode.com/problems/number-of-subarrays-with-bounded-maximum/
// Difficulty: Medium
// Time: O(n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(numSubarrayBoundedMax([]int{2, 1, 4, 3}, 2, 3))
	fmt.Println(numSubarrayBoundedMax([]int{2, 9, 2, 5, 6}, 2, 8))
}

func numSubarrayBoundedMax(nums []int, left int, right int) int {
	count := 0
	prevCount := 0
	prevLessIdx := -1

	for i, val := range nums {
		if val >= left && val <= right {
			prevCount = i - prevLessIdx
			count += prevCount
		} else if val < left {
			count += prevCount
		} else {
			prevCount = 0
			prevLessIdx = i
		}
	}

	return count
}
