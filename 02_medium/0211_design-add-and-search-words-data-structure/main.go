package main

// LeetCode #211: Design Add and Search Words Data Structure
// https://leetcode.com/problems/design-add-and-search-words-data-structure/
// Difficulty: Medium
// Time: O(n) for add, O(26^m) worst case for search with wildcards, Space: O(total chars)

import "fmt"

type WordDictionaryNode struct {
	children [26]*WordDictionaryNode
	isEnd    bool
}

type WordDictionary struct {
	root *WordDictionaryNode
}

func Constructor() WordDictionary {
	return WordDictionary{&WordDictionaryNode{}}
}

func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &WordDictionaryNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (this *WordDictionary) Search(word string) bool {
	return this.search(this.root, word, 0)
}

func (this *WordDictionary) search(node *WordDictionaryNode, word string, idx int) bool {
	if node == nil {
		return false
	}
	if idx == len(word) {
		return node.isEnd
	}

	if word[idx] == '.' {
		for i := 0; i < 26; i++ {
			if this.search(node.children[i], word, idx+1) {
				return true
			}
		}
		return false
	}

	child := node.children[word[idx]-'a']
	return this.search(child, word, idx+1)
}

func main() {
	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	fmt.Println(wd.Search("pad"))
	fmt.Println(wd.Search("bad"))
	fmt.Println(wd.Search(".ad"))
	fmt.Println(wd.Search("b.."))
}
