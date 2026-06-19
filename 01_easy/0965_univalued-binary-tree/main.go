package main

// LeetCode #965: Univalued Binary Tree
// https://leetcode.com/problems/univalued-binary-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}},
	}
	fmt.Println(isUnivalTree(root)) // true

	root2 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isUnivalTree(root2)) // false
}

// isUnivalTree checks if all nodes in the tree have the same value.
// Time: O(n). Space: O(n).
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsUni(root, root.Val)
}

func dfsUni(node *TreeNode, val int) bool {
	if node == nil {
		return true
	}
	if node.Val != val {
		return false
	}
	return dfsUni(node.Left, val) && dfsUni(node.Right, val)
}
