package main

// LeetCode #3120: Count the Number of Special Characters I
// https://leetcode.com/problems/count-the-number-of-special-characters-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: numberOfSpecialChars
	fmt.Println(CountTheNumberOfSpecialCharactersI("aaAbcBC")) // 3
	fmt.Println(CountTheNumberOfSpecialCharactersI("abcd"))    // 0
	fmt.Println(CountTheNumberOfSpecialCharactersI("abAB"))   // 2
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: numberOfSpecialChars
func CountTheNumberOfSpecialCharactersI(word string) int {
	lower := make(map[byte]bool)
	upper := make(map[byte]bool)
	for i := 0; i < len(word); i++ {
		c := word[i]
		if c >= 'a' && c <= 'z' {
			lower[c] = true
		} else if c >= 'A' && c <= 'Z' {
			upper[c] = true
		}
	}
	count := 0
	for c := byte('a'); c <= 'z'; c++ {
		if lower[c] && upper[c-'a'+'A'] {
			count++
		}
	}
	return count
}
