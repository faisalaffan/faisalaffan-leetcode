package main

// LeetCode #3931: Check Adjacent Digit Differences
// https://leetcode.com/problems/check-adjacent-digit-differences/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckAdjacentDigitDifferences("132"))
	fmt.Println(CheckAdjacentDigitDifferences("129"))
}

// Time: O(n)
// Space: O(1)
func CheckAdjacentDigitDifferences(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		diff := int(s[i]) - int(s[i+1])
		if diff < 0 {
			diff = -diff
		}
		if diff > 2 {
			return false
		}
	}
	return true
}
