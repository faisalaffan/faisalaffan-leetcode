# 0606 — Construct String From Binary Tree

## Deskripsi

**Soal:** [0606. Construct String From Binary Tree](https://leetcode.com/problems/construct-string-from-binary-tree/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h is tree height

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #606: Construct String from Binary Tree
// https://leetcode.com/problems/construct-string-from-binary-tree/
// Difficulty: Medium
// Time: O(n)
// Space: O(h) where h is tree height

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}}
	fmt.Println(Tree2str(root))
}

func Tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	result := strconv.Itoa(root.Val)
	if root.Left != nil || root.Right != nil {
		result += "(" + Tree2str(root.Left) + ")"
	}
	if root.Right != nil {
		result += "(" + Tree2str(root.Right) + ")"
	}
	return result
}
```
