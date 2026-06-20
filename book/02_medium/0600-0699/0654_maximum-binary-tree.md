# 0654 — Maximum Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func constructMaximumBinaryTree(nums []int) *TreeNode`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n^2) worst case, O(n log n) average  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #654: Maximum Binary Tree
// https://leetcode.com/problems/maximum-binary-tree/
// Difficulty: Medium
// Time: O(n^2) worst case, O(n log n) average
// Space: O(n)

import "fmt"

func main() {
	root := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	fmt.Println(treeToSlice(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
  // Edge case: input kosong
	if len(nums) == 0 {
		return nil
	}

	maxIdx := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[maxIdx] {
			maxIdx = i
		}
	}

	root := &TreeNode{Val: nums[maxIdx]}
	root.Left = constructMaximumBinaryTree(nums[:maxIdx])
	root.Right = constructMaximumBinaryTree(nums[maxIdx+1:])

	return root
}

func treeToSlice(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result := []int{root.Val}
	result = append(result, treeToSlice(root.Left)...)
	result = append(result, treeToSlice(root.Right)...)
	return result
}
```
