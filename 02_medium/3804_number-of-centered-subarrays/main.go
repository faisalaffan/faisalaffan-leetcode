package main

// LeetCode #3804: Number of Centered Subarrays
// https://leetcode.com/problems/number-of-centered-subarrays/
// Difficulty: Medium
// Time: O(n^2) | Space: O(n)
// Approach: For each start index, expand subarrays and track running sum
// with a set of seen elements. A subarray is centered if its sum equals
// at least one element within it.

import "fmt"

func NumberOfCenteredSubarrays(nums []int) int {
	n := len(nums)
	ans := 0

	for i := 0; i < n; i++ {
		sum := 0
		seen := make(map[int]bool)
		for j := i; j < n; j++ {
			seen[nums[j]] = true
			sum += nums[j]
			if seen[sum] {
				ans++
			}
		}
	}

	return ans
}

func main() {
	// Example 1
	fmt.Println(NumberOfCenteredSubarrays([]int{-1, 1, 0})) // Expected: 5

	// Example 2
	fmt.Println(NumberOfCenteredSubarrays([]int{2, -3})) // Expected: 2

	// Example 3
	fmt.Println(NumberOfCenteredSubarrays([]int{1, 2, 3})) // Expected: 3 (all single elements are centered)
}
