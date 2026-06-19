package main

// LeetCode #1641: Count Sorted Vowel Strings
// https://leetcode.com/problems/count-sorted-vowel-strings/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountVowelStrings(1))
	fmt.Println(CountVowelStrings(2))
	fmt.Println(CountVowelStrings(33))
}

func CountVowelStrings(n int) int {
	// Time: O(N), Space: O(1)
	// Combinatorics: C(n+4, 4) = (n+4)*(n+3)*(n+2)*(n+1)/24
	return (n + 4) * (n + 3) * (n + 2) * (n + 1) / 24
}
