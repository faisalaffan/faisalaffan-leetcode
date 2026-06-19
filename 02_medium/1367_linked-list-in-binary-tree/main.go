package main

// LeetCode #1367: Linked List in Binary Tree
// https://leetcode.com/problems/linked-list-in-binary-tree/
// Difficulty: Medium

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	head := &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 8}}}
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 4,
			Right: &TreeNode{
				Val: 2,
				Left: &TreeNode{Val: 1},
			},
		},
		Right: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 2,
				Left: &TreeNode{Val: 6},
				Right: &TreeNode{
					Val: 8,
					Left: &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
			},
		},
	}
	fmt.Println(isSubPath(head, root)) // true

	// Test case 2
	head2 := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 2, Next: &ListNode{Val: 6}}}}
	fmt.Println(isSubPath(head2, root)) // true

	// Test case 3
	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	head3 := &ListNode{Val: 4}
	fmt.Println(isSubPath(head3, root3)) // false
}

// Time: O(N * L) where N = tree nodes, L = list length
// Space: O(N) for recursion stack
func isSubPath(head *ListNode, root *TreeNode) bool {
	if root == nil {
		return false
	}
	if dfs(head, root) {
		return true
	}
	return isSubPath(head, root.Left) || isSubPath(head, root.Right)
}

func dfs(head *ListNode, node *TreeNode) bool {
	if head == nil {
		return true
	}
	if node == nil {
		return false
	}
	if head.Val != node.Val {
		return false
	}
	return dfs(head.Next, node.Left) || dfs(head.Next, node.Right)
}
