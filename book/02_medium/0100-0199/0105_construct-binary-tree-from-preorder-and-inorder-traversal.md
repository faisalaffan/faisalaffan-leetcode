# 0105 — Construct Binary Tree From Preorder And Inorder Traversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func buildTree(preorder []int, inorder []int) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #105: Construct Binary Tree from Preorder and Inorder Traversal
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

  // Membuat map (HashMap) — pencarian O(1)
	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}

	var build func(preStart, preEnd, inStart, inEnd int) *TreeNode
	build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart > preEnd || inStart > inEnd {
			return nil
		}

		rootVal := preorder[preStart]
		root := &TreeNode{Val: rootVal}
		inIdx := inorderMap[rootVal]
		leftSize := inIdx - inStart

		root.Left = build(preStart+1, preStart+leftSize, inStart, inIdx-1)
		root.Right = build(preStart+leftSize+1, preEnd, inIdx+1, inEnd)
		return root
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}

func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInorder(root.Right)
}

func main() {
	// Test case 1
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	printInorder(root) // 9 3 15 20 7
	fmt.Println()

	// Test case 2
	root = buildTree([]int{-1}, []int{-1})
	printInorder(root) // -1
	fmt.Println()
}

// Time: O(n) | Space: O(n)
```
