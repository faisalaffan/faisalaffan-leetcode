# 0208 — Implement Trie Prefix Tree

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

**Kompleksitas Waktu:** O(n) per operation, Space: O(total characters)  
**Kompleksitas Ruang:** O(total characters)

> **Untuk fresh graduate:** Kuasai dulu teknik **Trie** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

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
  // Loop linear O(n): iterasi setiap elemen
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
  // Loop linear O(n): iterasi setiap elemen
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
  // Loop linear O(n): iterasi setiap elemen
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
