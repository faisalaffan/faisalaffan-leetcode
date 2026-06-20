# 0545 — Boundary Of Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func BoundaryOfBinaryTree(root *TreeNode) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #545: Boundary of Binary Tree
// https://leetcode.com/problems/boundary-of-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
	fmt.Println(BoundaryOfBinaryTree(root))
}

func BoundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{root.Val}
	addLeftBoundary(root.Left, &result)
	addLeaves(root.Left, &result)
	addLeaves(root.Right, &result)
	addRightBoundary(root.Right, &result)
	return result
}

func addLeftBoundary(node *TreeNode, result *[]int) {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return
	}
	*result = append(*result, node.Val)
	if node.Left != nil {
		addLeftBoundary(node.Left, result)
	} else {
		addLeftBoundary(node.Right, result)
	}
}

func addRightBoundary(node *TreeNode, result *[]int) {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return
	}
	if node.Right != nil {
		addRightBoundary(node.Right, result)
	} else {
		addRightBoundary(node.Left, result)
	}
	*result = append(*result, node.Val) // post-order for reverse
}

func addLeaves(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*result = append(*result, node.Val)
		return
	}
	addLeaves(node.Left, result)
	addLeaves(node.Right, result)
}
```
