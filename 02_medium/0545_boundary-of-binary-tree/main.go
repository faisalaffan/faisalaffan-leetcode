package main

// LeetCode #545: Boundary of Binary Tree
// https://leetcode.com/problems/boundary-of-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}
	fmt.Println(BoundaryOfBinaryTree(root))
}

func BoundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	result := []int{root.Val}
	addLeftBoundary(root.Left, &result)
	addLeaves(root.Left, &result)
	addLeaves(root.Right, &result)
	addRightBoundary(root.Right, &result)
	return result
}

func addLeftBoundary(node *TreeNode, result *[]int) {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return
	}
	*result = append(*result, node.Val)
	if node.Left != nil {
		addLeftBoundary(node.Left, result)
	} else {
		addLeftBoundary(node.Right, result)
	}
}

func addRightBoundary(node *TreeNode, result *[]int) {
	if node == nil || (node.Left == nil && node.Right == nil) {
		return
	}
	if node.Right != nil {
		addRightBoundary(node.Right, result)
	} else {
		addRightBoundary(node.Left, result)
	}
	*result = append(*result, node.Val) // post-order for reverse
}

func addLeaves(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*result = append(*result, node.Val)
		return
	}
	addLeaves(node.Left, result)
	addLeaves(node.Right, result)
}
