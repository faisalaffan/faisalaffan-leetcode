# 0776 — Split Bst

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func splitBST(root *TreeNode, target int) (*TreeNode, *TreeNode)`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(h) where h is tree height  |  **Ruang:** O(h)


## 💻 Solusi Go

```go
package main

// LeetCode #776: Split BST
// https://leetcode.com/problems/split-bst/
// Difficulty: Medium [Paid]
// Time: O(h) where h is tree height
// Space: O(h)

import "fmt"

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 6}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 5}
	root.Right.Right = &TreeNode{Val: 7}

	left, right := splitBST(root, 4)
	fmt.Println(left.Val)
	fmt.Println(right.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func splitBST(root *TreeNode, target int) (*TreeNode, *TreeNode) {
	if root == nil {
		return nil, nil
	}

	if root.Val <= target {
		left, right := splitBST(root.Right, target)
		root.Right = left
		return root, right
	}

	left, right := splitBST(root.Left, target)
	root.Left = right
	return left, root
}
```
