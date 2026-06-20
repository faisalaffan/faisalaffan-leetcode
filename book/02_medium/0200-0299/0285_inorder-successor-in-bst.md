# 0285 — Inorder Successor In Bst

## Deskripsi

**Soal:** [0285. Inorder Successor In Bst](https://leetcode.com/problems/inorder-successor-in-bst/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(h), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode`

## Solusi Go

```go
package main

// LeetCode #285: Inorder Successor in BST
// https://leetcode.com/problems/inorder-successor-in-bst/
// Difficulty: Medium [Paid]
// Time: O(h), Space: O(1)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var successor *TreeNode

	for root != nil {
		if p.Val < root.Val {
			successor = root
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return successor
}

func main() {
	root := &TreeNode{5, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, &TreeNode{7, nil, nil}}}
	p := root.Left
	fmt.Println(inorderSuccessor(root, p).Val)

	p2 := root.Right
	fmt.Println(inorderSuccessor(root, p2))

	p3 := root.Left.Right
	fmt.Println(inorderSuccessor(root, p3).Val)
}
```
