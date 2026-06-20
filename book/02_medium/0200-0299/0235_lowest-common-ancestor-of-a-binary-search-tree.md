# 0235 — Lowest Common Ancestor Of A Binary Search Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search

**Waktu:** O(h), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #235: Lowest Common Ancestor of a Binary Search Tree
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// Difficulty: Medium
// Time: O(h), Space: O(1)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}

func main() {
	root := &TreeNode{6, &TreeNode{2, &TreeNode{0, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}}, &TreeNode{8, &TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}}}
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right).Val)

	p := root.Left
	q := root.Left.Right
	fmt.Println(lowestCommonAncestor(root, p, q).Val)

	p2 := root.Left
	q2 := root.Left.Left
	fmt.Println(lowestCommonAncestor(root, p2, q2).Val)
}
```
