package main

// LeetCode #2299: Strong Password Checker II
// https://leetcode.com/problems/strong-password-checker-ii/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(StrongPasswordCheckerIi("IloveLe3tcode!")) // true
	fmt.Println(StrongPasswordCheckerIi("Me+You--IsMyDream")) // false
	fmt.Println(StrongPasswordCheckerIi("1aB!")) // false
}

func StrongPasswordCheckerIi(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasLower, hasUpper, hasDigit, hasSpecial := false, false, false, false
	special := "!@#$%^&*()-+"

	for i := 0; i < len(password); i++ {
		if i > 0 && password[i] == password[i-1] {
			return false
		}
		ch := password[i]
		switch {
		case ch >= 'a' && ch <= 'z':
			hasLower = true
		case ch >= 'A' && ch <= 'Z':
			hasUpper = true
		case ch >= '0' && ch <= '9':
			hasDigit = true
		default:
			for j := 0; j < len(special); j++ {
				if ch == special[j] {
					hasSpecial = true
					break
				}
			}
		}
	}
	return hasLower && hasUpper && hasDigit && hasSpecial
}
