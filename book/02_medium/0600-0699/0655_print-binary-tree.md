# 0655 — Print Binary Tree

## Deskripsi

**Soal:** [0655. Print Binary Tree](https://leetcode.com/problems/print-binary-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * h) where h is height  
**Kompleksitas Ruang:** O(n * h)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #655: Print Binary Tree
// https://leetcode.com/problems/print-binary-tree/
// Difficulty: Medium
// Time: O(n * h) where h is height
// Space: O(n * h)

import (
	"fmt"
	"strconv"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}

	fmt.Println(printTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	height := getHeight(root)
	width := (1 << height) - 1
  // Membuat slice 2D untuk DP/tabel
	result := make([][]string, height)
  // Iterasi seluruh elemen
	for i := range result {
		result[i] = make([]string, width)
	}

	fill(root, result, 0, 0, width-1)
	return result
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}

func fill(root *TreeNode, result [][]string, level int, left int, right int) {
	if root == nil {
		return
	}
	mid := (left + right) / 2
	result[level][mid] = strconv.Itoa(root.Val)
	fill(root.Left, result, level+1, left, mid-1)
	fill(root.Right, result, level+1, mid+1, right)
}
```
