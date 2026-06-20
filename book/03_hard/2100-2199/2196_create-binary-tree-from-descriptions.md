# 2196 — Create Binary Tree From Descriptions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func createBinaryTree(descriptions [][]int) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2196: Create Binary Tree From Descriptions
// https://leetcode.com/problems/create-binary-tree-from-descriptions/
// Difficulty: Medium (listed as Hard)

import "fmt"

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// createBinaryTree builds a binary tree from descriptions.
// Each description: [parent, child, isLeft] where isLeft=1 means left child.
func createBinaryTree(descriptions [][]int) *TreeNode {
  // Membuat map (HashMap) — pencarian O(1)
	children := make(map[int]*TreeNode)
  // Membuat map (HashMap) — pencarian O(1)
	hasParent := make(map[int]bool)

	for _, d := range descriptions {
		parentVal, childVal, isLeft := d[0], d[1], d[2]

		// get or create parent
		parent, ok := children[parentVal]
		if !ok {
			parent = &TreeNode{Val: parentVal}
			children[parentVal] = parent
		}

		// get or create child
		child, ok := children[childVal]
		if !ok {
			child = &TreeNode{Val: childVal}
			children[childVal] = child
		}

		// set left/right
		if isLeft == 1 {
			parent.Left = child
		} else {
			parent.Right = child
		}

		hasParent[childVal] = true
		// ensure parent also tracked
		if !hasParent[parentVal] {
			hasParent[parentVal] = false
		}
	}

	// find root: node without a parent
	var root *TreeNode
	for val, node := range children {
		if !hasParent[val] {
			root = node
			break
		}
	}

	return root
}

func main() {
	// Example 1
	descriptions1 := [][]int{
		{20, 15, 1},
		{20, 17, 0},
		{50, 20, 1},
		{50, 80, 0},
		{80, 19, 1},
	}
	root1 := createBinaryTree(descriptions1)
	fmt.Println("Root value:", root1.Val) // Expected: 50

	// Example 2
	descriptions2 := [][]int{
		{1, 2, 1},
		{2, 3, 0},
		{3, 4, 1},
	}
	root2 := createBinaryTree(descriptions2)
	fmt.Println("Root value:", root2.Val) // Expected: 1
}
```
