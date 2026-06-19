package main

// LeetCode #637: Average of Levels in Binary Tree
// https://leetcode.com/problems/average-of-levels-in-binary-tree/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(n)
func AverageOfLevelsInBinaryTree(root *TreeNode) []float64 {
	var result []float64
	if root == nil {
		return result
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		sum := 0
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, float64(sum)/float64(levelSize))
	}
	return result
}

func main() {
	// Test: [3,9,20,null,null,15,7]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(AverageOfLevelsInBinaryTree(root1))

	// Test: [3,9,20,15,7]
	root2 := &TreeNode{
		Val:   3,
		Left:  &TreeNode{Val: 9, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}},
		Right: &TreeNode{Val: 20},
	}
	fmt.Println(AverageOfLevelsInBinaryTree(root2))
}
