# 0156 — Binary Tree Upside Down

## Deskripsi

**Soal:** [0156. Binary Tree Upside Down](https://leetcode.com/problems/binary-tree-upside-down/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(h) for recursion  
**Kompleksitas Ruang:** O(h) for recursion

**Algoritma:** —

**Fungsi Solusi:** `func upsideDownBinaryTree(root *TreeNode) *TreeNode`

## Solusi Go

```go
package main

// LeetCode #156: Binary Tree Upside Down
// https://leetcode.com/problems/binary-tree-upside-down/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(h) for recursion

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	newRoot := upsideDownBinaryTree(root.Left)
	root.Left.Left = root.Right
	root.Left.Right = root
	root.Left = nil
	root.Right = nil

	return newRoot
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	newRoot := upsideDownBinaryTree(root)
	fmt.Println(newRoot.Val)

	root2 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	newRoot2 := upsideDownBinaryTree(root2)
	fmt.Println(newRoot2.Val)

	var root3 *TreeNode
	newRoot3 := upsideDownBinaryTree(root3)
	fmt.Println(newRoot3)
}
```
