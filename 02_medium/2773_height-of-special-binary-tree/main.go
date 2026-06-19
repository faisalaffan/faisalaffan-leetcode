package main

// LeetCode #2773: Height of Special Binary Tree
// https://leetcode.com/problems/height-of-special-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(h)

import "fmt"

type SpecialTreeNode struct {
	Val       int
	Left      *SpecialTreeNode
	Right     *SpecialTreeNode
	IsSpecial bool
}

func HeightOfSpecialBinaryTree(root *SpecialTreeNode) int {
	var dfs func(*SpecialTreeNode) int
	dfs = func(node *SpecialTreeNode) int {
		if node == nil || node.IsSpecial {
			return 0
		}
		leftH := dfs(node.Left)
		rightH := dfs(node.Right)
		if leftH > rightH {
			return leftH + 1
		}
		return rightH + 1
	}
	return dfs(root)
}

func main() {
	// Tree: 1(not special) -> 2(special), 3(not special)
	root := &SpecialTreeNode{
		Val: 1,
		Left: &SpecialTreeNode{
			Val:       2,
			IsSpecial: true,
		},
		Right: &SpecialTreeNode{
			Val: 3,
			Left: &SpecialTreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(HeightOfSpecialBinaryTree(root))
	fmt.Println(HeightOfSpecialBinaryTree(nil))
}
