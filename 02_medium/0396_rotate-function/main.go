package main

// LeetCode #396: Rotate Function
// https://leetcode.com/problems/rotate-function/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maxRotateFunction(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	sum := 0
	f0 := 0
	for i, num := range nums {
		sum += num
		f0 += i * num
	}

	maxVal := f0
	prev := f0
	for k := 1; k < n; k++ {
		curr := prev + sum - n*nums[n-k]
		if curr > maxVal {
			maxVal = curr
		}
		prev = curr
	}
	return maxVal
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxRotateFunction([]int{4, 3, 2, 6}))
	// Expected: 26

	// Test case 2
	fmt.Println("Test 2:", maxRotateFunction([]int{100}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", maxRotateFunction([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
	// Expected: 330
}
