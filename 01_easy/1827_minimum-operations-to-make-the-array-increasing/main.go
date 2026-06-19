package main

// LeetCode #1827: Minimum Operations to Make the Array Increasing
// https://leetcode.com/problems/minimum-operations-to-make-the-array-increasing/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func MinOperations(nums []int) int {
	ops := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			needed := nums[i-1] - nums[i] + 1
			nums[i] += needed
			ops += needed
		}
	}
	return ops
}

func main() {
	fmt.Println(MinOperations([]int{1, 1, 1}))
	fmt.Println(MinOperations([]int{1, 5, 2, 4, 1}))
	fmt.Println(MinOperations([]int{8}))
}
