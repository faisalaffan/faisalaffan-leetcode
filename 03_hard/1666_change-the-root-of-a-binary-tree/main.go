package main

// LeetCode #1666: Change the Root of a Binary Tree
// https://leetcode.com/problems/change-the-root-of-a-binary-tree/
// Difficulty: Hard [Premium]

import "fmt"

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func main() {
	// Build tree: root=3
	//     3
	//    / \
	//   5   1
	//  / \  / \
	// 6   2 0  8
	//    / \
	//   7   4
	root := &Node{Val: 3}
	n5 := &Node{Val: 5}
	n1 := &Node{Val: 1}
	n6 := &Node{Val: 6}
	n2 := &Node{Val: 2}
	n7 := &Node{Val: 7}
	n4 := &Node{Val: 4}
	n0 := &Node{Val: 0}
	n8 := &Node{Val: 8}

	root.Left, root.Right = n5, n1
	n5.Parent, n1.Parent = root, root
	n5.Left, n5.Right = n6, n2
	n6.Parent, n2.Parent = n5, n5
	n2.Left, n2.Right = n7, n4
	n7.Parent, n4.Parent = n2, n2
	n1.Left, n1.Right = n0, n8
	n0.Parent, n8.Parent = n1, n1

	// Change root to node 2
	newRoot := flipBinaryTree(root, n2)
	fmt.Printf("Test 1 - New root value: %d (Expected: 2)\n", newRoot.Val)
	fmt.Printf("Right child (was old parent 5): %d (Expected: 5)\n", newRoot.Right.Val)
	fmt.Printf("Left child of 5 (was old parent 3): %d (Expected: 3)\n", n5.Left.Val)
	fmt.Printf("Parent of new root is nil: %v (Expected: true)\n", newRoot.Parent == nil)
	fmt.Printf("Parent of old root is child: %d (Expected: 5)\n", root.Parent.Val)
}

func flipBinaryTree(root *Node, leaf *Node) *Node {
	// Walk from leaf up to root, reversing parent-child relationships
	return flip(leaf, nil)
}

func flip(node, newParent *Node) *Node {
	oldParent := node.Parent
	node.Parent = newParent

	// Disconnect the child pointer that now points to the new parent
	if node.Left == newParent {
		node.Left = nil
	} else if node.Right == newParent {
		node.Right = nil
	}

	if oldParent != nil {
		// Disconnect old parent's pointer to this node
		if oldParent.Left == node {
			oldParent.Left = nil
		} else if oldParent.Right == node {
			oldParent.Right = nil
		}
		flip(oldParent, node)
		// Attach old parent as a child of current node
		if node.Left == nil {
			node.Left = oldParent
		} else {
			node.Right = oldParent
		}
	}

	return node
}
