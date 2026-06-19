package main

// LeetCode #1698: Number of Distinct Substrings in a String
// https://leetcode.com/problems/number-of-distinct-substrings-in-a-string/
// Difficulty: Medium [Paid]
// Time: O(n^2), Space: O(n^2) using trie

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
}

func countDistinct(s string) int {
	root := &TrieNode{}
	count := 0

	for i := 0; i < len(s); i++ {
		node := root
		for j := i; j < len(s); j++ {
			idx := s[j] - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
				count++
			}
			node = node.children[idx]
		}
	}
	return count
}

func main() {
	fmt.Println(countDistinct("aabbaba")) // Expected: 21
	fmt.Println(countDistinct("abcdef"))  // Expected: 21
	fmt.Println(countDistinct("a"))       // Expected: 1
}
