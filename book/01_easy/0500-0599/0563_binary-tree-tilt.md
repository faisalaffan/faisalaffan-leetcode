# 0563 — Binary Tree Tilt

## Deskripsi

**Soal:** [0563. Binary Tree Tilt](https://leetcode.com/problems/binary-tree-tilt/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n), Space: O(h)  
**Kompleksitas Ruang:** O(h)

**Algoritma:** DFS (Depth-First Search / pencarian kedalaman)

**Fungsi Solusi:** `func BinaryTreeTilt(root *TreeNode) int`

## Solusi Go

```go
package main

// LeetCode #563: Binary Tree Tilt
// https://leetcode.com/problems/binary-tree-tilt/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(h)
func BinaryTreeTilt(root *TreeNode) int {
	totalTilt := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		tilt := left - right
		if tilt < 0 {
			tilt = -tilt
		}
		totalTilt += tilt
		return left + right + node.Val
	}
	dfs(root)
	return totalTilt
}

func main() {
	// Test: [1,2,3]
	root1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(BinaryTreeTilt(root1))

	// Test: [4,2,9,3,5,null,7]
	root2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   9,
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(BinaryTreeTilt(root2))
}
```
