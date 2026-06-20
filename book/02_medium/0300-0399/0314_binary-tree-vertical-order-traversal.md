# 0314 — Binary Tree Vertical Order Traversal

## Deskripsi

**Soal:** [0314. Binary Tree Vertical Order Traversal](https://leetcode.com/problems/binary-tree-vertical-order-traversal/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Queue (antrian FIFO)

**Fungsi Solusi:** `func verticalOrder(root *TreeNode) [][]int`

## Solusi Go

```go
package main

// LeetCode #314: Binary Tree Vertical Order Traversal
// https://leetcode.com/problems/binary-tree-vertical-order-traversal/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queueItem struct {
	node *TreeNode
	col  int
}

func verticalOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

  // Membuat map untuk pencarian O(1): key → value
	colMap := make(map[int][]int)
	minCol, maxCol := 0, 0
	queue := []queueItem{{root, 0}}

	for len(queue) > 0 {
		item := queue[0]
		queue = queue[1:]
		node, col := item.node, item.col

		colMap[col] = append(colMap[col], node.Val)

		if col < minCol {
			minCol = col
		}
		if col > maxCol {
			maxCol = col
		}

		if node.Left != nil {
			queue = append(queue, queueItem{node.Left, col - 1})
		}
		if node.Right != nil {
			queue = append(queue, queueItem{node.Right, col + 1})
		}
	}

  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, 0, maxCol-minCol+1)
	for c := minCol; c <= maxCol; c++ {
		result = append(result, colMap[c])
	}
	return result
}

func main() {
	// Test case 1: Example tree [3,9,20,null,null,15,7]
	root1 := &TreeNode{Val: 3}
	root1.Left = &TreeNode{Val: 9}
	root1.Right = &TreeNode{Val: 20}
	root1.Right.Left = &TreeNode{Val: 15}
	root1.Right.Right = &TreeNode{Val: 7}
	fmt.Println("Test 1:", verticalOrder(root1))
	// Expected: [[9],[3,15],[20],[7]]

	// Test case 2: Single node
	root2 := &TreeNode{Val: 1}
	fmt.Println("Test 2:", verticalOrder(root2))
	// Expected: [[1]]

	// Test case 3: Nil
	fmt.Println("Test 3:", verticalOrder(nil))
	// Expected: []
}
```
