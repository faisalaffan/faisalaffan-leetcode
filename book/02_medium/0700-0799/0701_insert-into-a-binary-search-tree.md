# 0701 — Insert Into A Binary Search Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func insertIntoBST(root *TreeNode, val int) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(h) where h is tree height  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #701: Insert into a Binary Search Tree
// https://leetcode.com/problems/insert-into-a-binary-search-tree/
// Difficulty: Medium
// Time: O(h) where h is tree height
// Space: O(1)

import "fmt"

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}

	result := insertIntoBST(root, 5)
	fmt.Println(result.Val)
	fmt.Println(result.Right.Left.Val)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return &TreeNode{Val: val}
	}

	curr := root
	for {
		if val < curr.Val {
			if curr.Left == nil {
				curr.Left = &TreeNode{Val: val}
				break
			}
			curr = curr.Left
		} else {
			if curr.Right == nil {
				curr.Right = &TreeNode{Val: val}
				break
			}
			curr = curr.Right
		}
	}

	return root
}
```
