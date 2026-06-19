package main

// LeetCode #2351: First Letter to Appear Twice
// https://leetcode.com/problems/first-letter-to-appear-twice/
// Difficulty: Easy
// Time O(n) | Space O(1)

import "fmt"

func main() {
	fmt.Println(string(FirstLetterToAppearTwice("abccbaacz"))) // "c"
	fmt.Println(string(FirstLetterToAppearTwice("abcdd")))      // "d"
}

func FirstLetterToAppearTwice(s string) byte {
	seen := [26]bool{}
	for i := 0; i < len(s); i++ {
		idx := s[i] - 'a'
		if seen[idx] {
			return s[i]
		}
		seen[idx] = true
	}
	return 0
}
