package main

// LeetCode #230: Kth Smallest Element in a BST
// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
// Difficulty: Medium
// Time: O(h+k), Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	node := root

	for {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--

		if k == 0 {
			return node.Val
		}

		node = node.Right
	}
}

func main() {
	root := &TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}
	fmt.Println(kthSmallest(root, 1))

	root2 := &TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, nil}}
	fmt.Println(kthSmallest(root2, 3))
}
