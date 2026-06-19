package main

// LeetCode #98: Validate Binary Search Tree
// https://leetcode.com/problems/validate-binary-search-tree/
// Difficulty: Medium

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	return validate(root, math.MinInt64, math.MaxInt64)
}

func validate(node *TreeNode, min, max int64) bool {
	if node == nil {
		return true
	}
	if int64(node.Val) <= min || int64(node.Val) >= max {
		return false
	}
	return validate(node.Left, min, int64(node.Val)) && validate(node.Right, int64(node.Val), max)
}

func main() {
	// Test case 1: [2,1,3] -> true
	root := &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}
	fmt.Println(isValidBST(root)) // true

	// Test case 2: [5,1,4,null,null,3,6] -> false
	root = &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}}
	fmt.Println(isValidBST(root)) // false

	// Test case 3: [2,2,2] -> false
	root = &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 2}}
	fmt.Println(isValidBST(root)) // false
}

// Time: O(n) | Space: O(n)
