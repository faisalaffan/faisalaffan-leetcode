# 0226 — Invert Binary Tree

## Deskripsi

**Soal:** [0226. Invert Binary Tree](https://leetcode.com/problems/invert-binary-tree/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h)

**Algoritma:** —

**Fungsi Solusi:** `func InvertTree(root *TreeNode) *TreeNode`

## Solusi Go

```go
package main

// LeetCode #226: Invert Binary Tree
// https://leetcode.com/problems/invert-binary-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	return root
}

func preorder(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Print(root.Val, " ")
	preorder(root.Left)
	preorder(root.Right)
}

func main() {
	root := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{7, &TreeNode{6, nil, nil}, &TreeNode{9, nil, nil}}}
	preorder(InvertTree(root))
	fmt.Println()
}
```
