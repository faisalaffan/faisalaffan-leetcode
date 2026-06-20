# 2689 — Extract Kth Character From The Rope Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func ExtractKthCharacterFromTheRopeTree(root *RopeNode, k int) byte
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2689: Extract Kth Character From The Rope Tree
// https://leetcode.com/problems/extract-kth-character-from-the-rope-tree/
// Difficulty: Easy [Paid]
// Time: O(n) | Space: O(h)
// Note: A rope tree is a binary tree where leaves contain substrings and internal nodes concatenate children.

import "fmt"

func main() {
	// Rope tree representing "abc" + "defg" = "abcdefg"
	root := &RopeNode{left: &RopeNode{val: "abc"}, right: &RopeNode{val: "defg"}}
	fmt.Println(string(ExtractKthCharacterFromTheRopeTree(root, 5))) // 'e'
}

type RopeNode struct {
	val   string
	left  *RopeNode
	right *RopeNode
}

func ExtractKthCharacterFromTheRopeTree(root *RopeNode, k int) byte {
	var dfs func(*RopeNode) string
	dfs = func(node *RopeNode) string {
		if node == nil {
			return ""
		}
		if node.left == nil && node.right == nil {
			return node.val
		}
		return dfs(node.left) + dfs(node.right)
	}
	return dfs(root)[k-1]
}
```
