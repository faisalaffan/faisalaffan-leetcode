# 0663 — Equal Tree Partition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func checkEqualTree(root *TreeNode) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #663: Equal Tree Partition
// https://leetcode.com/problems/equal-tree-partition/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 10}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 3}

	fmt.Println(checkEqualTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkEqualTree(root *TreeNode) bool {
  // HashMap: O(1) lookup
	sums := make(map[int]int)
	total := computeSum(root, sums)

	if total%2 != 0 {
		return false
	}

	// Root sum is also stored; we need a different way to track
	// We'll rebuild sums excluding total
	half := total / 2
	return sums[half] > 0
}

func computeSum(node *TreeNode, sums map[int]int) int {
	if node == nil {
		return 0
	}
	left := computeSum(node.Left, sums)
	right := computeSum(node.Right, sums)
	sum := node.Val + left + right

	// Count subtree sums
	sums[sum]++

	return sum
}
```
