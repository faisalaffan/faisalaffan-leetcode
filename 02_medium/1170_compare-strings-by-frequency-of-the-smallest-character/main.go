package main

import (
	"fmt"
	"sort"
)

// LeetCode #1170: Compare Strings by Frequency of the Smallest Character
// https://leetcode.com/problems/compare-strings-by-frequency-of-the-smallest-character/
// Difficulty: Medium

// f(s) = frequency of smallest character in s.
// For each query word, count words in words[] with f(w) > f(query).

// Time: O((n + m) * L) where L = average string length
// Space: O(n)

func f(s string) int {
	minChar := s[0]
	count := 0
	for i := 0; i < len(s); i++ {
		if s[i] < minChar {
			minChar = s[i]
			count = 1
		} else if s[i] == minChar {
			count++
		}
	}
	return count
}

func numSmallerByFrequency(queries []string, words []string) []int {
	wordFreqs := make([]int, len(words))
	for i, w := range words {
		wordFreqs[i] = f(w)
	}
	sort.Ints(wordFreqs)

	result := make([]int, len(queries))
	for i, q := range queries {
		qf := f(q)
		// Binary search for first wordFreq > qf
		idx := sort.Search(len(wordFreqs), func(j int) bool {
			return wordFreqs[j] > qf
		})
		result[i] = len(wordFreqs) - idx
	}
	return result
}

func main() {
	fmt.Printf("%v (expected: [1])\n", numSmallerByFrequency([]string{"cbd"}, []string{"zaaaz"}))
	fmt.Printf("%v (expected: [1 2])\n", numSmallerByFrequency([]string{"bbb", "cc"}, []string{"a", "aa", "aaa", "aaaa"}))
}
