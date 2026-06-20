# 0104 — Maximum Depth Of Binary Tree

## Deskripsi

**Soal:** [0104. Maximum Depth Of Binary Tree](https://leetcode.com/problems/maximum-depth-of-binary-tree/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h)

**Algoritma:** —

**Fungsi Solusi:** `func MaxDepth(root *TreeNode) int`

## Solusi Go

```go
package main

// LeetCode #104: Maximum Depth of Binary Tree
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(MaxDepth(root))
	fmt.Println(MaxDepth(nil))
}
```
