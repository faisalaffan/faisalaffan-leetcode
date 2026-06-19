package main

// LeetCode #472: Concatenated Words
// https://leetcode.com/problems/concatenated-words/
// Difficulty: Hard
// Approach: DP word break. Sort words by length, for each word check if it
// can be formed by concatenating shorter words (already processed).

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("472 - Concatenated Words")

	// Test cases
	words1 := []string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"}
	fmt.Printf("findAllConcatenatedWordsInADict(%v) = %v (expected: [catsdogcats dogcatsdog ratcatdogcat])\n",
		words1, findAllConcatenatedWordsInADict(words1))

	words2 := []string{"cat", "dog", "catdog"}
	fmt.Printf("findAllConcatenatedWordsInADict(%v) = %v (expected: [catdog])\n",
		words2, findAllConcatenatedWordsInADict(words2))

	words3 := []string{}
	fmt.Printf("findAllConcatenatedWordsInADict(%v) = %v (expected: [])\n",
		words3, findAllConcatenatedWordsInADict(words3))

	words4 := []string{"a"}
	fmt.Printf("findAllConcatenatedWordsInADict(%v) = %v (expected: [])\n",
		words4, findAllConcatenatedWordsInADict(words4))

	words5 := []string{"a", "aa", "aaa", "aaaa"}
	fmt.Printf("findAllConcatenatedWordsInADict(%v) = %v (expected: [aaa aaaa] or similar, depends on order)\n",
		words5, findAllConcatenatedWordsInADict(words5))

	words6 := []string{"", "a", "ab", "abc", "ababc"}
	fmt.Printf("findAllConcatenatedWordsInADict(%v) = %v\n",
		words6, findAllConcatenatedWordsInADict(words6))

	// Empty string should be handled - it's not a concatenated word
	// unless formed by 2+ empty strings, which conceptually doesn't count
}

func findAllConcatenatedWordsInADict(words []string) []string {
	// Sort words by length
	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	wordSet := make(map[string]bool)
	result := []string{}

	for _, word := range words {
		if word == "" {
			continue
		}
		if canForm(word, wordSet) {
			result = append(result, word)
		}
		wordSet[word] = true
	}

	return result
}

// canForm checks if word can be formed by concatenating 2+ words from wordSet
func canForm(word string, wordSet map[string]bool) bool {
	if len(word) == 0 {
		return false
	}

	// dp[i] = can form word[0:i]
	dp := make([]bool, len(word)+1)
	dp[0] = true

	for i := 1; i <= len(word); i++ {
		for j := 0; j < i; j++ {
			if !dp[j] {
				continue
			}
			// Skip the case where we use the entire word itself
			if j == 0 && i == len(word) {
				// We need at least 2 words, so this single word match doesn't count
				continue
			}
			if wordSet[word[j:i]] {
				dp[i] = true
				break
			}
		}
	}

	return dp[len(word)]
}
