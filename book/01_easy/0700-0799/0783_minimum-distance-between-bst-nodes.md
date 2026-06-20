# 0783 — Minimum Distance Between Bst Nodes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func minDiffInBST(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #783: Minimum Distance Between BST Nodes
// https://leetcode.com/problems/minimum-distance-between-bst-nodes/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 6},
	}
	fmt.Println(minDiffInBST(root)) // 1
}

// minDiffInBST finds the minimum difference between values of any two nodes in a BST.
// Time: O(n). Space: O(n).
func minDiffInBST(root *TreeNode) int {
	minDiff := math.MaxInt32
	var prev *int
	inorder(root, &prev, &minDiff)
	return minDiff
}

func inorder(node *TreeNode, prev **int, minDiff *int) {
	if node == nil {
		return
	}
	inorder(node.Left, prev, minDiff)
	if *prev != nil {
		diff := node.Val - **prev
		if diff < *minDiff {
			*minDiff = diff
		}
	}
	*prev = &node.Val
	inorder(node.Right, prev, minDiff)
}
```
