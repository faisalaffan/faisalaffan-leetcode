package main

// LeetCode #2068: Check Whether Two Strings are Almost Equivalent
// https://leetcode.com/problems/check-whether-two-strings-are-almost-equivalent/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(CheckWhetherTwoStringsAreAlmostEquivalent("aaaa", "bccb"))   // false
	fmt.Println(CheckWhetherTwoStringsAreAlmostEquivalent("abcdeef", "abaaacc")) // true
	fmt.Println(CheckWhetherTwoStringsAreAlmostEquivalent("cccddabba", "babababab")) // true
}

// Time: O(n), Space: O(1)
func CheckWhetherTwoStringsAreAlmostEquivalent(word1 string, word2 string) bool {
	freq := make([]int, 26)
	for i := 0; i < len(word1); i++ {
		freq[word1[i]-'a']++
	}
	for i := 0; i < len(word2); i++ {
		freq[word2[i]-'a']--
	}
	for _, v := range freq {
		if v < 0 {
			v = -v
		}
		if v > 3 {
			return false
		}
	}
	return true
}
