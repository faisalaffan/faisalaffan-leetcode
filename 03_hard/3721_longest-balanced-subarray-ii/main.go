package main

// LeetCode #3721: Longest Balanced Subarray II
// https://leetcode.com/problems/longest-balanced-subarray-ii/
// Difficulty: Hard
//
// Find longest subarray where number of distinct even numbers equals
// number of distinct odd numbers.
//
// Approach: Sliding window with frequency maps. For each window,
// track distinct even and odd counts using hash maps.

import "fmt"

func main() {
	// Example 1
	fmt.Println(longestBalanced([]int{2, 5, 4, 3}))
	// Example 2
	fmt.Println(longestBalanced([]int{1, 3, 5, 2, 4}))
	// Edge: single element
	fmt.Println(longestBalanced([]int{1}))
	// Edge: all even
	fmt.Println(longestBalanced([]int{2, 4, 6}))
}

func longestBalanced(nums []int) int {
	n := len(nums)
	result := 0

	// For each starting position, expand window
	for i := 0; i < n; i++ {
		evenSet := make(map[int]bool)
		oddSet := make(map[int]bool)
		evenCount := 0
		oddCount := 0

		for j := i; j < n; j++ {
			if nums[j]%2 == 0 {
				if !evenSet[nums[j]] {
					evenSet[nums[j]] = true
					evenCount++
				}
			} else {
				if !oddSet[nums[j]] {
					oddSet[nums[j]] = true
					oddCount++
				}
			}
			if evenCount == oddCount {
				length := j - i + 1
				if length > result {
					result = length
				}
			}
		}
	}

	return result
}
