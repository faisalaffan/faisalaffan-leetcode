# 0669 — Trim A Binary Search Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func trimBST(root *TreeNode, low int, high int) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #669: Trim a Binary Search Tree
// https://leetcode.com/problems/trim-a-binary-search-tree/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 2}

	result := trimBST(root, 1, 2)
	fmt.Println(result.Val)
	if result.Left != nil {
		fmt.Println(result.Left.Val)
	}
	if result.Right != nil {
		fmt.Println(result.Right.Val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
```
