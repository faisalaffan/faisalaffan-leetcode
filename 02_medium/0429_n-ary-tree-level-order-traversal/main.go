package main

// LeetCode #429: N-ary Tree Level Order Traversal
// https://leetcode.com/problems/n-ary-tree-level-order-traversal/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, 0, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			queue = append(queue, node.Children...)
		}
		result = append(result, level)
	}
	return result
}

func main() {
	// Test case 1: [1,null,3,2,4,null,5,6]
	root1 := &Node{Val: 1}
	root1.Children = []*Node{
		{Val: 3, Children: []*Node{{Val: 5}, {Val: 6}}},
		{Val: 2},
		{Val: 4},
	}
	fmt.Println("Test 1:", levelOrder(root1))
	// Expected: [[1],[3,2,4],[5,6]]

	// Test case 2: Single node
	root2 := &Node{Val: 1}
	fmt.Println("Test 2:", levelOrder(root2))
	// Expected: [[1]]

	// Test case 3: Nil
	fmt.Println("Test 3:", levelOrder(nil))
	// Expected: []
}
