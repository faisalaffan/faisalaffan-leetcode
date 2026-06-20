# 0501 — Find Mode In Binary Search Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindModeInBinarySearchTree(root *TreeNode) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #501: Find Mode in Binary Search Tree
// https://leetcode.com/problems/find-mode-in-binary-search-tree/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(n)
func FindModeInBinarySearchTree(root *TreeNode) []int {
	var result []int
	maxCount, currentCount := 0, 0
	var prev *int

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil && *prev == node.Val {
			currentCount++
		} else {
			currentCount = 1
		}
		prev = &node.Val
		if currentCount > maxCount {
			maxCount = currentCount
			result = []int{node.Val}
		} else if currentCount == maxCount {
			result = append(result, node.Val)
		}
		inorder(node.Right)
	}
	inorder(root)
	return result
}

func main() {
	// Test: [1,null,2,2]
	root1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 2},
		},
	}
	fmt.Println(FindModeInBinarySearchTree(root1))

	// Test: [0]
	root2 := &TreeNode{Val: 0}
	fmt.Println(FindModeInBinarySearchTree(root2))
}
```
