# 0637 — Average Of Levels In Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func AverageOfLevelsInBinaryTree(root *TreeNode) []float64
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, BFS

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #637: Average of Levels in Binary Tree
// https://leetcode.com/problems/average-of-levels-in-binary-tree/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(n)
func AverageOfLevelsInBinaryTree(root *TreeNode) []float64 {
	var result []float64
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, float64(sum)/float64(levelSize))
	}
	return result
}

func main() {
	// Test: [3,9,20,null,null,15,7]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(AverageOfLevelsInBinaryTree(root1))

	// Test: [3,9,20,15,7]
	root2 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
		Right: &TreeNode{Val: 20},
	}
	fmt.Println(AverageOfLevelsInBinaryTree(root2))
}
```
