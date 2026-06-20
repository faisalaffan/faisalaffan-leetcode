# 0617 — Merge Two Binary Trees

## Deskripsi

**Soal:** [0617. Merge Two Binary Trees](https://leetcode.com/problems/merge-two-binary-trees/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n+m), Space: O(h)  
**Kompleksitas Ruang:** O(h)

**Algoritma:** —

**Fungsi Solusi:** `func MergeTwoBinaryTrees(root1, root2 *TreeNode) *TreeNode`

## Solusi Go

```go
package main

// LeetCode #617: Merge Two Binary Trees
// https://leetcode.com/problems/merge-two-binary-trees/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n+m), Space: O(h)
func MergeTwoBinaryTrees(root1, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	return &TreeNode{
		Val:   root1.Val + root2.Val,
		Left:  MergeTwoBinaryTrees(root1.Left, root2.Left),
		Right: MergeTwoBinaryTrees(root1.Right, root2.Right),
	}
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("null ")
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

func main() {
	// Test: root1=[1,3,2,5], root2=[2,1,3,null,4,null,7]
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 2},
	}
	root2 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:   1,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 7},
		},
	}
	printTree(MergeTwoBinaryTrees(root1, root2))
	fmt.Println()

	// Test: root1=[1], root2=[1,2]
	r1 := &TreeNode{Val: 1}
	r2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	printTree(MergeTwoBinaryTrees(r1, r2))
	fmt.Println()
}
```
