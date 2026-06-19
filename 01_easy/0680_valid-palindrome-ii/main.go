package main

// LeetCode #680: Valid Palindrome II
// https://leetcode.com/problems/valid-palindrome-ii/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(validPalindrome("aba"))    // true
	fmt.Println(validPalindrome("abca"))   // true
	fmt.Println(validPalindrome("abc"))    // false
	fmt.Println(validPalindrome("deeee"))  // true
}

// validPalindrome checks if the string can be a palindrome after deleting at most one character.
// Time: O(n). Space: O(1).
func validPalindrome(s string) bool {
	l, r := 0, len(s)-1
	for l < r {
		if s[l] != s[r] {
			return isPal(s, l+1, r) || isPal(s, l, r-1)
		}
		l++
		r--
	}
	return true
}

func isPal(s string, l, r int) bool {
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
