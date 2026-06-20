# 1516 — Move Sub Tree Of N Ary Tree

## Deskripsi

**Soal:** [1516. Move Sub Tree Of N Ary Tree](https://leetcode.com/problems/move-sub-tree-of-n-ary-tree/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func moveSubTree(root, p, q *Node) *Node`

## Solusi Go

```go
package main

import "fmt"

// LeetCode #1516: Move Sub-Tree of N-Ary Tree
// https://leetcode.com/problems/move-sub-tree-of-n-ary-tree/
// Difficulty: Hard [Paid]
//
// Given the root of an N-ary tree, a node p, and a node q,
// make q the new parent of the subtree rooted at p.
// p is not an ancestor of q (guaranteed by the problem).

// Node represents an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// moveSubTree makes q the new parent of the subtree rooted at p.
// It returns the new root (which may change if the root itself is moved).
func moveSubTree(root, p, q *Node) *Node {
	if root == nil || p == nil || q == nil {
		return root
	}
	if p == root {
		// p is the root – we cannot move root under q unless we detach first.
		// Problem usually guarantees p is not root, but handle gracefully.
		removeChild(root, p)
		q.Children = append(q.Children, p)
		return root
	}

	// Find parent of p
	parentP, _ := findParent(root, p, nil)
	if parentP == nil {
		return root // p not found
	}

	// Remove p from its current parent
	removeChild(parentP, p)

	// Make q the new parent
	q.Children = append(q.Children, p)

	return root
}

// findParent traverses the tree to find parent of target.
func findParent(node, target, parent *Node) (*Node, *Node) {
	if node == nil {
		return nil, nil
	}
	if node == target {
		return parent, node
	}
	for _, child := range node.Children {
		if p, found := findParent(child, target, node); found != nil {
			return p, found
		}
	}
	return nil, nil
}

// removeChild removes child from parent's Children slice.
func removeChild(parent, child *Node) {
	for i, c := range parent.Children {
		if c == child {
			parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
			return
		}
	}
}

// Helper to build a tree from adjacency list.
// nodes[0] is the root.
func buildTree(adj [][]int) *Node {
	if len(adj) == 0 {
		return nil
	}
  // Membuat slice untuk menyimpan hasil
	nodes := make([]*Node, len(adj))
  // Iterasi seluruh elemen
	for i := range adj {
		nodes[i] = &Node{Val: i}
	}
	for i, children := range adj {
		for _, c := range children {
			nodes[i].Children = append(nodes[i].Children, nodes[c])
		}
	}
	return nodes[0]
}

// Helper to collect tree values via preorder traversal for verification.
func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	res := []int{root.Val}
	for _, c := range root.Children {
		res = append(res, preorder(c)...)
	}
	return res
}

func main() {
	// Test case:
	// Tree: 0 -> [1, 2], 1 -> [3, 4], 2 -> [5]
	// Root: 0
	// Move subtree of node 1 under node 2 (q=2, p=1)
	// Before: 0 has children [1, 2]; 1 has children [3, 4]; 2 has child [5]
	// After:  0 has child [2]; 2 has children [5, 1]; 1 has children [3, 4]

	adj := [][]int{
		{1, 2}, // 0's children
		{3, 4}, // 1's children
		{5},    // 2's children
		{},     // 3's children
		{},     // 4's children
		{},     // 5's children
	}
	root := buildTree(adj)

	// Find nodes p=1 and q=2
	var p, q *Node
	var findNodes func(*Node)
	findNodes = func(n *Node) {
		if n == nil {
			return
		}
		if n.Val == 1 {
			p = n
		}
		if n.Val == 2 {
			q = n
		}
		for _, c := range n.Children {
			findNodes(c)
		}
	}
	findNodes(root)

	fmt.Println("Before move:")
	fmt.Println("Root preorder:", preorder(root))

	newRoot := moveSubTree(root, p, q)

	fmt.Println("After moving subtree at 1 under 2:")
	fmt.Println("Root preorder:", preorder(newRoot))
	fmt.Println("Expected: [0 2 5 1 3 4]")
}
```
