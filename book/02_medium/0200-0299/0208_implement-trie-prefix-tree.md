# 0208 — Implement Trie Prefix Tree

## Deskripsi

**Soal:** [0208. Implement Trie Prefix Tree](https://leetcode.com/problems/implement-trie-prefix-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) per operation, Space: O(total characters)  
**Kompleksitas Ruang:** O(total characters)

**Algoritma:** Trie (pohon awalan)

**Fungsi Solusi:** `func Constructor() Trie`

## Solusi Go

```go
package main

// LeetCode #208: Implement Trie (Prefix Tree)
// https://leetcode.com/problems/implement-trie-prefix-tree/
// Difficulty: Medium
// Time: O(n) per operation, Space: O(total characters)

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{&TrieNode{}}
}

func (this *Trie) Insert(word string) {
	node := this.root
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &TrieNode{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(word); i++ {
		idx := word[i] - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return node.isEnd
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
  // Loop standar: indeks 0 sampai n-1
	for i := 0; i < len(prefix); i++ {
		idx := prefix[i] - 'a'
		if node.children[idx] == nil {
			return false
		}
		node = node.children[idx]
	}
	return true
}

func main() {
	trie := Constructor()
	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))
	fmt.Println(trie.Search("app"))
	fmt.Println(trie.StartsWith("app"))
	trie.Insert("app")
	fmt.Println(trie.Search("app"))
}
```
