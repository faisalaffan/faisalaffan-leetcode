package main

// LeetCode #2697: Lexicographically Smallest Palindrome
// https://leetcode.com/problems/lexicographically-smallest-palindrome/
// Difficulty: Easy
// Time: O(n) | Space: O(n)

import "fmt"

func main() {
	fmt.Println(LexicographicallySmallestPalindrome("egcfe"))
	fmt.Println(LexicographicallySmallestPalindrome("abcd"))
}

func LexicographicallySmallestPalindrome(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1
	for i < j {
		if runes[i] < runes[j] {
			runes[j] = runes[i]
		} else {
			runes[i] = runes[j]
		}
		i++
		j--
	}
	return string(runes)
}
