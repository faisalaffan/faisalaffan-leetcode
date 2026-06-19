package main

// LeetCode #2273: Find Resultant Array After Removing Anagrams
// https://leetcode.com/problems/find-resultant-array-after-removing-anagrams/
// Difficulty: Easy
// Time O(n * m log m) | Space O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindResultantArrayAfterRemovingAnagrams([]string{"abba", "baba", "bbaa", "cd", "cd"})) // ["abba","cd"]
	fmt.Println(FindResultantArrayAfterRemovingAnagrams([]string{"a", "b", "c", "d", "e"}))            // ["a","b","c","d","e"]
}

func FindResultantArrayAfterRemovingAnagrams(words []string) []string {
	result := []string{words[0]}
	prev := sortWord(words[0])

	for i := 1; i < len(words); i++ {
		curr := sortWord(words[i])
		if curr != prev {
			result = append(result, words[i])
			prev = curr
		}
	}
	return result
}

func sortWord(w string) string {
	b := []byte(w)
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })
	return string(b)
}
