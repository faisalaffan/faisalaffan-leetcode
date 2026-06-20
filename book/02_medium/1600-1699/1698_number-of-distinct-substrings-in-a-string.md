# 1698 — Number Of Distinct Substrings In A String

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func countDistinct(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Trie

**Kompleksitas Waktu:** O(n^2), Space: O(n^2) using trie  
**Kompleksitas Ruang:** O(n^2) using trie

> **Untuk fresh graduate:** Kuasai dulu teknik **Trie** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1698: Number of Distinct Substrings in a String
// https://leetcode.com/problems/number-of-distinct-substrings-in-a-string/
// Difficulty: Medium [Paid]
// Time: O(n^2), Space: O(n^2) using trie

import "fmt"

type TrieNode struct {
	children [26]*TrieNode
}

func countDistinct(s string) int {
	root := &TrieNode{}
	count := 0

  // Loop linear O(n): iterasi setiap elemen
	for i := 0; i < len(s); i++ {
		node := root
		for j := i; j < len(s); j++ {
			idx := s[j] - 'a'
			if node.children[idx] == nil {
				node.children[idx] = &TrieNode{}
				count++
			}
			node = node.children[idx]
		}
	}
	return count
}

func main() {
	fmt.Println(countDistinct("aabbaba")) // Expected: 21
	fmt.Println(countDistinct("abcdef"))  // Expected: 21
	fmt.Println(countDistinct("a"))       // Expected: 1
}
```
