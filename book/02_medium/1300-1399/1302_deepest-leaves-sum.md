# 1302 — Deepest Leaves Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func deepestLeavesSum(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, BFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h = tree height

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1302: Deepest Leaves Sum
// https://leetcode.com/problems/deepest-leaves-sum/
// Difficulty: Medium

// Sum values of the deepest leaves in a binary tree.

// Time: O(n)
// Space: O(h) where h = tree height

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSum := 0
		size := len(queue)

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			levelSum += node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		// If this is the last level, return its sum
		if len(queue) == 0 {
			return levelSum
		}
	}

	return 0
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 7}},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 8}},
		},
	}
	fmt.Printf("%d (expected: 15)\n", deepestLeavesSum(root))

	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Printf("%d (expected: 5)\n", deepestLeavesSum(root2))
}
```
