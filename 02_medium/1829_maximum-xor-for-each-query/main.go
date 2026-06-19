package main

// LeetCode #1829: Maximum XOR for Each Query
// https://leetcode.com/problems/maximum-xor-for-each-query/
// Difficulty: Medium
// Time: O(n), Space: O(1) excluding output

import "fmt"

func getMaximumXor(nums []int, maximumBit int) []int {
	n := len(nums)
	xor := 0
	for _, v := range nums {
		xor ^= v
	}

	maxVal := (1 << maximumBit) - 1
	result := make([]int, n)

	for i := n - 1; i >= 0; i-- {
		// Best k is the one that maximizes xor ^ k, i.e., xor ^ maxVal
		result[n-1-i] = xor ^ maxVal
		// Remove last element for next query
		xor ^= nums[i]
	}
	return result
}

func main() {
	fmt.Println(getMaximumXor([]int{0, 1, 1, 3}, 2)) // Expected: [0, 3, 2, 3]
	fmt.Println(getMaximumXor([]int{2, 3, 4, 7}, 3)) // Expected: [5, 2, 6, 5]
	fmt.Println(getMaximumXor([]int{0, 1, 2, 2, 5, 7}, 3)) // Expected: [4, 3, 6, 4, 6, 7]
}
