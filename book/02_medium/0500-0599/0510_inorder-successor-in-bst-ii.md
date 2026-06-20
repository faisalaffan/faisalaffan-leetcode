# 0510 — Inorder Successor In Bst Ii

## Deskripsi

**Soal:** [0510. Inorder Successor In Bst Ii](https://leetcode.com/problems/inorder-successor-in-bst-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(h) where h is height of tree  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #510: Inorder Successor in BST II
// https://leetcode.com/problems/inorder-successor-in-bst-ii/
// Difficulty: Medium [Paid]
// Time: O(h) where h is height of tree
// Space: O(1)

import "fmt"

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func main() {
	// Build tree: [2,1,3]
	root := &Node{Val: 2}
	root.Left = &Node{Val: 1, Parent: root}
	root.Right = &Node{Val: 3, Parent: root}
	fmt.Println(InorderSuccessorInBstIi(root.Left).Val) // node 1 -> successor 2
	fmt.Println(InorderSuccessorInBstIi(root).Val)       // node 2 -> successor 3

	// For node 3, successor should be nil
	successor := InorderSuccessorInBstIi(root.Right)
	if successor == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(successor.Val)
	}
}

func InorderSuccessorInBstIi(node *Node) *Node {
	if node == nil {
		return nil
	}

	// If right child exists, find leftmost in right subtree
	if node.Right != nil {
		cur := node.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		return cur
	}

	// Otherwise, go up until we find a node that is a left child
	cur := node
	for cur.Parent != nil && cur.Parent.Right == cur {
		cur = cur.Parent
	}

	return cur.Parent
}
```
