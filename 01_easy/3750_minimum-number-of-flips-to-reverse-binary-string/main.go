package main

// LeetCode #3750: Minimum Number of Flips to Reverse Binary String
// https://leetcode.com/problems/minimum-number-of-flips-to-reverse-binary-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumNumberOfFlipsToReverseBinaryString(7))
	fmt.Println(MinimumNumberOfFlipsToReverseBinaryString(10))
}

// Time: O(log n)
// Space: O(1)
func MinimumNumberOfFlipsToReverseBinaryString(n int) int {
	// XOR n with its reverse, count bits
	rev := 0
	for x := n; x > 0; x >>= 1 {
		rev = (rev << 1) | (x & 1)
	}
	xor := n ^ rev
	ans := 0
	for xor > 0 {
		ans += xor & 1
		xor >>= 1
	}
	return ans
}
