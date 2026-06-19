package main

// LeetCode #3689: Maximum Total Subarray Value I
// https://leetcode.com/problems/maximum-total-subarray-value-i/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"
import "math"

func maximumTotalSubarrayValueI(nums []int, k int) int64 {
	minVal := math.MaxInt32
	maxVal := math.MinInt32
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}
	return int64(maxVal-minVal) * int64(k)
}

func main() {
	fmt.Println(maximumTotalSubarrayValueI([]int{1, 3, 2}, 2))
	fmt.Println(maximumTotalSubarrayValueI([]int{4, 2, 5, 1}, 3))
	fmt.Println(maximumTotalSubarrayValueI([]int{1, 1, 1}, 5))
}
