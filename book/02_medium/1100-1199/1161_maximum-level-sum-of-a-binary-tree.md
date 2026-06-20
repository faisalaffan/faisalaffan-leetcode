# 1161 — Maximum Level Sum Of A Binary Tree

## Deskripsi

**Soal:** [1161. Maximum Level Sum Of A Binary Tree](https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** BFS (Breadth-First Search / pencarian lebar), Queue (antrian FIFO)

**Fungsi Solusi:** `func maxLevelSum(root *TreeNode) int`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1161: Maximum Level Sum of a Binary Tree
// https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
// Difficulty: Medium

// BFS level order traversal, track sum per level.

// Time: O(n)
// Space: O(n)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	maxSum := root.Val
	maxLevel := 1
	currentLevel := 0

	for len(queue) > 0 {
		currentLevel++
		size := len(queue)
		sum := 0

		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		if sum > maxSum {
			maxSum = sum
			maxLevel = currentLevel
		}
	}

	return maxLevel
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 7},
			Right: &TreeNode{Val: -8},
		},
		Right: &TreeNode{Val: 0},
	}
	fmt.Printf("%d (expected: 2)\n", maxLevelSum(root))

	root2 := &TreeNode{
		Val:  -100,
		Left: &TreeNode{Val: -200},
		Right: &TreeNode{
			Val:   -300,
			Left:  &TreeNode{Val: -20},
			Right: &TreeNode{Val: -5},
		},
	}
	fmt.Printf("%d (expected: 3)\n", maxLevelSum(root2))
}
```
