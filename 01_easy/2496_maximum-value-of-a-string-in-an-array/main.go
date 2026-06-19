package main

// LeetCode #2496: Maximum Value of a String in an Array
// https://leetcode.com/problems/maximum-value-of-a-string-in-an-array/
// Difficulty: Easy
// Time O(n * m) | Space O(1)

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(MaximumValueOfAStringInAnArray([]string{"alic3", "bob", "3", "4", "00000"})) // 5
	fmt.Println(MaximumValueOfAStringInAnArray([]string{"1", "01", "001", "0001"}))           // 1
}

func MaximumValueOfAStringInAnArray(strs []string) int {
	maxVal := 0
	for _, s := range strs {
		val := 0
		if isNumeric(s) {
			val, _ = strconv.Atoi(s)
		} else {
			val = len(s)
		}
		if val > maxVal {
			maxVal = val
		}
	}
	return maxVal
}

func isNumeric(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return len(s) > 0
}
