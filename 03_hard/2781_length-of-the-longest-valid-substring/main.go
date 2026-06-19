package main

// LeetCode #2781: Length of the Longest Valid Substring
// https://leetcode.com/problems/length-of-the-longest-valid-substring/
// Difficulty: Hard
//
// Trie + sliding window. Build a trie from reversed forbidden words. For each
// right bound, advance the left bound to exclude any forbidden word ending at
// right. O(N * maxForbiddenLen * alphabet) time, O(total forbidden chars) space.

import "fmt"

type trieNode struct {
	children [26]*trieNode
	isEnd    bool
}

func longestValidSubstring(word string, forbidden []string) int {
	root := &trieNode{}
	for _, f := range forbidden {
		if len(f) > len(word) {
			continue
		}
		node := root
		for i := len(f) - 1; i >= 0; i-- {
			idx := f[i] - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &trieNode{}
			}
			node = node.children[idx]
		}
		node.isEnd = true
	}

	maxLen := 0
	left := 0
	n := len(word)

	for right := 0; right < n; right++ {
		node := root
		for i := right; i >= left; i-- {
			idx := word[i] - 'a'
			if node.children[idx] == nil {
				break
			}
			node = node.children[idx]
			if node.isEnd {
				left = i + 1
				break
			}
		}
		cur := right - left + 1
		if cur > maxLen {
			maxLen = cur
		}
	}

	return maxLen
}

func main() {
	// Example: word="cbaaaabc", forbidden=["aaa","cb"] => 4
	fmt.Println(longestValidSubstring("cbaaaabc", []string{"aaa", "cb"}))
	// Single char matches forbidden
	fmt.Println(longestValidSubstring("a", []string{"a"}))
	// Single char no match
	fmt.Println(longestValidSubstring("a", []string{"b"}))
	// No forbidden
	fmt.Println(longestValidSubstring("abc", []string{"def", "gh"}))
	// Forbidden at start
	fmt.Println(longestValidSubstring("abcde", []string{"ab"}))
	// Forbidden at end
	fmt.Println(longestValidSubstring("abcde", []string{"de"}))
	// Multiple overlapping
	fmt.Println(longestValidSubstring("leetcode", []string{"leet", "code"}))
}
