package main

// LeetCode #783: Minimum Distance Between BST Nodes
// https://leetcode.com/problems/minimum-distance-between-bst-nodes/
// Difficulty: Easy

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 6},
	}
	fmt.Println(minDiffInBST(root)) // 1
}

// minDiffInBST finds the minimum difference between values of any two nodes in a BST.
// Time: O(n). Space: O(n).
func minDiffInBST(root *TreeNode) int {
	minDiff := math.MaxInt32
	var prev *int
	inorder(root, &prev, &minDiff)
	return minDiff
}

func inorder(node *TreeNode, prev **int, minDiff *int) {
	if node == nil {
		return
	}
	inorder(node.Left, prev, minDiff)
	if *prev != nil {
		diff := node.Val - **prev
		if diff < *minDiff {
			*minDiff = diff
		}
	}
	*prev = &node.Val
	inorder(node.Right, prev, minDiff)
}
