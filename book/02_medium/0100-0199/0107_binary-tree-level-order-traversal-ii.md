# 0107 — Binary Tree Level Order Traversal Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func levelOrderBottom(root *TreeNode) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #107: Binary Tree Level Order Traversal II
// https://leetcode.com/problems/binary-tree-level-order-traversal-ii/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
  // Alokasi slice integer
		level := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			level[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append([][]int{level}, result...)
		queue = queue[levelSize:]
	}

	return result
}

func main() {
	// Test case 1: [3,9,20,null,null,15,7] -> [[15,7],[9,20],[3]]
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(levelOrderBottom(root))

	// Test case 2
	fmt.Println(levelOrderBottom(nil)) // []

	// Test case 3: [1] -> [[1]]
	fmt.Println(levelOrderBottom(&TreeNode{Val: 1})) // [[1]]
}

// Time: O(n) | Space: O(n)
```
