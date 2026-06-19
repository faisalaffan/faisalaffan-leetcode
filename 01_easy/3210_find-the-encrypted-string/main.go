package main

// LeetCode #3210: Find the Encrypted String
// https://leetcode.com/problems/find-the-encrypted-string/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(FindTheEncryptedString("dart", 3))
	fmt.Println(FindTheEncryptedString("aaa", 1))
	fmt.Println(FindTheEncryptedString("abcd", 5))
}

// FindTheEncryptedString returns the encrypted string by rotating each character by k positions forward.
// Time: O(n). Space: O(n).
func FindTheEncryptedString(s string, k int) string {
	n := len(s)
	if n == 0 {
		return s
	}
	k %= n
	return s[k:] + s[:k]
}
