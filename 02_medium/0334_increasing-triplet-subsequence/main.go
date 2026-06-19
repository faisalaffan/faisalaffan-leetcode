package main

// LeetCode #334: Increasing Triplet Subsequence
// https://leetcode.com/problems/increasing-triplet-subsequence/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func increasingTriplet(nums []int) bool {
	first, second := 1<<31-1, 1<<31-1
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}
	return false
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", increasingTriplet([]int{1, 2, 3, 4, 5}))
	// Expected: true

	// Test case 2
	fmt.Println("Test 2:", increasingTriplet([]int{5, 4, 3, 2, 1}))
	// Expected: false

	// Test case 3
	fmt.Println("Test 3:", increasingTriplet([]int{2, 1, 5, 0, 4, 6}))
	// Expected: true
}
