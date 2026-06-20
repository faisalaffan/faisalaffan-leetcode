# 0114 — Flatten Binary Tree To Linked List

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func flatten(root *TreeNode) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #114: Flatten Binary Tree to Linked List
// https://leetcode.com/problems/flatten-binary-tree-to-linked-list/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	curr := root
	for curr != nil {
		if curr.Left != nil {
			// Find rightmost node in left subtree
			prev := curr.Left
			for prev.Right != nil {
				prev = prev.Right
			}
			// Rewire
			prev.Right = curr.Right
			curr.Right = curr.Left
			curr.Left = nil
		}
		curr = curr.Right
	}
}

func printPreorder(root *TreeNode) {
	for root != nil {
		fmt.Printf("%d ", root.Val)
		if root.Left != nil {
			fmt.Print("(has left) ")
		}
		root = root.Right
	}
	fmt.Println()
}

func main() {
	// Test case 1: [1,2,5,3,4,null,6] -> [1,null,2,null,3,null,4,null,5,null,6]
	root := &TreeNode{Val: 1,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}},
		Right: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}}
	flatten(root)
	printPreorder(root) // 1 2 3 4 5 6

	// Test case 2
	root = nil
	flatten(root)
	printPreorder(root)

	// Test case 3: [0]
	root = &TreeNode{Val: 0}
	flatten(root)
	printPreorder(root) // 0
}

// Time: O(n) | Space: O(1)
```
