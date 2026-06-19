package main

// LeetCode #2415: Reverse Odd Levels of Binary Tree
// https://leetcode.com/problems/reverse-odd-levels-of-binary-tree/
// Difficulty: Medium
// Time: O(n) | Space: O(n)
// BFS level-order, reverse values at odd levels.

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [2,3,5,8,13,21,34]
	root := &TreeNode{2,
		&TreeNode{3,
			&TreeNode{8, nil, nil},
			&TreeNode{13, nil, nil},
		},
		&TreeNode{5,
			&TreeNode{21, nil, nil},
			&TreeNode{34, nil, nil},
		},
	}
	reversed := reverseOddLevels(root)
	// After reverse: level 1: 3<->5, becomes [2,5,3,...]
	fmt.Println(reversed.Left.Val)  // 5
	fmt.Println(reversed.Right.Val) // 3

	root2 := &TreeNode{7,
		&TreeNode{13, nil, nil},
		&TreeNode{11, nil, nil},
	}
	rev2 := reverseOddLevels(root2)
	fmt.Println(rev2.Left.Val)  // 11
	fmt.Println(rev2.Right.Val) // 13
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	q := []*TreeNode{root}
	level := 0
	for len(q) > 0 {
		if level%2 == 1 {
			// reverse values at this level
			for i, j := 0, len(q)-1; i < j; i, j = i+1, j-1 {
				q[i].Val, q[j].Val = q[j].Val, q[i].Val
			}
		}
		next := make([]*TreeNode, 0)
		for _, node := range q {
			if node.Left != nil {
				next = append(next, node.Left)
				next = append(next, node.Right)
			}
		}
		q = next
		level++
	}
	return root
}
