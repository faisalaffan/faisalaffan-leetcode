package main

// LeetCode #1490: Clone N-ary Tree
// https://leetcode.com/problems/clone-n-ary-tree/
// Difficulty: Medium [Paid]

import "fmt"

// Node is an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

func main() {
	// Build tree: root = 1 -> [3, 2, 4]; 3 -> [5, 6]
	root := &Node{Val: 1}
	child3 := &Node{Val: 3}
	child2 := &Node{Val: 2}
	child4 := &Node{Val: 4}
	root.Children = []*Node{child3, child2, child4}
	child3.Children = []*Node{{Val: 5}, {Val: 6}}

	cloned := CloneTree(root)
	fmt.Println("Root cloned:", cloned != nil && cloned != root)
	fmt.Println("Root val:", cloned.Val)
	fmt.Println("Children count:", len(cloned.Children))
	fmt.Println("Deep cloned:", cloned.Children[0].Children[0].Val == 5 && cloned.Children[0] != child3.Children[0])

	// Test nil
	fmt.Println(CloneTree(nil))
}

func CloneTree(root *Node) *Node {
	// Time: O(N), Space: O(N) (recursion stack)
	if root == nil {
		return nil
	}

	clone := &Node{Val: root.Val}
	clone.Children = make([]*Node, len(root.Children))
	for i, child := range root.Children {
		clone.Children[i] = CloneTree(child)
	}
	return clone
}
