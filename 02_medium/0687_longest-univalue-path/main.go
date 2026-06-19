package main

// LeetCode #687: Longest Univalue Path
// https://leetcode.com/problems/longest-univalue-path/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 5}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 1}
	root.Right.Right = &TreeNode{Val: 5}

	fmt.Println(longestUnivaluePath(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestUnivaluePath(root *TreeNode) int {
	maxLen := 0

	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		leftLen := dfs(node.Left)
		rightLen := dfs(node.Right)

		leftArrow, rightArrow := 0, 0
		if node.Left != nil && node.Left.Val == node.Val {
			leftArrow = leftLen + 1
		}
		if node.Right != nil && node.Right.Val == node.Val {
			rightArrow = rightLen + 1
		}

		maxLen = max(maxLen, leftArrow+rightArrow)

		return max(leftArrow, rightArrow)
	}

	dfs(root)
	return maxLen
}
