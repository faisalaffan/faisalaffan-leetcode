package main

// LeetCode #1754: Largest Merge Of Two Strings
// https://leetcode.com/problems/largest-merge-of-two-strings/
// Difficulty: Medium
// Time: O(n^2), Space: O(n)

import "fmt"

func largestMerge(word1 string, word2 string) string {
	result := make([]byte, 0, len(word1)+len(word2))
	i, j := 0, 0

	for i < len(word1) && j < len(word2) {
		// Pick the character that leads to a larger overall string
		if word1[i:] > word2[j:] {
			result = append(result, word1[i])
			i++
		} else {
			result = append(result, word2[j])
			j++
		}
	}

	// Append remaining characters
	result = append(result, word1[i:]...)
	result = append(result, word2[j:]...)
	return string(result)
}

func main() {
	fmt.Println(largestMerge("cabaa", "bcaaa")) // Expected: "cbcabaaaaa"
	fmt.Println(largestMerge("abcabc", "abdcaba")) // Expected: "abdcabcabcaba"
	fmt.Println(largestMerge("a", "b")) // Expected: "ba"
}
