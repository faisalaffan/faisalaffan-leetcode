package main

// LeetCode #2828: Check if a String Is an Acronym of Words
// https://leetcode.com/problems/check-if-a-string-is-an-acronym-of-words/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(CheckIfAStringIsAnAcronymOfWords([]string{"alice", "bob", "charlie"}, "abc"))
	fmt.Println(CheckIfAStringIsAnAcronymOfWords([]string{"an", "apple"}, "a"))
}

func CheckIfAStringIsAnAcronymOfWords(words []string, s string) bool {
	if len(words) != len(s) {
		return false
	}
	for i, w := range words {
		if w[0] != s[i] {
			return false
		}
	}
	return true
}
