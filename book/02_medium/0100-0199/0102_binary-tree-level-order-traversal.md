# 0102 — Binary Tree Level Order Traversal

## Deskripsi

**Soal:** [0102. Binary Tree Level Order Traversal](https://leetcode.com/problems/binary-tree-level-order-traversal/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Queue (antrian FIFO)

**Fungsi Solusi:** `func levelOrder(root *TreeNode) [][]int`

## Solusi Go

```go
package main

// LeetCode #102: Binary Tree Level Order Traversal
// https://leetcode.com/problems/binary-tree-level-order-traversal/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
  // Membuat slice untuk menyimpan hasil
		level := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			level[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, level)
		queue = queue[levelSize:]
	}

	return result
}

func main() {
	// Test case 1: [3,9,20,null,null,15,7] -> [[3],[9,20],[15,7]]
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}}
	fmt.Println(levelOrder(root))

	// Test case 2
	fmt.Println(levelOrder(nil)) // []

	// Test case 3: [1] -> [[1]]
	fmt.Println(levelOrder(&TreeNode{Val: 1})) // [[1]]
}

// Time: O(n) | Space: O(n)
```
