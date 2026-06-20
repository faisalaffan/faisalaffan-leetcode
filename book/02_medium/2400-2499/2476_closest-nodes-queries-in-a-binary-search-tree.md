# 2476 — Closest Nodes Queries In A Binary Search Tree

## Deskripsi

**Soal:** [2476. Closest Nodes Queries In A Binary Search Tree](https://leetcode.com/problems/closest-nodes-queries-in-a-binary-search-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + q log n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Binary Search (pencarian biner)

## Solusi Go

```go
package main

// LeetCode #2476: Closest Nodes Queries in a Binary Search Tree
// https://leetcode.com/problems/closest-nodes-queries-in-a-binary-search-tree/
// Difficulty: Medium
// Time: O(n + q log n) | Space: O(n)
// Inorder traversal to sorted array, then binary search each query.

import (
	"fmt"
	"sort"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{6,
		&TreeNode{2,
			&TreeNode{1, nil, nil},
			&TreeNode{4, nil, nil},
		},
		&TreeNode{13,
			&TreeNode{9, nil, nil},
			&TreeNode{15, nil, nil},
		},
	}
	fmt.Println(closestNodes(root, []int{2, 5, 16})) // [[2,2],[2,4],[15,15]]

	root2 := &TreeNode{4, nil, nil}
	fmt.Println(closestNodes(root2, []int{1, 5})) // [[-1,4],[4,-1]]
}

func closestNodes(root *TreeNode, queries []int) [][]int {
  // Membuat slice untuk menyimpan hasil
	vals := make([]int, 0)
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		vals = append(vals, node.Val)
		inorder(node.Right)
	}
	inorder(root)

  // Membuat slice 2D untuk DP/tabel
	ans := make([][]int, len(queries))
	for i, q := range queries {
		// Find smallest >= q
		idx := sort.SearchInts(vals, q)
		minVal := -1
		if idx < len(vals) {
			minVal = vals[idx]
		}
		maxVal := -1
		if idx > 0 {
			maxVal = vals[idx-1]
		}
		if idx < len(vals) && vals[idx] == q {
			// exact match
			minVal, maxVal = q, q
		}
		// Fix: if exact match, both are q
		if minVal == q && maxVal == -1 {
			maxVal = q
		}
		if minVal == q {
			maxVal = q
			minVal = q
		} else {
			// min is smallest >= q, max is largest < q
			if idx < len(vals) {
				minVal = vals[idx]
			}
			if idx > 0 {
				maxVal = vals[idx-1]
			}
		}
		ans[i] = []int{maxVal, minVal}
	}
	return ans
}
```
