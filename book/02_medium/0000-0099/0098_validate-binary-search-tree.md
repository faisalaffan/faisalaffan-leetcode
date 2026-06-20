# 0098 — Validate Binary Search Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func isValidBST(root *TreeNode) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #98: Validate Binary Search Tree
// https://leetcode.com/problems/validate-binary-search-tree/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}
	if int64(node.Val) <= min || int64(node.Val) >= max {
		return false
	}
	return validate(node.Left, min, int64(node.Val)) && validate(node.Right, int64(node.Val), max)
}

func main() {
	// Test case 1: [2,1,3] -> true
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	fmt.Println(isValidBST(root)) // true

	// Test case 2: [5,1,4,null,null,3,6] -> false
	root = &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	fmt.Println(isValidBST(root)) // false

	// Test case 3: [2,2,2] -> false
	root = &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}
	fmt.Println(isValidBST(root)) // false
}

// Time: O(n) | Space: O(n)
```
