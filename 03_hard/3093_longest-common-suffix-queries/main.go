package main

// LeetCode #3093: Longest Common Suffix Queries
// https://leetcode.com/problems/longest-common-suffix-queries/
// Difficulty: Hard
// Time: O((N + Q) * M) | Space: O(N * M)

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	length   int // shortest container word length passing through this node
	idx      int // index of that shortest word
}

func stringIndices(wordsContainer []string, wordsQuery []string) []int {
	root := &TrieNode{length: 1 << 30, idx: -1}

	// Insert each container word into trie (reversed)
	for i, w := range wordsContainer {
		node := root
		// Update root with shortest word
		if len(w) < node.length {
			node.length = len(w)
			node.idx = i
		}
		// Traverse reversed word
		for k := len(w) - 1; k >= 0; k-- {
			c := int(w[k] - 'a')
			if node.children[c] == nil {
				node.children[c] = &TrieNode{length: 1 << 30, idx: -1}
			}
			node = node.children[c]
			if len(w) < node.length {
				node.length = len(w)
				node.idx = i
			}
		}
	}

	// Query each word
	ans := make([]int, len(wordsQuery))
	for qi, q := range wordsQuery {
		node := root
		for k := len(q) - 1; k >= 0; k-- {
			c := int(q[k] - 'a')
			if node.children[c] == nil {
				break
			}
			node = node.children[c]
		}
		ans[qi] = node.idx
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", stringIndices(
		[]string{"abcd", "bcd", "xbcd"},
		[]string{"bcd", "ab", "cd"},
	))
	// Expected: [1, 0, 1]  (Note: "ab" matches root, shortest word = "bcd" idx 1)

	// Test case 2
	fmt.Println("Test 2:", stringIndices(
		[]string{"abcd", "bcd", "xbcd"},
		[]string{"cd", "bcd", "xyz"},
	))
	// Expected: [1, 1, 1]

	// Test case 3
	fmt.Println("Test 3:", stringIndices(
		[]string{"aaaa", "a"},
		[]string{"a"},
	))
	// Expected: [1]
}
