# 0513 — Find Bottom Left Tree Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindBottomLeftTreeValue(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #513: Find Bottom Left Tree Value
// https://leetcode.com/problems/find-bottom-left-tree-value/
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
	// Test: root = [2,1,3]
	root1 := &TreeNode{Val: 2}
	root1.Left = &TreeNode{Val: 1}
	root1.Right = &TreeNode{Val: 3}
	fmt.Println(FindBottomLeftTreeValue(root1))

	// Test: root = [1,2,3,4,null,5,6,null,null,7]
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}
	root2.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 5, Left: &TreeNode{Val: 7}}, Right: &TreeNode{Val: 6}}
	fmt.Println(FindBottomLeftTreeValue(root2))
}

func FindBottomLeftTreeValue(root *TreeNode) int {
	queue := []*TreeNode{root}
	var leftmost int

	for len(queue) > 0 {
		levelSize := len(queue)
		leftmost = queue[0].Val
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return leftmost
}
```
