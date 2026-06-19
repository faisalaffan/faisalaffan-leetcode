package main

// LeetCode #3941: Password Strength
// https://leetcode.com/problems/password-strength/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Track distinct characters per category. Sum points:
// lowercase=1, uppercase=2, digit=3, special=5.

import "fmt"

func PasswordStrength(password string) int {
	lower := make(map[byte]bool)
	upper := make(map[byte]bool)
	digit := make(map[byte]bool)
	special := make(map[byte]bool)

	for i := 0; i < len(password); i++ {
		ch := password[i]
		if ch >= 'a' && ch <= 'z' {
			lower[ch] = true
		} else if ch >= 'A' && ch <= 'Z' {
			upper[ch] = true
		} else if ch >= '0' && ch <= '9' {
			digit[ch] = true
		} else {
			// Special: ! @ # $
			special[ch] = true
		}
	}

	return len(lower)*1 + len(upper)*2 + len(digit)*3 + len(special)*5
}

func main() {
	// Example 1
	fmt.Println(PasswordStrength("aA1!")) // Expected: 11

	// Example 2
	fmt.Println(PasswordStrength("bbB11#")) // Expected: 11
}
