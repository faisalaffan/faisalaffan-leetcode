# 2331 — Evaluate Boolean Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func EvaluateBooleanBinaryTree(root *TreeNode) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2331: Evaluate Boolean Binary Tree
// https://leetcode.com/problems/evaluate-boolean-binary-tree/
// Difficulty: Easy
// Time O(n) | Space O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// full binary tree: OR(2, AND(3,1)) = true
	root1 := &TreeNode{2,
		&TreeNode{1, nil, nil},
		&TreeNode{3, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}},
	}
	fmt.Println(EvaluateBooleanBinaryTree(root1)) // true

	root2 := &TreeNode{0, nil, nil}
	fmt.Println(EvaluateBooleanBinaryTree(root2)) // false
}

func EvaluateBooleanBinaryTree(root *TreeNode) bool {
	switch root.Val {
	case 0:
		return false
	case 1:
		return true
	case 2:
		return EvaluateBooleanBinaryTree(root.Left) || EvaluateBooleanBinaryTree(root.Right)
	default: // 3
		return EvaluateBooleanBinaryTree(root.Left) && EvaluateBooleanBinaryTree(root.Right)
	}
}
```
