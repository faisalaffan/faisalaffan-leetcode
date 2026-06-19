package main

// LeetCode #1680: Concatenation of Consecutive Binary Numbers
// https://leetcode.com/problems/concatenation-of-consecutive-binary-numbers/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

const mod = 1_000_000_007

func concatenatedBinary(n int) int {
	result := 0
	lenBits := 0

	for i := 1; i <= n; i++ {
		if i&(i-1) == 0 {
			lenBits++
		}
		result = (result<<lenBits | i) % mod
	}
	return result
}

func main() {
	fmt.Println(concatenatedBinary(1))   // Expected: 1
	fmt.Println(concatenatedBinary(3))   // Expected: 27
	fmt.Println(concatenatedBinary(12))  // Expected: 505379714
}
