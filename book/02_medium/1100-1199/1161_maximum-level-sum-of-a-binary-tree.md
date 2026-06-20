# 1161 — Maximum Level Sum Of A Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxLevelSum(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1161: Maximum Level Sum of a Binary Tree
// https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
// Difficulty: Medium

// BFS level order traversal, track sum per level.

// Time: O(n)
// Space: O(n)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	maxSum := root.Val
	maxLevel := 1
	currentLevel := 0

	for len(queue) > 0 {
		currentLevel++
		size := len(queue)
		sum := 0

		for i := 0; i < size; i++ {
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

		if sum > maxSum {
			maxSum = sum
			maxLevel = currentLevel
		}
	}

	return maxLevel
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: -8},
		},
		Right: &TreeNode{Val: 0},
	}
	fmt.Printf("%d (expected: 2)\n", maxLevelSum(root))

	root2 := &TreeNode{
		Val:  -100,
		Left: &TreeNode{Val: -200},
		Right: &TreeNode{
			Val:   -300,
			Left:  &TreeNode{Val: -20},
			Right: &TreeNode{Val: -5},
		},
	}
	fmt.Printf("%d (expected: 3)\n", maxLevelSum(root2))
}
```
