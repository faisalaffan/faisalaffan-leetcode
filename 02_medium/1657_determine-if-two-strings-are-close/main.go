package main

// LeetCode #1657: Determine if Two Strings Are Close
// https://leetcode.com/problems/determine-if-two-strings-are-close/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(CloseStrings("abc", "bca"))
	fmt.Println(CloseStrings("a", "aa"))
	fmt.Println(CloseStrings("cabbba", "abbccc"))
}

func CloseStrings(word1 string, word2 string) bool {
	// Time: O(N + M + 26 log 26), Space: O(1)
	if len(word1) != len(word2) {
		return false
	}

	freq1 := make([]int, 26)
	freq2 := make([]int, 26)
	set1 := make([]bool, 26)
	set2 := make([]bool, 26)

	for _, ch := range word1 {
		freq1[ch-'a']++
		set1[ch-'a'] = true
	}
	for _, ch := range word2 {
		freq2[ch-'a']++
		set2[ch-'a'] = true
	}

	// Check same character set
	for i := 0; i < 26; i++ {
		if set1[i] != set2[i] {
			return false
		}
	}

	// Check same frequency multiset
	sort.Ints(freq1)
	sort.Ints(freq2)
	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}
