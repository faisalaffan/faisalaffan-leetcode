package main

// LeetCode #1941: Check if All Characters Have Equal Number of Occurrences
// https://leetcode.com/problems/check-if-all-characters-have-equal-number-of-occurrences/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckIfAllCharactersHaveEqualNumberOfOccurrences("abacbc")) // true
	fmt.Println(CheckIfAllCharactersHaveEqualNumberOfOccurrences("aaabb"))  // false
}

// Time: O(n), Space: O(1) (max 26 chars)
func CheckIfAllCharactersHaveEqualNumberOfOccurrences(s string) bool {
	freq := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		freq[s[i]]++
	}

	var target int
	for _, v := range freq {
		target = v
		break
	}
	for _, v := range freq {
		if v != target {
			return false
		}
	}
	return true
}
