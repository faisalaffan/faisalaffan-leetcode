# 2196 — Create Binary Tree From Descriptions

## Deskripsi

**Soal:** [2196. Create Binary Tree From Descriptions](https://leetcode.com/problems/create-binary-tree-from-descriptions/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func createBinaryTree(descriptions [][]int) *TreeNode`

## Solusi Go

```go
package main

// LeetCode #2196: Create Binary Tree From Descriptions
// https://leetcode.com/problems/create-binary-tree-from-descriptions/
// Difficulty: Medium (listed as Hard)

import "fmt"

// TreeNode definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// createBinaryTree builds a binary tree from descriptions.
// Each description: [parent, child, isLeft] where isLeft=1 means left child.
func createBinaryTree(descriptions [][]int) *TreeNode {
  // Membuat map untuk pencarian O(1): key → value
	children := make(map[int]*TreeNode)
  // Membuat map untuk pencarian O(1): key → value
	hasParent := make(map[int]bool)

	for _, d := range descriptions {
		parentVal, childVal, isLeft := d[0], d[1], d[2]

		// get or create parent
		parent, ok := children[parentVal]
		if !ok {
			parent = &TreeNode{Val: parentVal}
			children[parentVal] = parent
		}

		// get or create child
		child, ok := children[childVal]
		if !ok {
			child = &TreeNode{Val: childVal}
			children[childVal] = child
		}

		// set left/right
		if isLeft == 1 {
			parent.Left = child
		} else {
			parent.Right = child
		}

		hasParent[childVal] = true
		// ensure parent also tracked
		if !hasParent[parentVal] {
			hasParent[parentVal] = false
		}
	}

	// find root: node without a parent
	var root *TreeNode
	for val, node := range children {
		if !hasParent[val] {
			root = node
			break
		}
	}

	return root
}

func main() {
	// Example 1
	descriptions1 := [][]int{
		{20, 15, 1},
		{20, 17, 0},
		{50, 20, 1},
		{50, 80, 0},
		{80, 19, 1},
	}
	root1 := createBinaryTree(descriptions1)
	fmt.Println("Root value:", root1.Val) // Expected: 50

	// Example 2
	descriptions2 := [][]int{
		{1, 2, 1},
		{2, 3, 0},
		{3, 4, 1},
	}
	root2 := createBinaryTree(descriptions2)
	fmt.Println("Root value:", root2.Val) // Expected: 1
}
```
