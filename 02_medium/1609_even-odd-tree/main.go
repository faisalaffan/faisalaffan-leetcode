package main

// LeetCode #1609: Even Odd Tree
// https://leetcode.com/problems/even-odd-tree/
// Difficulty: Medium

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [1,10,4,3,null,7,9,12,8,6,null,null,2]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 10, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 12}, Right: &TreeNode{Val: 8}}}
	root.Right = &TreeNode{Val: 4, Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 9, Right: &TreeNode{Val: 2}}}
	fmt.Println(IsEvenOddTree(root))

	// Simple tree
	root2 := &TreeNode{Val: 5, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 2}}
	fmt.Println(IsEvenOddTree(root2))

	// Tree: [5,9,1,3,5,7]
	root3 := &TreeNode{Val: 5}
	root3.Left = &TreeNode{Val: 9, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}
	root3.Right = &TreeNode{Val: 1, Left: &TreeNode{Val: 7}}
	fmt.Println(IsEvenOddTree(root3))
}

func IsEvenOddTree(root *TreeNode) bool {
	// Time: O(N), Space: O(N)
	if root == nil {
		return true
	}

	queue := []*TreeNode{root}
	level := 0

	for len(queue) > 0 {
		levelSize := len(queue)
		var prevVal int

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if level%2 == 0 {
				// Even level: odd values, strictly increasing
				if node.Val%2 == 0 {
					return false
				}
				if i > 0 && node.Val <= prevVal {
					return false
				}
			} else {
				// Odd level: even values, strictly decreasing
				if node.Val%2 != 0 {
					return false
				}
				if i > 0 && node.Val >= prevVal {
					return false
				}
			}

			prevVal = node.Val

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		level++
	}

	return true
}
