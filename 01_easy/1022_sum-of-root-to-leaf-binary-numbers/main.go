package main

// LeetCode #1022: Sum of Root To Leaf Binary Numbers
// https://leetcode.com/problems/sum-of-root-to-leaf-binary-numbers/
// Difficulty: Easy
// Time: O(n) | Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//     1
	//    / \
	//   0   1
	//  / \ / \
	// 0  1 0  1
	root := &TreeNode{1,
		&TreeNode{0, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}},
		&TreeNode{1, &TreeNode{0, nil, nil}, &TreeNode{1, nil, nil}},
	}
	fmt.Println(sumRootToLeaf(root)) // 22

	// Single node
	fmt.Println(sumRootToLeaf(&TreeNode{1, nil, nil})) // 1
}

// LeetCode submission: sumRootToLeaf
func sumRootToLeaf(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode, val int) int {
	if node == nil {
		return 0
	}
	val = val*2 + node.Val
	if node.Left == nil && node.Right == nil {
		return val
	}
	return dfs(node.Left, val) + dfs(node.Right, val)
}
