package main

// LeetCode #1616: Split Two Strings to Make Palindrome
// https://leetcode.com/problems/split-two-strings-to-make-palindrome/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CheckPalindromeFormation("x", "y"))
	fmt.Println(CheckPalindromeFormation("abdef", "fecab"))
	fmt.Println(CheckPalindromeFormation("ulacfd", "jizalu"))
}

func CheckPalindromeFormation(a string, b string) bool {
	// Time: O(N), Space: O(1)
	return canForm(a, b) || canForm(b, a)
}

func canForm(a, b string) bool {
	left := 0
	right := len(a) - 1

	// Find the first mismatch from the outside
	for left < right && a[left] == b[right] {
		left++
		right--
	}

	if left >= right {
		return true
	}

	// Check if a[left..right] is palindrome
	return isPalindrome(a, left, right) || isPalindrome(b, left, right)
}

func isPalindrome(s string, left, right int) bool {
	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
}
