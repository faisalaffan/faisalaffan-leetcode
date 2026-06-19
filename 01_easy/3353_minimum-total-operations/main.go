package main

// LeetCode #3353: Minimum Total Operations
// https://leetcode.com/problems/minimum-total-operations/
// Difficulty: Easy [Paid]

import "fmt"

func main() {
	fmt.Println(MinimumTotalOperations([]int{1, 2, 3, 4}))
	fmt.Println(MinimumTotalOperations([]int{1, 1, 1}))
}

// MinimumTotalOperations returns the minimum number of operations to make all elements zero.
// Each operation picks a subarray and subtracts the minimum value in it from all its elements.
// Time: O(n). Space: O(1).
func MinimumTotalOperations(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	ops := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			ops += nums[i] - nums[i-1]
		}
	}
	return ops
}
