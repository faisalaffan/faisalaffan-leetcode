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

**Teknik yang digunakan:** HashMap, BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2196: Create Binary Tree From Descriptions
// https://leetcode.com/problems/create-binary-tree-from-descriptions/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func createBinaryTree(descriptions [][]int) *TreeNode {
  // Membuat map (HashMap) — pencarian O(1)
	children := make(map[int]bool)
  // Membuat map (HashMap) — pencarian O(1)
	nodes := make(map[int]*TreeNode)

	for _, desc := range descriptions {
		parent, child, isLeft := desc[0], desc[1], desc[2]
		if nodes[parent] == nil {
			nodes[parent] = &TreeNode{Val: parent}
		}
		if nodes[child] == nil {
			nodes[child] = &TreeNode{Val: child}
		}
		if isLeft == 1 {
			nodes[parent].Left = nodes[child]
		} else {
			nodes[parent].Right = nodes[child]
		}
		children[child] = true
	}

	for _, desc := range descriptions {
		if !children[desc[0]] {
			return nodes[desc[0]]
		}
	}
	return nil
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("null ")
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			fmt.Print("null ")
			continue
		}
		fmt.Printf("%d ", node.Val)
		queue = append(queue, node.Left, node.Right)
	}
	fmt.Println()
}

func main() {
	// Test case 1
	root1 := createBinaryTree([][]int{{20, 15, 1}, {20, 17, 0}, {50, 20, 1}, {50, 80, 0}})
	printTree(root1)

	// Test case 2
	root2 := createBinaryTree([][]int{{1, 2, 1}, {2, 3, 0}, {3, 4, 1}})
	printTree(root2)
}
```
