# 0559 — Maximum Depth Of N Ary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaximumDepthOfNAryTree(root *Node) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(h)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #559: Maximum Depth of N-ary Tree
// https://leetcode.com/problems/maximum-depth-of-n-ary-tree/
// Difficulty: Easy

import "fmt"

// Node represents an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// Time: O(n), Space: O(h)
func MaximumDepthOfNAryTree(root *Node) int {
	if root == nil {
		return 0
	}
	maxDepth := 0
	for _, child := range root.Children {
		if depth := MaximumDepthOfNAryTree(child); depth > maxDepth {
			maxDepth = depth
		}
	}
	return maxDepth + 1
}

func main() {
	// Test: [1,null,3,2,4,null,5,6]
	root1 := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 3, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 2},
			{Val: 4},
		},
	}
	fmt.Println(MaximumDepthOfNAryTree(root1))

	// Test: single node
	root2 := &Node{Val: 1}
	fmt.Println(MaximumDepthOfNAryTree(root2))
}
```
