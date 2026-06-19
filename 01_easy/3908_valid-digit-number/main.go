package main

// LeetCode #3908: Valid Digit Number
// https://leetcode.com/problems/valid-digit-number/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ValidDigitNumber(101, 0))
	fmt.Println(ValidDigitNumber(232, 2))
	fmt.Println(ValidDigitNumber(5, 1))
}

// Time: O(log n)
// Space: O(1)
func ValidDigitNumber(n int, x int) bool {
	// Check if n starts with digit x
	first := n
	for first >= 10 {
		first /= 10
	}
	if first == x {
		return false
	}

	// Check if n contains digit x
	m := n
	for m > 0 {
		if m%10 == x {
			return true
		}
		m /= 10
	}
	return false
}
