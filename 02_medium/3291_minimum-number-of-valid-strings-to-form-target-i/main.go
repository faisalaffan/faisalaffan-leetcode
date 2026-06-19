package main

// LeetCode #3291: Minimum Number of Valid Strings to Form Target I
// https://leetcode.com/problems/minimum-number-of-valid-strings-to-form-target-i/
// Difficulty: Medium
// Time: O(n * L) Space: O(total_chars + n) where L = average prefix length

import (
	"fmt"
)

func main() {
	fmt.Println(minValidStrings([]string{"abc", "aaaaa", "bcdef"}, "aabcdabc")) // 3
	fmt.Println(minValidStrings([]string{"ab", "bc", "cd"}, "abc"))            // 2
	fmt.Println(minValidStrings([]string{"a", "b", "c"}, "xyz"))               // -1
}

type trieNode struct {
	children [26]*trieNode
}

func minValidStrings(words []string, target string) int {
	root := &trieNode{}
	for _, w := range words {
		node := root
		for _, ch := range w {
			idx := ch - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &trieNode{}
			}
			node = node.children[idx]
		}
	}

	n := len(target)
	inf := int(1e9)
	dp := make([]int, n+1)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		if dp[i] == inf {
			continue
		}
		node := root
		for j := i; j < n; j++ {
			idx := target[j] - 'a'
			if node.children[idx] == nil {
				break
			}
			node = node.children[idx]
			if dp[i]+1 < dp[j+1] {
				dp[j+1] = dp[i] + 1
			}
		}
	}

	if dp[n] == inf {
		return -1
	}
	return dp[n]
}
