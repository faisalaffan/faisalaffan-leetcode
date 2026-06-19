package main

// LeetCode #3775: Reverse Words With Same Vowel Count
// https://leetcode.com/problems/reverse-words-with-same-vowel-count/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"strings"
)

func reverseWordsWithSameVowelCount(s string) string {
	words := strings.Fields(s)
	if len(words) == 0 {
		return s
	}

	vowels := "aeiou"
	countVowels := func(w string) int {
		c := 0
		for _, ch := range w {
			if strings.ContainsRune(vowels, ch) {
				c++
			}
		}
		return c
	}

	firstCnt := countVowels(words[0])
	for i := 1; i < len(words); i++ {
		if countVowels(words[i]) == firstCnt {
			// Reverse word
			runes := []rune(words[i])
			for l, r := 0, len(runes)-1; l < r; l, r = l+1, r-1 {
				runes[l], runes[r] = runes[r], runes[l]
			}
			words[i] = string(runes)
		}
	}

	return strings.Join(words, " ")
}

func main() {
	fmt.Println(reverseWordsWithSameVowelCount("cat and mice"))
	fmt.Println(reverseWordsWithSameVowelCount("book is nice"))
	fmt.Println(reverseWordsWithSameVowelCount("banana healthy"))
}
