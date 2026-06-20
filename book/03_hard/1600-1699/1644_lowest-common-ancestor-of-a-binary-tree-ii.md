# 1644 — Lowest Common Ancestor Of A Binary Tree Ii

## Deskripsi

**Soal:** [1644. Lowest Common Ancestor Of A Binary Tree Ii](https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode`

## Solusi Go

```go
package main

import "fmt"

// LeetCode #1644: Lowest Common Ancestor of a Binary Tree II
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-ii/
// Difficulty: Medium (listed in Hard section)
//
// Given the root of a binary tree and two nodes p and q, return their lowest
// common ancestor. Unlike LCA I, p and q may not exist in the tree.
// If either node does not exist, return nil.

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor returns the LCA of p and q, or nil if either is missing.
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	// First pass: check existence of both nodes
	foundP := false
	foundQ := false

	var search func(*TreeNode)
	search = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node == p {
			foundP = true
		}
		if node == q {
			foundQ = true
		}
		search(node.Left)
		search(node.Right)
	}
	search(root)

	if !foundP || !foundQ {
		return nil
	}

	// Both exist, find LCA
	var lca func(*TreeNode) *TreeNode
	lca = func(node *TreeNode) *TreeNode {
		if node == nil || node == p || node == q {
			return node
		}
		left := lca(node.Left)
		right := lca(node.Right)
		if left != nil && right != nil {
			return node
		}
		if left != nil {
			return left
		}
		return right
	}

	return lca(root)
}

// BuildBT builds a binary tree from a level-order slice (-1 for nil).
func BuildBT(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == -1 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != -1 {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != -1 {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// findNode finds a node with given value in the tree.
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func main() {
	// Example 1:
	// Tree:
	//       3
	//      / \
	//     5   1
	//    / \  / \
	//   6  2 0  8
	//     / \
	//    7   4
	//
	// LCA of 5 and 1 = 3
	// LCA of 5 and 4 = 5
	// LCA of 5 and 999 (non-existent) = nil

	root := BuildBT([]int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4})

	node5 := findNode(root, 5)
	node1 := findNode(root, 1)
	node4 := findNode(root, 4)

	fmt.Println("LCA of 5 and 1:", lowestCommonAncestor(root, node5, node1)) // 3
	if lca := lowestCommonAncestor(root, node5, node1); lca != nil {
		fmt.Println("  Val:", lca.Val)
	}

	fmt.Println("LCA of 5 and 4:", lowestCommonAncestor(root, node5, node4)) // 5
	if lca := lowestCommonAncestor(root, node5, node4); lca != nil {
		fmt.Println("  Val:", lca.Val)
	}

	// Non-existent node
	fakeNode := &TreeNode{Val: 999}
	fmt.Println("LCA of 5 and 999 (fake):", lowestCommonAncestor(root, node5, fakeNode)) // nil

	// Self LCA
	fmt.Println("LCA of 5 and 5:", lowestCommonAncestor(root, node5, node5)) // 5
	if lca := lowestCommonAncestor(root, node5, node5); lca != nil {
		fmt.Println("  Val:", lca.Val)
	}

	// Empty tree
	fmt.Println("LCA nil:", lowestCommonAncestor(nil, node5, node1)) // nil
}
```
