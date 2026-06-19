package main

// LeetCode #1950: Maximum of Minimum Values in All Subarrays
// https://leetcode.com/problems/maximum-of-minimum-values-in-all-subarrays/
// Difficulty: Medium [Paid]

import "fmt"

func main() {
	fmt.Println(MaxOfMinValues([]int{10, 20, 30, 50, 10, 70, 30}))
	fmt.Println(MaxOfMinValues([]int{1, 2, 3, 4, 5}))
}

// Time: O(n), Space: O(n)
func MaxOfMinValues(nums []int) []int {
	n := len(nums)
	result := make([]int, n+1) // result[k] for k-length subarrays
	for i := range result {
		result[i] = 0
	}

	// Find previous smaller and next smaller elements
	left := make([]int, n)
	right := make([]int, n)
	for i := 0; i < n; i++ {
		left[i] = -1
		right[i] = n
	}

	stack := make([]int, 0)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			left[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	stack = make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) > 0 && nums[stack[len(stack)-1]] >= nums[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) > 0 {
			right[i] = stack[len(stack)-1]
		}
		stack = append(stack, i)
	}

	for i := 0; i < n; i++ {
		length := right[i] - left[i] - 1
		if nums[i] > result[length] {
			result[length] = nums[i]
		}
	}

	// Fill missing values: max of min values for larger windows
	for i := n - 1; i >= 1; i-- {
		if result[i] < result[i+1] {
			result[i] = result[i+1]
		}
	}
	return result[1:]
}
