package main

// LeetCode #3364: Minimum Positive Sum Subarray
// https://leetcode.com/problems/minimum-positive-sum-subarray/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumPositiveSumSubarray([]int{3, -2, 1, 4}, 2, 3))
	fmt.Println(MinimumPositiveSumSubarray([]int{-2, 2, -2, 2}, 1, 2))
	fmt.Println(MinimumPositiveSumSubarray([]int{1, 2, 3, 4}, 2, 4))
}

// MinimumPositiveSumSubarray returns the minimum positive sum of any subarray with length between l and r.
// Time: O(n * (r-l+1)). Space: O(1).
func MinimumPositiveSumSubarray(nums []int, l int, r int) int {
	n := len(nums)
	minPos := -1
	for length := l; length <= r; length++ {
		for start := 0; start <= n-length; start++ {
			sum := 0
			for i := start; i < start+length; i++ {
				sum += nums[i]
			}
			if sum > 0 && (minPos == -1 || sum < minPos) {
				minPos = sum
			}
		}
	}
	return minPos
}
