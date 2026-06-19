package main

// LeetCode #1804: Implement Trie II (Prefix Tree)
// https://leetcode.com/problems/implement-trie-ii-prefix-tree/
// Difficulty: Medium [Paid]
// Time: O(L) per operation, Space: O(total characters)

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	wordCnt  int
	prefixCnt int
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{&TrieNode{}}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
		node.prefixCnt++
	}
	node.wordCnt++
}

func (t *Trie) CountWordsEqualTo(word string) int {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return 0
		}
		node = node.children[idx]
	}
	return node.wordCnt
}

func (t *Trie) CountWordsStartingWith(prefix string) int {
	node := t.root
	for _, ch := range prefix {
		idx := ch - 'a'
		if node.children[idx] == nil {
			return 0
		}
		node = node.children[idx]
	}
	return node.prefixCnt
}

func (t *Trie) Erase(word string) {
	node := t.root
	for _, ch := range word {
		idx := ch - 'a'
		node = node.children[idx]
		node.prefixCnt--
	}
	node.wordCnt--
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	trie.Insert("apple")
	fmt.Println(trie.CountWordsEqualTo("apple"))  // Expected: 2
	fmt.Println(trie.CountWordsStartingWith("app")) // Expected: 2
	trie.Erase("apple")
	fmt.Println(trie.CountWordsEqualTo("apple"))  // Expected: 1
	fmt.Println(trie.CountWordsStartingWith("app")) // Expected: 1
	trie.Erase("apple")
	fmt.Println(trie.CountWordsEqualTo("apple"))  // Expected: 0
}
