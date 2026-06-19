package main

// LeetCode #105: Construct Binary Tree from Preorder and Inorder Traversal
// https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	inorderMap := make(map[int]int)
	for i, v := range inorder {
		inorderMap[v] = i
	}

	var build func(preStart, preEnd, inStart, inEnd int) *TreeNode
	build = func(preStart, preEnd, inStart, inEnd int) *TreeNode {
		if preStart > preEnd || inStart > inEnd {
			return nil
		}

		rootVal := preorder[preStart]
		root := &TreeNode{Val: rootVal}
		inIdx := inorderMap[rootVal]
		leftSize := inIdx - inStart

		root.Left = build(preStart+1, preStart+leftSize, inStart, inIdx-1)
		root.Right = build(preStart+leftSize+1, preEnd, inIdx+1, inEnd)
		return root
	}

	return build(0, len(preorder)-1, 0, len(inorder)-1)
}

func printInorder(root *TreeNode) {
	if root == nil {
		return
	}
	printInorder(root.Left)
	fmt.Printf("%d ", root.Val)
	printInorder(root.Right)
}

func main() {
	// Test case 1
	root := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	printInorder(root) // 9 3 15 20 7
	fmt.Println()

	// Test case 2
	root = buildTree([]int{-1}, []int{-1})
	printInorder(root) // -1
	fmt.Println()
}

// Time: O(n) | Space: O(n)
