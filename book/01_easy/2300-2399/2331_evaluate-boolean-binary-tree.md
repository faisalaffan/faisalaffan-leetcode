# 2331 — Evaluate Boolean Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func EvaluateBooleanBinaryTree(root *TreeNode) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


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
