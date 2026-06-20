# 1302 — Deepest Leaves Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func deepestLeavesSum(root *TreeNode) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, BFS

**Waktu:** O(n)  |  **Ruang:** O(h) where h = tree height

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

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
