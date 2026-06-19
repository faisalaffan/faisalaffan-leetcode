package main

// LeetCode #792: Number of Matching Subsequences
// https://leetcode.com/problems/number-of-matching-subsequences/
// Difficulty: Medium
// Time: O(n + m * L) where n = len(s), m = len(words)
// Space: O(m)

import "fmt"

func main() {
	fmt.Println(numMatchingSubseq("abcde", []string{"a", "bb", "acd", "ace"}))
	fmt.Println(numMatchingSubseq("dsahjpjauf", []string{"ahjpjau", "ja", "ahbwzgqnuk", "tnmlanowax"}))
}

func numMatchingSubseq(s string, words []string) int {
	buckets := make([][]string, 26)
	for i := range buckets {
		buckets[i] = make([]string, 0)
	}

	for _, w := range words {
		buckets[w[0]-'a'] = append(buckets[w[0]-'a'], w)
	}

	count := 0

	for _, c := range s {
		idx := c - 'a'
		curr := buckets[idx]
		buckets[idx] = make([]string, 0)

		for _, w := range curr {
			if len(w) == 1 {
				count++
			} else {
				buckets[w[1]-'a'] = append(buckets[w[1]-'a'], w[1:])
			}
		}
	}

	return count
}
