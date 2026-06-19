package main

// LeetCode #643: Maximum Average Subarray I
// https://leetcode.com/problems/maximum-average-subarray-i/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findMaxAverage([]int{1, 12, -5, -6, 50, 3}, 4)) // 12.75
	fmt.Println(findMaxAverage([]int{5}, 1))                     // 5.0
	fmt.Println(findMaxAverage([]int{-1}, 1))                    // -1.0
}

// findMaxAverage finds a contiguous subarray of length k that has the maximum average value.
// Time: O(n). Space: O(1).
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum += nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}
	return float64(maxSum) / float64(k)
}
