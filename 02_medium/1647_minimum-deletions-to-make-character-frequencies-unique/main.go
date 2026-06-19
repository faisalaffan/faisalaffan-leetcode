package main

// LeetCode #1647: Minimum Deletions to Make Character Frequencies Unique
// https://leetcode.com/problems/minimum-deletions-to-make-character-frequencies-unique/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MinDeletions("aab"))
	fmt.Println(MinDeletions("aaabbbcc"))
	fmt.Println(MinDeletions("ceabaacb"))
}

func MinDeletions(s string) int {
	// Time: O(N), Space: O(1)
	freq := make([]int, 26)
	for _, ch := range s {
		freq[ch-'a']++
	}

	used := make(map[int]bool)
	deletions := 0

	for _, f := range freq {
		for f > 0 && used[f] {
			f--
			deletions++
		}
		if f > 0 {
			used[f] = true
		}
	}

	return deletions
}
