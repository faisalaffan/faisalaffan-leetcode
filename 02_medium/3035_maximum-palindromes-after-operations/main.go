package main

// LeetCode #3035: Maximum Palindromes After Operations
// https://leetcode.com/problems/maximum-palindromes-after-operations/
// Difficulty: Medium
// Time: O(n * L + A log A) | Space: O(A)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxPalindromesAfterOperations([]string{"abbb", "ba", "aa"}))
	fmt.Println(maxPalindromesAfterOperations([]string{"abc", "ab"}))
	fmt.Println(maxPalindromesAfterOperations([]string{"cd", "ef", "a"}))
}

func maxPalindromesAfterOperations(words []string) int {
	freq := [26]int{}
	lens := make([]int, len(words))
	for i, w := range words {
		lens[i] = len(w)
		for _, ch := range w {
			freq[ch-'a']++
		}
	}
	pairs := 0
	for _, c := range freq {
		pairs += c / 2
	}
	sort.Slice(lens, func(i, j int) bool {
		return lens[i] < lens[j]
	})
	ans := 0
	for _, l := range lens {
		need := l / 2
		if pairs >= need {
			pairs -= need
			ans++
		}
	}
	return ans
}
