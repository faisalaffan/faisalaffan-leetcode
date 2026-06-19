package main

// LeetCode #124: Binary Tree Maximum Path Sum
// https://leetcode.com/problems/binary-tree-maximum-path-sum/
// Difficulty: Hard

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := root.Val

	var dfs func(*TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftGain := max(0, dfs(node.Left))
		rightGain := max(0, dfs(node.Right))

		currentPathSum := node.Val + leftGain + rightGain
		if currentPathSum > maxSum {
			maxSum = currentPathSum
		}

		return node.Val + max(leftGain, rightGain)
	}

	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example: [-10,9,20,null,null,15,7] -> 42
	root := &TreeNode{Val: -10}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}

	result := maxPathSum(root)
	expected := 42

	fmt.Printf("maxPathSum([-10,9,20,null,null,15,7]) = %d\n", result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
