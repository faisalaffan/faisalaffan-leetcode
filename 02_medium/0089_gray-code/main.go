package main

// LeetCode #89: Gray Code
// https://leetcode.com/problems/gray-code/
// Difficulty: Medium

import "fmt"

func grayCode(n int) []int {
	result := make([]int, 1<<n)
	for i := 0; i < len(result); i++ {
		result[i] = i ^ (i >> 1)
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(grayCode(2)) // [0 1 3 2]

	// Test case 2
	fmt.Println(grayCode(1)) // [0 1]

	// Test case 3
	fmt.Println(grayCode(3)) // [0 1 3 2 6 7 5 4]
}

// Time: O(2^n) | Space: O(1)
