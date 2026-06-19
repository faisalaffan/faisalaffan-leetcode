package main

import "fmt"

// LeetCode #1602: Find Nearest Right Node in Binary Tree
// https://leetcode.com/problems/find-nearest-right-node-in-binary-tree/
// Difficulty: Medium (listed in Hard section)
//
// Given the root of a binary tree and a node u, find the nearest node to the
// right of u on the same level. If u is the rightmost node on its level,
// return nil.

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// findNearestRightNode returns the nearest node to the right of u at the same level.
func findNearestRightNode(root *TreeNode, u *TreeNode) *TreeNode {
	if root == nil || u == nil {
		return nil
	}

	// BFS level-order traversal
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			if node == u {
				// Found u; return the next node at this level, if any
				if i+1 < levelSize {
					return queue[0] // queue[0] is the next node at this level
				}
				return nil
			}

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}

	return nil
}

// BuildBT builds a binary tree from a level-order slice (-1 for nil).
func BuildBT(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == -1 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != -1 {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != -1 {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// findNode finds a node with given value in the tree.
func findNode(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNode(root.Left, val); left != nil {
		return left
	}
	return findNode(root.Right, val)
}

func main() {
	// Example:
	// Tree:
	//       1
	//      / \
	//     2   3
	//    /   / \
	//   4   5   6
	//
	// Find nearest right node of 2: should be 3
	// Find nearest right node of 4: should be nil (rightmost)
	// Find nearest right node of 5: should be 6

	root := BuildBT([]int{1, 2, 3, 4, -1, 5, 6})

	node2 := findNode(root, 2)
	node4 := findNode(root, 4)
	node5 := findNode(root, 5)

	rightOf2 := findNearestRightNode(root, node2)
	rightOf4 := findNearestRightNode(root, node4)
	rightOf5 := findNearestRightNode(root, node5)

	fmt.Println("Nearest right of 2:", rightOf2) // 3
	if rightOf2 != nil {
		fmt.Println("  Val:", rightOf2.Val)
	}

	fmt.Println("Nearest right of 4:", rightOf4) // nil
	if rightOf4 != nil {
		fmt.Println("  Val:", rightOf4.Val)
	}

	fmt.Println("Nearest right of 5:", rightOf5) // 6
	if rightOf5 != nil {
		fmt.Println("  Val:", rightOf5.Val)
	}
}
