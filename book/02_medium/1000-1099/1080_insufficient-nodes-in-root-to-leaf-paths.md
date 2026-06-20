# 1080 — Insufficient Nodes In Root To Leaf Paths

## Deskripsi

**Soal:** [1080. Insufficient Nodes In Root To Leaf Paths](https://leetcode.com/problems/insufficient-nodes-in-root-to-leaf-paths/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h is tree height

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

> **Ide Kunci:** DFS post-order. Delete node if sum from root to leaf < limit.

## Solusi Go

```go
package main

// LeetCode #1080: Insufficient Nodes in Root to Leaf Paths
// https://leetcode.com/problems/insufficient-nodes-in-root-to-leaf-paths/
// Difficulty: Medium
//
// Approach: DFS post-order. Delete node if sum from root to leaf < limit.
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4, Left: &TreeNode{Val: 8, Left: nil, Right: nil}, Right: nil},
			Right: &TreeNode{Val: 5, Left: nil, Right: nil},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6, Left: nil, Right: nil},
			Right: &TreeNode{Val: 7, Left: nil, Right: nil},
		},
	}
	result := sufficientSubset(root, 10)
	printTree(result)
	fmt.Println()
}

func sufficientSubset(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		if root.Val < limit {
			return nil
		}
		return root
	}

	root.Left = sufficientSubset(root.Left, limit-root.Val)
	root.Right = sufficientSubset(root.Right, limit-root.Val)

	if root.Left == nil && root.Right == nil {
		return nil
	}
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("null")
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			fmt.Print("null ")
			continue
		}
		fmt.Printf("%d ", node.Val)
		queue = append(queue, node.Left, node.Right)
	}
}
```
