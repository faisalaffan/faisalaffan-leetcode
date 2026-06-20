# 0530 — Minimum Absolute Difference In Bst

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func MinimumAbsoluteDifferenceInBst(root *TreeNode) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(h)  |  **Ruang:** O(h)


## 💻 Solusi Go

```go
package main

// LeetCode #530: Minimum Absolute Difference in BST
// https://leetcode.com/problems/minimum-absolute-difference-in-bst/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(h)
func MinimumAbsoluteDifferenceInBst(root *TreeNode) int {
	minDiff := math.MaxInt32
	var prev *int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil {
			diff := node.Val - *prev
			if diff < minDiff {
				minDiff = diff
			}
		}
		prev = &node.Val
		inorder(node.Right)
	}
	inorder(root)
	return minDiff
}

func main() {
	// Test: [4,2,6,1,3]
	root1 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 6},
	}
	fmt.Println(MinimumAbsoluteDifferenceInBst(root1))

	// Test: [1,0,48,null,null,12,49]
	root2 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 0},
		Right: &TreeNode{
			Val:   48,
			Left:  &TreeNode{Val: 12},
			Right: &TreeNode{Val: 49},
		},
	}
	fmt.Println(MinimumAbsoluteDifferenceInBst(root2))
}
```
