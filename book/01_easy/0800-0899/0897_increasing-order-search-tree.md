# 0897 — Increasing Order Search Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func printTree(root *TreeNode) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #897: Increasing Order Search Tree
// https://leetcode.com/problems/increasing-order-search-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
		},
	}
	result := increasingBST(root)
	// Print inorder to verify
	printTree(result) // 1 2 3 4 5 6 7 8 9
}

func printTree(root *TreeNode) {
	for root != nil {
		fmt.Print(root.Val, " ")
		root = root.Right
	}
	fmt.Println()
}

// increasingBST rearranges the BST into increasing order (only right children).
// Time: O(n). Space: O(n).
func increasingBST(root *TreeNode) *TreeNode {
	var newRoot, prev *TreeNode
	inorderTree(root, &newRoot, &prev)
	return newRoot
}

func inorderTree(node *TreeNode, newRoot **TreeNode, prev **TreeNode) {
	if node == nil {
		return
	}
	inorderTree(node.Left, newRoot, prev)
	if *newRoot == nil {
		*newRoot = node
	}
	if *prev != nil {
		(*prev).Right = node
	}
	node.Left = nil
	*prev = node
	inorderTree(node.Right, newRoot, prev)
}
```
