package main

// LeetCode #538: Convert BST to Greater Tree
// https://leetcode.com/problems/convert-bst-to-greater-tree/
// Difficulty: Medium
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}
	root.Right = &TreeNode{Val: 6, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 8}}}
	result := ConvertBST(root)
	printInorder(result)
	fmt.Println()
}

func ConvertBST(root *TreeNode) *TreeNode {
	sum := 0
	var reverseInorder func(node *TreeNode)
	reverseInorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		reverseInorder(node.Right)
		sum += node.Val
		node.Val = sum
		reverseInorder(node.Left)
	}
	reverseInorder(root)
	return root
}

func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInorder(root.Right)
}
