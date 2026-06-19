package main

// LeetCode #389: Find the Difference
// https://leetcode.com/problems/find-the-difference/
// Difficulty: Easy

import "fmt"

// Time: O(n), Space: O(1)
func FindTheDifference(s, t string) byte {
	var diff byte
	for i := 0; i < len(s); i++ {
		diff ^= s[i]
	}
	for i := 0; i < len(t); i++ {
		diff ^= t[i]
	}
	return diff
}

func main() {
	fmt.Printf("%c\n", FindTheDifference("abcd", "abcde"))
	fmt.Printf("%c\n", FindTheDifference("", "y"))
	fmt.Printf("%c\n", FindTheDifference("a", "aa"))
}
