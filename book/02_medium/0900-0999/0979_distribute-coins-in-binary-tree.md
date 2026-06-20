# 0979 — Distribute Coins In Binary Tree

## Deskripsi

**Soal:** [0979. Distribute Coins In Binary Tree](https://leetcode.com/problems/distribute-coins-in-binary-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h is tree height

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

> **Ide Kunci:** DFS - post-order traversal

## Solusi Go

```go
package main

// LeetCode #979: Distribute Coins in Binary Tree
// https://leetcode.com/problems/distribute-coins-in-binary-tree/
// Difficulty: Medium
//
// Approach: DFS - post-order traversal
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example 1: [3,0,0]
	root1 := &TreeNode{3, &TreeNode{0, nil, nil}, &TreeNode{0, nil, nil}}
	fmt.Println(distributeCoins(root1)) // 2

	// Example 2: [0,3,0]
	root2 := &TreeNode{0, &TreeNode{3, nil, nil}, &TreeNode{0, nil, nil}}
	fmt.Println(distributeCoins(root2)) // 3
}

func distributeCoins(root *TreeNode) int {
	moves := 0
	dfs(root, &moves)
	return moves
}

func dfs(node *TreeNode, moves *int) int {
	if node == nil {
		return 0
	}
	left := dfs(node.Left, moves)
	right := dfs(node.Right, moves)
	*moves += abs(left) + abs(right)
	return node.Val + left + right - 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
