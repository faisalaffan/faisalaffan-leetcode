package main

// LeetCode #1526: Minimum Number of Increments on Subarrays to Form a Target Array
// https://leetcode.com/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/
// Difficulty: Hard
//
// Approach: One Pass (Think of diff array)
// Consider the target array. Each time we need to "raise" the level from
// previous position, that's a new operation. The answer is:
//   target[0] + sum over i > 0 of max(0, target[i] - target[i-1])
// Intuition: Imagine we start from all zeros and apply operations. Each operation
// increments a contiguous segment. The difference array gives us the number of
// times we need to start a new segment at each position.

import "fmt"

func main() {
	// Example 1
	fmt.Println(minNumberOperations([]int{1, 2, 3, 2, 1}))
	// Expected: 3

	// Example 2
	fmt.Println(minNumberOperations([]int{3, 1, 1, 2}))
	// Expected: 4

	// Example 3
	fmt.Println(minNumberOperations([]int{3, 1, 5, 4, 2}))
	// Expected: 7
}

func minNumberOperations(target []int) int {
	if len(target) == 0 {
		return 0
	}
	ans := target[0]
	for i := 1; i < len(target); i++ {
		if target[i] > target[i-1] {
			ans += target[i] - target[i-1]
		}
	}
	return ans
}
