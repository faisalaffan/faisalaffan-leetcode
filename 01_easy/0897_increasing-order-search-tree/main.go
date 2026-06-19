package main

// LeetCode #897: Increasing Order Search Tree
// https://leetcode.com/problems/increasing-order-search-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 1}},
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   6,
			Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}},
		},
	}
	result := increasingBST(root)
	// Print inorder to verify
	printTree(result) // 1 2 3 4 5 6 7 8 9
}

func printTree(root *TreeNode) {
	for root != nil {
		fmt.Print(root.Val, " ")
		root = root.Right
	}
	fmt.Println()
}

// increasingBST rearranges the BST into increasing order (only right children).
// Time: O(n). Space: O(n).
func increasingBST(root *TreeNode) *TreeNode {
	var newRoot, prev *TreeNode
	inorderTree(root, &newRoot, &prev)
	return newRoot
}

func inorderTree(node *TreeNode, newRoot **TreeNode, prev **TreeNode) {
	if node == nil {
		return
	}
	inorderTree(node.Left, newRoot, prev)
	if *newRoot == nil {
		*newRoot = node
	}
	if *prev != nil {
		(*prev).Right = node
	}
	node.Left = nil
	*prev = node
	inorderTree(node.Right, newRoot, prev)
}
