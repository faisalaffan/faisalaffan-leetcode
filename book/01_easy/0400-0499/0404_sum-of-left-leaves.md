# 0404 — Sum Of Left Leaves

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func SumOfLeftLeaves(root *TreeNode) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n), Space: O(h)  |  **Ruang:** O(h)


## 💻 Solusi Go

```go
package main

// LeetCode #404: Sum of Left Leaves
// https://leetcode.com/problems/sum-of-left-leaves/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(h)
func SumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sum := 0
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		sum += root.Left.Val
	}
	sum += SumOfLeftLeaves(root.Left)
	sum += SumOfLeftLeaves(root.Right)
	return sum
}

func main() {
	// Test case 1: [3,9,20,null,null,15,7]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(SumOfLeftLeaves(root1))

	// Test case 2: [1]
	root2 := &TreeNode{Val: 1}
	fmt.Println(SumOfLeftLeaves(root2))
}
```
