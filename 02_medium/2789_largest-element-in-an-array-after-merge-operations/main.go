package main

// LeetCode #2789: Largest Element in an Array after Merge Operations
// https://leetcode.com/problems/largest-element-in-an-array-after-merge-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func LargestElementInAnArrayAfterMergeOperations(nums []int) int64 {
	n := len(nums)
	if n == 0 {
		return 0
	}

	result := int64(nums[n-1])
	for i := n - 2; i >= 0; i-- {
		if int64(nums[i]) <= result {
			result += int64(nums[i])
		} else {
			result = int64(nums[i])
		}
	}

	return result
}

func main() {
	fmt.Println(LargestElementInAnArrayAfterMergeOperations([]int{2, 3, 7, 9, 3}))
	fmt.Println(LargestElementInAnArrayAfterMergeOperations([]int{5, 3, 3}))
}
