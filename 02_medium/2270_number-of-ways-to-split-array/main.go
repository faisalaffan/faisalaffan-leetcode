package main

// LeetCode #2270: Number of Ways to Split Array
// https://leetcode.com/problems/number-of-ways-to-split-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func waysToSplitArray(nums []int) int {
	n := len(nums)
	total := int64(0)
	for _, v := range nums {
		total += int64(v)
	}

	var prefix int64 = 0
	ways := 0
	for i := 0; i < n-1; i++ {
		prefix += int64(nums[i])
		if prefix >= total-prefix {
			ways++
		}
	}
	return ways
}

func main() {
	// Test case 1
	fmt.Println(waysToSplitArray([]int{10, 4, -8, 7}))
	// Expected: 2

	// Test case 2
	fmt.Println(waysToSplitArray([]int{2, 3, 1, 0}))
	// Expected: 2

	// Test case 3
	fmt.Println(waysToSplitArray([]int{-1, -2, -3, -4}))
	// Expected: 0
}
