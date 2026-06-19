package main

// LeetCode #1930: Unique Length-3 Palindromic Subsequences
// https://leetcode.com/problems/unique-length-3-palindromic-subsequences/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountPalindromicSubsequence("aabca"))
	fmt.Println(CountPalindromicSubsequence("adc"))
	fmt.Println(CountPalindromicSubsequence("bbcbaba"))
}

// Time: O(n * 26) = O(n), Space: O(1)
func CountPalindromicSubsequence(s string) int {
	// For each character, find first and last occurrence
	first := make([]int, 26)
	last := make([]int, 26)
	for i := 0; i < 26; i++ {
		first[i] = -1
		last[i] = -1
	}
	for i := 0; i < len(s); i++ {
		idx := int(s[i] - 'a')
		if first[idx] == -1 {
			first[idx] = i
		}
		last[idx] = i
	}

	count := 0
	for c := 0; c < 26; c++ {
		if first[c] != -1 && last[c]-first[c] > 1 {
			// Count unique characters between first and last occurrence
			seen := make([]bool, 26)
			for i := first[c] + 1; i < last[c]; i++ {
				seen[s[i]-'a'] = true
			}
			for _, v := range seen {
				if v {
					count++
				}
			}
		}
	}
	return count
}
