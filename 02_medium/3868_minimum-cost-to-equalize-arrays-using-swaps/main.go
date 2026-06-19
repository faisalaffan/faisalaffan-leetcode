package main

// LeetCode #3868: Minimum Cost to Equalize Arrays Using Swaps
// https://leetcode.com/problems/minimum-cost-to-equalize-arrays-using-swaps/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Count frequencies in both arrays. If any value's total count is odd,
// impossible (return -1). Min cost = sum of positive differences / 2.

import "fmt"

func MinimumCostToEqualizeArraysUsingSwaps(nums1 []int, nums2 []int) int {
	cnt := make(map[int]int)
	for _, v := range nums1 {
		cnt[v]++
	}
	for _, v := range nums2 {
		cnt[v]--
	}

	posDiff := 0
	for _, c := range cnt {
		if c%2 != 0 {
			return -1
		}
		if c > 0 {
			posDiff += c
		}
	}
	return posDiff / 2
}

func main() {
	// Example 1
	fmt.Println(MinimumCostToEqualizeArraysUsingSwaps([]int{10, 20}, []int{20, 10})) // Expected: 0

	// Example 2
	fmt.Println(MinimumCostToEqualizeArraysUsingSwaps([]int{10, 10}, []int{20, 20})) // Expected: 1

	// Example 3
	fmt.Println(MinimumCostToEqualizeArraysUsingSwaps([]int{10, 20}, []int{30, 40})) // Expected: -1
}
