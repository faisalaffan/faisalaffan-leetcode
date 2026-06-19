package main

// LeetCode #589: N-ary Tree Preorder Traversal
// https://leetcode.com/problems/n-ary-tree-preorder-traversal/
// Difficulty: Easy

import "fmt"

// Node represents an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// Time: O(n), Space: O(n)
func NAryTreePreorderTraversal(root *Node) []int {
	var result []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		for _, child := range node.Children {
			dfs(child)
		}
	}
	dfs(root)
	return result
}

func main() {
	// Test: [1,null,3,2,4,null,5,6]
	root1 := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 3, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 2},
			{Val: 4},
		},
	}
	fmt.Println(NAryTreePreorderTraversal(root1))

	// Test: single node
	root2 := &Node{Val: 1}
	fmt.Println(NAryTreePreorderTraversal(root2))
}
