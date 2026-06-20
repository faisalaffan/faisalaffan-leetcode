# 0222 — Count Complete Tree Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func CountNodes(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(log^2 n)  
**Kompleksitas Ruang:** O(log n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #222: Count Complete Tree Nodes
// https://leetcode.com/problems/count-complete-tree-nodes/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(log^2 n) | Space: O(log n)
func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := func(node *TreeNode) int {
		d := 0
		for node != nil {
			d++
			node = node.Left
		}
		return d
	}
	rightDepth := func(node *TreeNode) int {
		d := 0
		for node != nil {
			d++
			node = node.Right
		}
		return d
	}
	l, r := leftDepth(root), rightDepth(root)
	if l == r {
		return (1 << l) - 1
	}
	return 1 + CountNodes(root.Left) + CountNodes(root.Right)
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, &TreeNode{6, nil, nil}, nil}}
	fmt.Println(CountNodes(root))
	fmt.Println(CountNodes(nil))
}
```
