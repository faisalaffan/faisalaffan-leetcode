# 0987 — Vertical Order Traversal Of A Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func verticalTraversal(root *TreeNode) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS, Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #987: Vertical Order Traversal of a Binary Tree
// https://leetcode.com/problems/vertical-order-traversal-of-a-binary-tree/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

// TreeNode definition
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeInfo struct {
	row int
	col int
	val int
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var nodes []nodeInfo
	var dfs func(node *TreeNode, row, col int)
	dfs = func(node *TreeNode, row, col int) {
		if node == nil {
			return
		}
		nodes = append(nodes, nodeInfo{row, col, node.Val})
		dfs(node.Left, row+1, col-1)
		dfs(node.Right, row+1, col+1)
	}
	dfs(root, 0, 0)

	// Sort: by column ascending, then row ascending, then value ascending
  // Custom sort
	sort.Slice(nodes, func(i, j int) bool {
		if nodes[i].col != nodes[j].col {
			return nodes[i].col < nodes[j].col
		}
		if nodes[i].row != nodes[j].row {
			return nodes[i].row < nodes[j].row
		}
		return nodes[i].val < nodes[j].val
	})

	var result [][]int
	currentCol := nodes[0].col
	var currentGroup []int

	for _, n := range nodes {
		if n.col != currentCol {
			result = append(result, currentGroup)
			currentGroup = nil
			currentCol = n.col
		}
		currentGroup = append(currentGroup, n.val)
	}
	result = append(result, currentGroup)

	return result
}

func main() {
	// Example 1: [3,9,20,null,null,15,7]
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}
	fmt.Println("Example 1:")
	fmt.Println(verticalTraversal(root1))
	// Expected: [[9],[3,15],[20],[7]]

	// Example 2: [1,2,3,4,5,6,7]
	root2 := &TreeNode{Val: 1}
	root2.Left = &TreeNode{Val: 2}
	root2.Right = &TreeNode{Val: 3}
	root2.Left.Left = &TreeNode{Val: 4}
	root2.Left.Right = &TreeNode{Val: 5}
	root2.Right.Left = &TreeNode{Val: 6}
	root2.Right.Right = &TreeNode{Val: 7}
	fmt.Println("Example 2:")
	fmt.Println(verticalTraversal(root2))
	// Expected: [[4],[2],[1,5,6],[3],[7]]

	// Example 3: [1,2,3,4,6,5,7]
	root3 := &TreeNode{Val: 1}
	root3.Left = &TreeNode{Val: 2}
	root3.Right = &TreeNode{Val: 3}
	root3.Left.Left = &TreeNode{Val: 4}
	root3.Left.Right = &TreeNode{Val: 6}
	root3.Right.Left = &TreeNode{Val: 5}
	root3.Right.Right = &TreeNode{Val: 7}
	fmt.Println("Example 3:")
	fmt.Println(verticalTraversal(root3))
	// Expected: [[4],[2],[1,5,6],[3],[7]]
}
```
