package main

// LeetCode #456: 132 Pattern
// https://leetcode.com/problems/132-pattern/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func find132pattern(nums []int) bool {
	stack := []int{}
	s3 := -1 << 31 // nums[k] (the "2" in 132)

	// Iterate from right to left
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < s3 {
			return true
		}
		// nums[i] > stack top means nums[i] could be the "3"
		for len(stack) > 0 && nums[i] > stack[len(stack)-1] {
			s3 = stack[len(stack)-1] // s3 is the current largest "2" candidate
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, nums[i])
	}
	return false
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", find132pattern([]int{1, 2, 3, 4}))
	// Expected: false

	// Test case 2
	fmt.Println("Test 2:", find132pattern([]int{3, 1, 4, 2}))
	// Expected: true

	// Test case 3
	fmt.Println("Test 3:", find132pattern([]int{-1, 3, 2, 0}))
	// Expected: true
}
