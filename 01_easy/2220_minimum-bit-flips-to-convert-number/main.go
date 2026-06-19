package main

// LeetCode #2220: Minimum Bit Flips to Convert Number
// https://leetcode.com/problems/minimum-bit-flips-to-convert-number/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumBitFlipsToConvertNumber(10, 7))  // 3
	fmt.Println(MinimumBitFlipsToConvertNumber(3, 4))   // 3
}

// Time: O(1), Space: O(1)
func MinimumBitFlipsToConvertNumber(start int, goal int) int {
	xor := start ^ goal
	count := 0
	for xor > 0 {
		count += xor & 1
		xor >>= 1
	}
	return count
}
