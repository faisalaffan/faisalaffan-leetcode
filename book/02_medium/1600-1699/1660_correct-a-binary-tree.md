# 1660 — Correct A Binary Tree

## Deskripsi

**Soal:** [1660. Correct A Binary Tree](https://leetcode.com/problems/correct-a-binary-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N), Space: O(N)  
**Kompleksitas Ruang:** O(N)

**Algoritma:** BFS (Breadth-First Search / pencarian lebar)

## Solusi Go

```go
package main

// LeetCode #1660: Correct a Binary Tree
// https://leetcode.com/problems/correct-a-binary-tree/
// Difficulty: Medium [Paid]

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree where a node has an invalid right pointer pointing to a node at the same level or below.
	// Example: 1 -> 2, 3; 2 -> 4 (right points to 3)
	//          1
	//         / \
	//        2   3
	//         \
	//          4 (right -> 3, invalid)
	n3 := &TreeNode{Val: 3}
	n4 := &TreeNode{Val: 4, Right: n3} // invalid pointer
	n2 := &TreeNode{Val: 2, Right: n4}
	root := &TreeNode{Val: 1, Left: n2, Right: n3}

	corrected := CorrectBinaryTree(root)
	fmt.Println("Root val:", corrected.Val)
	fmt.Println("Left:", corrected.Left.Val)
	fmt.Println("Right:", corrected.Right.Val)
	// Left.Right should be 4 but 4's Right should be nil (corrected)
	if corrected.Left.Right != nil {
		fmt.Println("Left.Right.Right (should be nil):", corrected.Left.Right.Right)
	}
}

func CorrectBinaryTree(root *TreeNode) *TreeNode {
	// Time: O(N), Space: O(N)
	// BFS to find the invalid node, remove it
  // Membuat map untuk pencarian O(1): key → value
	visited := make(map[*TreeNode]bool)
  // Membuat map untuk pencarian O(1): key → value
	parent := make(map[*TreeNode]*TreeNode)
	queue := []*TreeNode{root}
	visited[root] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left != nil {
			if visited[node.Left] {
				// node.Left points to an already visited node (invalid)
				removeNode(parent, node, 'L')
				return root
			}
			visited[node.Left] = true
			parent[node.Left] = node
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			if visited[node.Right] {
				// node.Right points to an already visited node (invalid)
				removeNode(parent, node, 'R')
				return root
			}
			visited[node.Right] = true
			parent[node.Right] = node
			queue = append(queue, node.Right)
		}
	}

	return root
}

func removeNode(parent map[*TreeNode]*TreeNode, node *TreeNode, child byte) {
	if p, ok := parent[node]; ok {
		if p.Left == node {
			p.Left = nil
		} else {
			p.Right = nil
		}
	} else {
		// node is root, can't easily remove - but this case doesn't happen
		// since root can't have invalid pointers
	}
}
```
