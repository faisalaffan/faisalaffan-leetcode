package main

// LeetCode #2710: Remove Trailing Zeros From a String
// https://leetcode.com/problems/remove-trailing-zeros-from-a-string/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(RemoveTrailingZerosFromAString("51230100"))
	fmt.Println(RemoveTrailingZerosFromAString("123"))
}

func RemoveTrailingZerosFromAString(num string) string {
	return strings.TrimRight(num, "0")
}
