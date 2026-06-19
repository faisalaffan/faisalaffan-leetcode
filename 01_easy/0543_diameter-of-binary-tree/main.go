package main

// LeetCode #543: Diameter of Binary Tree
// https://leetcode.com/problems/diameter-of-binary-tree/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(h)
func DiameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		if left+right > maxDiameter {
			maxDiameter = left + right
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}
	dfs(root)
	return maxDiameter
}

func main() {
	// Test: [1,2,3,4,5]
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(DiameterOfBinaryTree(root1))

	// Test: [1,2]
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	fmt.Println(DiameterOfBinaryTree(root2))
}
