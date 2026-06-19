package main

// LeetCode #3485: Longest Common Prefix of K Strings After Removal
// https://leetcode.com/problems/longest-common-prefix-of-k-strings-after-removal/
// Difficulty: Hard
//
// For each index i, find the length of the longest common prefix of any k
// strings after removing the i-th string from the array.
//
// Approach: Use a Trie to track frequencies. For each word, temporarily
// remove it, find longest common prefix among k remaining words.

import (
	"fmt"
)

func main() {
	// Example 1
	fmt.Println(longestCommonPrefix([]string{"abc", "abd", "ab"}, 2))
	// Example 2
	fmt.Println(longestCommonPrefix([]string{"a", "b", "c"}, 2))
	// Example 3: all same
	fmt.Println(longestCommonPrefix([]string{"abc", "abc", "abc"}, 2))
	// Edge: k = 1
	fmt.Println(longestCommonPrefix([]string{"abc", "def"}, 1))
}

type TrieNode struct {
	children [26]*TrieNode
	count    int
}

func longestCommonPrefix(words []string, k int) []int {
	n := len(words)
	if k > n {
		ans := make([]int, n)
		return ans
	}

	// Build trie with all words
	root := &TrieNode{}
	for _, w := range words {
		insert(root, w)
	}

	ans := make([]int, n)
	for i, w := range words {
		// Remove current word
		remove(root, w)
		// Find longest common prefix among k words
		ans[i] = findLCPK(root, k)
		// Reinsert
		insert(root, w)
	}

	return ans
}

func insert(root *TrieNode, w string) {
	node := root
	for _, ch := range w {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
		node.count++
	}
}

func remove(root *TrieNode, w string) {
	node := root
	var path []*TrieNode
	for _, ch := range w {
		idx := ch - 'a'
		path = append(path, node)
		node = node.children[idx]
		if node != nil {
			node.count--
		}
	}
}

func findLCPK(root *TrieNode, k int) int {
	// DFS to find longest prefix with >= k words
	var dfs func(node *TrieNode, depth int) int
	dfs = func(node *TrieNode, depth int) int {
		maxDepth := depth - 1 // depth before moving to child
		for i := 0; i < 26; i++ {
			if node.children[i] != nil && node.children[i].count >= k {
				d := dfs(node.children[i], depth+1)
				if d > maxDepth {
					maxDepth = d
				}
			}
		}
		return maxDepth
	}

	return dfs(root, 0)
}
