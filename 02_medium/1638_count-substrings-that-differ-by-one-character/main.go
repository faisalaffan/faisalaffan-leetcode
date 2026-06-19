package main

// LeetCode #1638: Count Substrings That Differ by One Character
// https://leetcode.com/problems/count-substrings-that-differ-by-one-character/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(CountSubstrings("aba", "baba"))
	fmt.Println(CountSubstrings("ab", "bb"))
	fmt.Println(CountSubstrings("abe", "bbc"))
}

func CountSubstrings(s string, t string) int {
	// Time: O(N*M*min(N,M)), Space: O(1)
	count := 0

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			diff := 0
			k := 0
			for i+k < len(s) && j+k < len(t) && diff <= 1 {
				if s[i+k] != t[j+k] {
					diff++
				}
				if diff == 1 {
					count++
				}
				k++
			}
		}
	}

	return count
}
