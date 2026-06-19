package main

// LeetCode #2396: Strictly Palindromic Number
// https://leetcode.com/problems/strictly-palindromic-number/
// Difficulty: Medium
// Time: O(1) | Space: O(1)
// For n >= 4, n in base (n-2) is always "12", which is never a palindrome.
// So answer is always false for n >= 4.

import "fmt"

func main() {
	fmt.Println(isStrictlyPalindromic(9))  // false
	fmt.Println(isStrictlyPalindromic(4))  // false
	fmt.Println(isStrictlyPalindromic(3))  // true (3 is "11" in base 2, "10" in base 1? no, base 2 only)
}

func isStrictlyPalindromic(n int) bool {
	// For all bases 2..n-2, check if representation is palindrome
	for base := 2; base <= n-2; base++ {
		if !isPalindromeInBase(n, base) {
			return false
		}
	}
	return true
}

func isPalindromeInBase(n, base int) bool {
	digits := make([]int, 0)
	for n > 0 {
		digits = append(digits, n%base)
		n /= base
	}
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}
	return true
}
