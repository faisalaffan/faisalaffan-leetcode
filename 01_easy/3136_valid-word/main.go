package main

// LeetCode #3136: Valid Word
// https://leetcode.com/problems/valid-word/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: isValid
	fmt.Println(ValidWord("Hello123")) // true
	fmt.Println(ValidWord("hi"))       // false
	fmt.Println(ValidWord("AAab"))     // false (no vowel)
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: isValid
func ValidWord(word string) bool {
	if len(word) < 3 {
		return false
	}
	hasVowel := false
	hasConsonant := false
	for i := 0; i < len(word); i++ {
		c := word[i]
		if c >= '0' && c <= '9' {
			continue
		}
		if c >= 'A' && c <= 'Z' {
			c = c - 'A' + 'a'
		}
		if c < 'a' || c > 'z' {
			return false
		}
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
			hasVowel = true
		} else {
			hasConsonant = true
		}
	}
	return hasVowel && hasConsonant
}
