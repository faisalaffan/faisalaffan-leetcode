package main

// LeetCode #2750: Ways to Split Array Into Good Subarrays
// https://leetcode.com/problems/ways-to-split-array-into-good-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func WaysToSplitArrayIntoGoodSubarrays(nums []int) int {
	// Find positions of 1s
	ones := make([]int, 0)
	for i, v := range nums {
		if v == 1 {
			ones = append(ones, i)
		}
	}

	if len(ones) == 0 {
		return 0
	}

	const mod = 1_000_000_007
	result := 1
	for i := 1; i < len(ones); i++ {
		gap := ones[i] - ones[i-1]
		result = (result * gap) % mod
	}

	return result
}

func main() {
	fmt.Println(WaysToSplitArrayIntoGoodSubarrays([]int{0, 1, 0, 0, 1}))
	fmt.Println(WaysToSplitArrayIntoGoodSubarrays([]int{1, 1, 1}))
	fmt.Println(WaysToSplitArrayIntoGoodSubarrays([]int{0, 0, 0}))
}
