# 1804 — Implement Trie Ii Prefix Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor() Trie
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Trie, Prefix Sum

**Kompleksitas Waktu:** O(L) per operation, Space: O(total characters)  
**Kompleksitas Ruang:** O(total characters)

> **Untuk fresh graduate:** Kuasai dulu teknik **Trie** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
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
```
