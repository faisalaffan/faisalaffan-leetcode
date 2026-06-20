# 0099 — Recover Binary Search Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func recoverTree(root *TreeNode) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #99: Recover Binary Search Tree
// https://leetcode.com/problems/recover-binary-search-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil && prev.Val > node.Val {
			if first == nil {
				first = prev
			}
			second = node
		}
		prev = node
		inorder(node.Right)
	}

	inorder(root)
	first.Val, second.Val = second.Val, first.Val
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	printTree(root.Left)
	fmt.Printf("%d ", root.Val)
	printTree(root.Right)
}

func main() {
	// Test case 1: [1,3,null,null,2] -> [3,1,null,null,2]
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 3, Right: &TreeNode{Val: 2}}}
	recoverTree(root)
	printTree(root) // 1 2 3
	fmt.Println()

	// Test case 2: [3,1,4,null,null,2] -> [2,1,4,null,null,3]
	root = &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 2}}}
	recoverTree(root)
	printTree(root) // 1 2 3 4
	fmt.Println()
}

// Time: O(n) | Space: O(n)
```
