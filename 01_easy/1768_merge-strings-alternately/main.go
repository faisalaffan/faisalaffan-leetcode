package main

// LeetCode #1768: Merge Strings Alternately
// https://leetcode.com/problems/merge-strings-alternately/
// Difficulty: Easy

import "fmt"

// Time: O(n+m), Space: O(n+m)
func MergeAlternately(word1 string, word2 string) string {
	result := make([]byte, 0, len(word1)+len(word2))
	i, j := 0, 0
	for i < len(word1) && j < len(word2) {
		result = append(result, word1[i], word2[j])
		i++
		j++
	}
	result = append(result, word1[i:]...)
	result = append(result, word2[j:]...)
	return string(result)
}

func main() {
	fmt.Println(MergeAlternately("abc", "pqr"))
	fmt.Println(MergeAlternately("ab", "pqrs"))
	fmt.Println(MergeAlternately("abcd", "pq"))
}
