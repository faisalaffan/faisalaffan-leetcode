# 0814 — Binary Tree Pruning

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func printTree(root *TreeNode) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h is tree height

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #814: Binary Tree Pruning
// https://leetcode.com/problems/binary-tree-pruning/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1: [1,null,0,0,1] -> [1,null,0,null,1]
	root1 := &TreeNode{1, nil, &TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}}}
	r1 := BinaryTreePruning(root1)
	printTree(r1)
	fmt.Println()

	// Test case 2: [1,0,1,0,0,0,1] -> [1,null,1,null,1]
	root2 := &TreeNode{1,
		&TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}},
		&TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}},
	}
	r2 := BinaryTreePruning(root2)
	printTree(r2)
	fmt.Println()

	// Test case 3: nil tree
	r3 := BinaryTreePruning(nil)
	printTree(r3)
	fmt.Println()
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("null")
		return
	}
	fmt.Print(root.Val)
	if root.Left != nil || root.Right != nil {
		fmt.Print(" ")
		printTree(root.Left)
		fmt.Print(" ")
		printTree(root.Right)
	}
}

// Time: O(n) | Space: O(h) where h is tree height
func BinaryTreePruning(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left = BinaryTreePruning(root.Left)
	root.Right = BinaryTreePruning(root.Right)

	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}

	return root
}
```
