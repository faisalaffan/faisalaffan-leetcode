# 1373 — Maximum Sum Bst In Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSumBST(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, DFS, BFS

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1373: Maximum Sum BST in Binary Tree
// https://leetcode.com/problems/maximum-sum-bst-in-binary-tree/
// Difficulty: Hard

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type SubtreeInfo struct {
	isBST bool
	min   int
	max   int
	sum   int
}

func maxSumBST(root *TreeNode) int {
	maxSum := 0

	var dfs func(node *TreeNode) SubtreeInfo
	dfs = func(node *TreeNode) SubtreeInfo {
		if node == nil {
			return SubtreeInfo{true, math.MaxInt32, math.MinInt32, 0}
		}

		left := dfs(node.Left)
		right := dfs(node.Right)

		if left.isBST && right.isBST &&
			node.Val > left.max && node.Val < right.min {
			sum := node.Val + left.sum + right.sum
			if sum > maxSum {
				maxSum = sum
			}
			minVal := left.min
			if node.Val < minVal {
				minVal = node.Val
			}
			maxVal := right.max
			if node.Val > maxVal {
				maxVal = node.Val
			}
			return SubtreeInfo{true, minVal, maxVal, sum}
		}

		return SubtreeInfo{false, 0, 0, 0}
	}

	dfs(root)
	return maxSum
}

// buildTree builds a binary tree from level-order array representation.
// null values are represented by math.MinInt32 sentinel.
func buildTree(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] == math.MinInt32 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]

		if i < len(arr) && arr[i] != math.MinInt32 {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(arr) && arr[i] != math.MinInt32 {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func main() {
	// Example 1
	// Tree: [1,4,3,2,4,2,5,null,null,null,null,null,null,4,6]
	// where math.MinInt32 represents null
	arr := []int{1, 4, 3, 2, 4, 2, 5, math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32, 4, 6}
	root := buildTree(arr)
	fmt.Println(maxSumBST(root))
	// Expected: 20

	// Example 2
	arr2 := []int{4, 3, math.MinInt32, 1, 2}
	root2 := buildTree(arr2)
	fmt.Println(maxSumBST(root2))
	// Expected: 2

	// Example 3
	arr3 := []int{-4, -2, -5}
	root3 := buildTree(arr3)
	fmt.Println(maxSumBST(root3))
	// Expected: 0
}
```
