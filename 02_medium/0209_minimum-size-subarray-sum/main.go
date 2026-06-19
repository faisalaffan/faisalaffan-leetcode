package main

// LeetCode #209: Minimum Size Subarray Sum
// https://leetcode.com/problems/minimum-size-subarray-sum/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import (
	"fmt"
	"math"
)

func minSubArrayLen(target int, nums []int) int {
	left, sum := 0, 0
	minLen := math.MaxInt32

	for right := 0; right < len(nums); right++ {
		sum += nums[right]

		for sum >= target {
			if right-left+1 < minLen {
				minLen = right - left + 1
			}
			sum -= nums[left]
			left++
		}
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func main() {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
	fmt.Println(minSubArrayLen(4, []int{1, 4, 4}))
	fmt.Println(minSubArrayLen(11, []int{1, 1, 1, 1, 1, 1, 1, 1}))
}
