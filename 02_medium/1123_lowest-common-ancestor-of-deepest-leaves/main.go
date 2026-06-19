package main

// LeetCode #1123: Lowest Common Ancestor of Deepest Leaves
// https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/
// Difficulty: Medium
//
// Approach: DFS. Return depth and LCA candidate for each subtree.
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7, Left: nil, Right: nil}, Right: &TreeNode{Val: 4, Left: nil, Right: nil}},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
			Right: &TreeNode{Val: 8, Left: nil, Right: nil},
		},
	}
	result := lcaDeepestLeaves(root)
	fmt.Println(result.Val) // 2
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, lca := dfs(root)
	return lca
}

func dfs(node *TreeNode) (int, *TreeNode) {
	if node == nil {
		return 0, nil
	}

	leftDepth, leftLCA := dfs(node.Left)
	rightDepth, rightLCA := dfs(node.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1, leftLCA
	} else if rightDepth > leftDepth {
		return rightDepth + 1, rightLCA
	}
	return leftDepth + 1, node
}
