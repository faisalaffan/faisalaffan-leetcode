package main

// LeetCode #2567: Minimum Score by Changing Two Elements
// https://leetcode.com/problems/minimum-score-by-changing-two-elements/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minimizeSum(nums []int) int {
	n := len(nums)
	if n <= 3 {
		return 0
	}
	sort.Ints(nums)

	// After changing two elements, the min sum possible is:
	// Option 1: change two smallest to = third smallest, score = nums[n-1] - nums[2]
	// Option 2: change two largest to = third largest, score = nums[n-3] - nums[0]
	// Option 3: change one smallest and one largest, score = nums[n-2] - nums[1]

	diff1 := nums[n-1] - nums[2]
	diff2 := nums[n-3] - nums[0]
	diff3 := nums[n-2] - nums[1]

	result := diff1
	if diff2 < result {
		result = diff2
	}
	if diff3 < result {
		result = diff3
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimizeSum([]int{1, 4, 3}))
	// Expected: 0

	// Test case 2
	fmt.Println("Test 2:", minimizeSum([]int{1, 4, 7, 8, 5}))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", minimizeSum([]int{1, 2, 3, 4, 5}))
	// Expected: 2
}
