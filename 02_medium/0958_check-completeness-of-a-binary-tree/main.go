package main

// LeetCode #958: Check Completeness of a Binary Tree
// https://leetcode.com/problems/check-completeness-of-a-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(n)
func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	seenNull := false
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			seenNull = true
		} else {
			if seenNull {
				return false
			}
			q = append(q, node.Left, node.Right)
		}
	}
	return true
}

func main() {
	// Complete tree: [1,2,3,4,5,6]
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
		},
	}
	fmt.Println(isCompleteTree(root))

	// Not complete: [1,2,3,4,5,null,7]
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(isCompleteTree(root2))
}
