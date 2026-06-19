package main

// LeetCode #951: Flip Equivalent Binary Trees
// https://leetcode.com/problems/flip-equivalent-binary-trees/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(n)
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}

func main() {
	// [1,2,3,4,5,6,null,null,null,7,8]
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{
				Val: 5,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 8},
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
		},
	}
	// [1,3,2,null,6,4,5,null,null,null,null,8,7]
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Right: &TreeNode{Val: 8},
			},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 7},
			},
		},
	}
	fmt.Println(flipEquiv(root1, root2))

	// Trivial: both nil
	fmt.Println(flipEquiv(nil, nil))

	// Trivial: one nil
	fmt.Println(flipEquiv(&TreeNode{Val: 1}, nil))
}
