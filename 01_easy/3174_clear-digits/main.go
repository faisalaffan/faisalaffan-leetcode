package main

// LeetCode #3174: Clear Digits
// https://leetcode.com/problems/clear-digits/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(ClearDigits("abc"))
	fmt.Println(ClearDigits("cb34"))
	fmt.Println(ClearDigits("a1b2c3"))
}

// ClearDigits removes all digits and their nearest non-digit character to the left.
// Time: O(n). Space: O(n).
func ClearDigits(s string) string {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
