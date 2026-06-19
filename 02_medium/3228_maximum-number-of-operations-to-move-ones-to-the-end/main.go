package main

// LeetCode #3228: Maximum Number of Operations to Move Ones to the End
// https://leetcode.com/problems/maximum-number-of-operations-to-move-ones-to-the-end/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maxOperations(s string) int {
	n := len(s)
	ans := 0
	ones := 0

	for i := 0; i < n; i++ {
		if s[i] == '1' {
			ones++
		} else if i > 0 && s[i-1] == '1' {
			ans += ones
		}
	}
	return ans
}

func main() {
	fmt.Println(maxOperations("1001101")) // Expected: 4
	fmt.Println(maxOperations("00111"))    // Expected: 0
}
