package main

// LeetCode #336: Palindrome Pairs
// https://leetcode.com/problems/palindrome-pairs/
// Difficulty: Hard

import "fmt"

func palindromePairs(words []string) [][]int {
	// Build trie of reversed words
	type trieNode struct {
		child [26]*trieNode
		idx   int // index of word ending here, -1 if none
	}

	root := &trieNode{idx: -1}

	// Insert reversed word into trie
	for i, w := range words {
		node := root
		for j := len(w) - 1; j >= 0; j-- {
			c := w[j] - 'a'
			if node.child[c] == nil {
				node.child[c] = &trieNode{idx: -1}
			}
			node = node.child[c]
		}
		node.idx = i
	}

	isPalindrome := func(s string, l, r int) bool {
		for l < r {
			if s[l] != s[r] {
				return false
			}
			l++
			r--
		}
		return true
	}

	result := [][]int{}

	for i, w := range words {
		node := root

		// Check if remaining part of current word is palindrome and trie has a word ending here
		for j := 0; j < len(w); j++ {
			if node.idx != -1 && node.idx != i {
				if isPalindrome(w, j, len(w)-1) {
					result = append(result, []int{i, node.idx})
				}
			}
			c := w[j] - 'a'
			if node.child[c] == nil {
				node = nil
				break
			}
			node = node.child[c]
		}

		if node == nil {
			continue
		}

		// Exact match: reversed word completely matches current word
		if node.idx != -1 && node.idx != i {
			result = append(result, []int{i, node.idx})
		}

		// Check remaining trie paths where remaining prefix is palindrome
		// (current word is shorter than the trie word)
		var dfs func(*trieNode, []byte)
		dfs = func(n *trieNode, prefix []byte) {
			for c := 0; c < 26; c++ {
				if n.child[c] != nil {
					prefix = append(prefix, byte('a'+c))
					if n.child[c].idx != -1 && n.child[c].idx != i {
						if isPalindrome(string(prefix), 0, len(prefix)-1) {
							result = append(result, []int{i, n.child[c].idx})
						}
					}
					dfs(n.child[c], prefix)
					prefix = prefix[:len(prefix)-1]
				}
			}
		}
		dfs(node, []byte{})
	}

	return result
}

func main() {
	// Example 1
	fmt.Println(palindromePairs([]string{"abcd", "dcba", "lls", "s", "sssll"}))
	// [[0 1] [1 0] [3 2] [2 4]]

	// Example 2
	fmt.Println(palindromePairs([]string{"bat", "tab", "cat"}))
	// [[0 1] [1 0]]

	// Example 3
	fmt.Println(palindromePairs([]string{"a", ""}))
	// [[0 1] [1 0]]
}
