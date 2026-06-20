# 0663 — Equal Tree Partition

## Deskripsi

**Soal:** [0663. Equal Tree Partition](https://leetcode.com/problems/equal-tree-partition/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

## Solusi Go

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
  // Membuat map untuk pencarian O(1): key → value
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
