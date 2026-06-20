# 0236 — Lowest Common Ancestor Of A Binary Tree

## Deskripsi

**Soal:** [0236. Lowest Common Ancestor Of A Binary Tree](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(h)  
**Kompleksitas Ruang:** O(h)

**Algoritma:** —

**Fungsi Solusi:** `func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode`

## Solusi Go

```go
package main

// LeetCode #236: Lowest Common Ancestor of a Binary Tree
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/
// Difficulty: Medium
// Time: O(n), Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}
	return right
}

func main() {
	root := &TreeNode{3, &TreeNode{5, &TreeNode{6, nil, nil}, &TreeNode{2, &TreeNode{7, nil, nil}, &TreeNode{4, nil, nil}}}, &TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{8, nil, nil}}}
	p, q := root.Left, root.Right
	fmt.Println(lowestCommonAncestor(root, p, q).Val)

	p, q = root.Left, root.Left.Right.Right
	fmt.Println(lowestCommonAncestor(root, p, q).Val)

	p, q = root.Left.Left, root.Left.Right
	fmt.Println(lowestCommonAncestor(root, p, q).Val)
}
```
