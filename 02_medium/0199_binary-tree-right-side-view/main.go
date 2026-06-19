package main

// LeetCode #199: Binary Tree Right Side View
// https://leetcode.com/problems/binary-tree-right-side-view/
// Difficulty: Medium
// Time: O(n), Space: O(h) for queue

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if i == levelSize-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return result
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, &TreeNode{4, nil, nil}}}
	fmt.Println(rightSideView(root))

	root2 := &TreeNode{1, nil, &TreeNode{3, nil, nil}}
	fmt.Println(rightSideView(root2))

	fmt.Println(rightSideView(nil))
}
