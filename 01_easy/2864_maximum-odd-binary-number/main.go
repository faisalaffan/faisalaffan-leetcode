package main

// LeetCode #2864: Maximum Odd Binary Number
// https://leetcode.com/problems/maximum-odd-binary-number/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(MaximumOddBinaryNumber("010"))
	fmt.Println(MaximumOddBinaryNumber("0101"))
}

func MaximumOddBinaryNumber(s string) string {
	ones := strings.Count(s, "1")
	zeros := len(s) - ones
	return strings.Repeat("1", ones-1) + strings.Repeat("0", zeros) + "1"
}
