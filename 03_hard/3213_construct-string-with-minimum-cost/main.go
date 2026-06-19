package main

// LeetCode #3213: Construct String with Minimum Cost
// https://leetcode.com/problems/construct-string-with-minimum-cost/
// Difficulty: Hard
//
// Trie + DP. Build a trie from words (storing min cost per node).
// For each position i in target (right to left), walk the trie to find
// all words matching target[i:j], update dp[i] = min(dp[i], cost + dp[j]).
// Return dp[0] or -1 if impossible.

import (
	"fmt"
	"math"
)

type trieNode struct {
	child [26]*trieNode
	cost  int
}

func newTrieNode() *trieNode {
	return &trieNode{cost: math.MaxInt32}
}

func main() {
	// Example 1: target="abcdef", words=["abdef","abc","d","def","ef"], costs=[100,1,1,10,5] => 7
	fmt.Println(minimumCost("abcdef", []string{"abdef", "abc", "d", "def", "ef"}, []int{100, 1, 1, 10, 5}))
	// Example 2: no match
	fmt.Println(minimumCost("xyz", []string{"ab", "cd"}, []int{5, 5}))
	// Example 3: single word
	fmt.Println(minimumCost("hello", []string{"hello", "world"}, []int{3, 7}))
	// Example 4: multiple same prefix words
	fmt.Println(minimumCost("aaa", []string{"a", "aa", "aaa"}, []int{5, 3, 1}))
	// Example 5: empty target
	fmt.Println(minimumCost("", []string{"a"}, []int{1}))
}

func minimumCost(target string, words []string, costs []int) int {
	if len(target) == 0 {
		return 0
	}

	root := newTrieNode()
	for i, w := range words {
		node := root
		for _, ch := range w {
			idx := ch - 'a'
			if node.child[idx] == nil {
				node.child[idx] = newTrieNode()
			}
			node = node.child[idx]
		}
		if costs[i] < node.cost {
			node.cost = costs[i]
		}
	}

	n := len(target)
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[n] = 0

	for i := n - 1; i >= 0; i-- {
		node := root
		for j := i; j < n; j++ {
			idx := target[j] - 'a'
			if node.child[idx] == nil {
				break
			}
			node = node.child[idx]
			if node.cost != math.MaxInt32 && dp[j+1] != math.MaxInt32 {
				candidate := node.cost + dp[j+1]
				if candidate < dp[i] {
					dp[i] = candidate
				}
			}
		}
	}

	if dp[0] == math.MaxInt32 {
		return -1
	}
	return dp[0]
}
