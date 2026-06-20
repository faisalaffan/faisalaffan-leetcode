# 3157 — Find The Level Of Tree With Minimum Sum

## Deskripsi

**Soal:** [3157. Find The Level Of Tree With Minimum Sum](https://leetcode.com/problems/find-the-level-of-tree-with-minimum-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Queue (antrian FIFO)

**Fungsi Solusi:** `func minimumLevel(root *TreeNode) int`

## Solusi Go

```go
package main

// LeetCode #3157: Find the Level of Tree with Minimum Sum
// https://leetcode.com/problems/find-the-level-of-tree-with-minimum-sum/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumLevel(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := []*TreeNode{root}
	minSum := root.Val
	minLevel := 1
	level := 1

	for len(queue) > 0 {
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

		if sum < minSum {
			minSum = sum
			minLevel = level
		}
		level++
	}

	return minLevel
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(minimumLevel(root)) // Expected: 1

	root2 := &TreeNode{10, &TreeNode{1, nil, nil}, &TreeNode{1, nil, nil}}
	fmt.Println(minimumLevel(root2)) // Expected: 1

	root3 := &TreeNode{5,
		&TreeNode{3, &TreeNode{100, nil, nil}, nil},
		&TreeNode{8, nil, nil},
	}
	fmt.Println(minimumLevel(root3)) // Expected: 2 (level 2 sum = 100, level 1 sum = 5+8=13)
}
```
