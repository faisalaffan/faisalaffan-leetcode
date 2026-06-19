package main

// LeetCode #2592: Maximize Greatness of an Array
// https://leetcode.com/problems/maximize-greatness-of-an-array/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximizeGreatness(nums []int) int {
	sort.Ints(nums)
	j := 0
	for _, v := range nums {
		if v > nums[j] {
			j++
		}
	}
	return j
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximizeGreatness([]int{1, 3, 5, 2, 1, 3, 1}))
	// Expected: 4

	// Test case 2
	fmt.Println("Test 2:", maximizeGreatness([]int{1, 2, 3, 4}))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", maximizeGreatness([]int{1, 1, 1}))
	// Expected: 0
}
