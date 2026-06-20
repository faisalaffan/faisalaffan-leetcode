# 0285 — Inorder Successor In Bst

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(h), Space: O(1)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #285: Inorder Successor in BST
// https://leetcode.com/problems/inorder-successor-in-bst/
// Difficulty: Medium [Paid]
// Time: O(h), Space: O(1)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var successor *TreeNode

	for root != nil {
		if p.Val < root.Val {
			successor = root
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return successor
}

func main() {
	root := &TreeNode{5, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, &TreeNode{7, nil, nil}}}
	p := root.Left
	fmt.Println(inorderSuccessor(root, p).Val)

	p2 := root.Right
	fmt.Println(inorderSuccessor(root, p2))

	p3 := root.Left.Right
	fmt.Println(inorderSuccessor(root, p3).Val)
}
```
