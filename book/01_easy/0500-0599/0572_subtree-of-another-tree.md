# 0572 — Subtree Of Another Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func isSameTree(p, q *TreeNode) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n), Space: O(h)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #572: Subtree of Another Tree
// https://leetcode.com/problems/subtree-of-another-tree/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// Time: O(m*n), Space: O(h)
func SubtreeOfAnotherTree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return SubtreeOfAnotherTree(root.Left, subRoot) || SubtreeOfAnotherTree(root.Right, subRoot)
}

func main() {
	// Test: root=[3,4,5,1,2], subRoot=[4,1,2]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 5},
	}
	sub1 := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(SubtreeOfAnotherTree(root1, sub1))

	// Test: root=[3,4,5,1,2,null,null,null,null,0], subRoot=[4,1,2]
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 0},
			},
		},
		Right: &TreeNode{Val: 5},
	}
	sub2 := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(SubtreeOfAnotherTree(root2, sub2))
}
```
