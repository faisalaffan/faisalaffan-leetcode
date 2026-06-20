# 0623 — Add One Row To Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func AddOneRow(root *TreeNode, val int, depth int) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #623: Add One Row to Tree
// https://leetcode.com/problems/add-one-row-to-tree/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 5}}}
	result := AddOneRow(root, 1, 2)
	printTree(result)
	fmt.Println()
}

func AddOneRow(root *TreeNode, val int, depth int) *TreeNode {
	if depth == 1 {
		return &TreeNode{Val: val, Left: root}
	}
	addRowDFS(root, val, depth, 1)
	return root
}

func addRowDFS(node *TreeNode, val int, depth int, curDepth int) {
	if node == nil {
		return
	}
	if curDepth == depth-1 {
		oldLeft, oldRight := node.Left, node.Right
		node.Left = &TreeNode{Val: val, Left: oldLeft}
		node.Right = &TreeNode{Val: val, Right: oldRight}
		return
	}
	addRowDFS(node.Left, val, depth, curDepth+1)
	addRowDFS(node.Right, val, depth, curDepth+1)
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}
```
