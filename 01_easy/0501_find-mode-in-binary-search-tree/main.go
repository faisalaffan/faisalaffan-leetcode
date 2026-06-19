package main

// LeetCode #501: Find Mode in Binary Search Tree
// https://leetcode.com/problems/find-mode-in-binary-search-tree/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(n)
func FindModeInBinarySearchTree(root *TreeNode) []int {
	var result []int
	maxCount, currentCount := 0, 0
	var prev *int

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil && *prev == node.Val {
			currentCount++
		} else {
			currentCount = 1
		}
		prev = &node.Val
		if currentCount > maxCount {
			maxCount = currentCount
			result = []int{node.Val}
		} else if currentCount == maxCount {
			result = append(result, node.Val)
		}
		inorder(node.Right)
	}
	inorder(root)
	return result
}

func main() {
	// Test: [1,null,2,2]
	root1 := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 2},
		},
	}
	fmt.Println(FindModeInBinarySearchTree(root1))

	// Test: [0]
	root2 := &TreeNode{Val: 0}
	fmt.Println(FindModeInBinarySearchTree(root2))
}
