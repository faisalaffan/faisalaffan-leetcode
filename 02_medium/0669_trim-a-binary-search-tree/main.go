package main

// LeetCode #669: Trim a Binary Search Tree
// https://leetcode.com/problems/trim-a-binary-search-tree/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 0}
	root.Right = &TreeNode{Val: 2}

	result := trimBST(root, 1, 2)
	fmt.Println(result.Val)
	if result.Left != nil {
		fmt.Println(result.Left.Val)
	}
	if result.Right != nil {
		fmt.Println(result.Right.Val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < low {
		return trimBST(root.Right, low, high)
	}
	if root.Val > high {
		return trimBST(root.Left, low, high)
	}

	root.Left = trimBST(root.Left, low, high)
	root.Right = trimBST(root.Right, low, high)
	return root
}
