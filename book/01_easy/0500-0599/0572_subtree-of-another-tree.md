# 0572 — Subtree Of Another Tree

## Deskripsi

**Soal:** [0572. Subtree Of Another Tree](https://leetcode.com/problems/subtree-of-another-tree/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(m*n), Space: O(h)  
**Kompleksitas Ruang:** O(h)

**Algoritma:** —

**Fungsi Solusi:** `func isSameTree(p, q *TreeNode) bool`

## Solusi Go

```go
package main

// LeetCode #572: Subtree of Another Tree
// https://leetcode.com/problems/subtree-of-another-tree/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

// Time: O(m*n), Space: O(h)
func SubtreeOfAnotherTree(root, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return SubtreeOfAnotherTree(root.Left, subRoot) || SubtreeOfAnotherTree(root.Right, subRoot)
}

func main() {
	// Test: root=[3,4,5,1,2], subRoot=[4,1,2]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 5},
	}
	sub1 := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(SubtreeOfAnotherTree(root1, sub1))

	// Test: root=[3,4,5,1,2,null,null,null,null,0], subRoot=[4,1,2]
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:   2,
				Left:  &TreeNode{Val: 0},
			},
		},
		Right: &TreeNode{Val: 5},
	}
	sub2 := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(SubtreeOfAnotherTree(root2, sub2))
}
```
