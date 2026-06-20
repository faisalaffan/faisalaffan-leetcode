# 0530 — Minimum Absolute Difference In Bst

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumAbsoluteDifferenceInBst(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(h)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #530: Minimum Absolute Difference in BST
// https://leetcode.com/problems/minimum-absolute-difference-in-bst/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(h)
func MinimumAbsoluteDifferenceInBst(root *TreeNode) int {
	minDiff := math.MaxInt32
	var prev *int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil {
			diff := node.Val - *prev
			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = &node.Val
		inorder(node.Right)
	}
	inorder(root)
	return minDiff
}

func main() {
	// Test: [4,2,6,1,3]
	root1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 6},
	}
	fmt.Println(MinimumAbsoluteDifferenceInBst(root1))

	// Test: [1,0,48,null,null,12,49]
	root2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{
			Val:   48,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 49},
		},
	}
	fmt.Println(MinimumAbsoluteDifferenceInBst(root2))
}
```
