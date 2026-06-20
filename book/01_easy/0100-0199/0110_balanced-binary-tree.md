# 0110 — Balanced Binary Tree

## Deskripsi

**Soal:** [0110. Balanced Binary Tree](https://leetcode.com/problems/balanced-binary-tree/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h)

**Algoritma:** —

**Fungsi Solusi:** `func IsBalanced(root *TreeNode) bool`

## Solusi Go

```go
package main

// LeetCode #110: Balanced Binary Tree
// https://leetcode.com/problems/balanced-binary-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func IsBalanced(root *TreeNode) bool {
	var height func(*TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := height(node.Left)
		r := height(node.Right)
		if l == -1 || r == -1 || l-r > 1 || r-l > 1 {
			return -1
		}
		if l > r {
			return l + 1
		}
		return r + 1
	}
	return height(root) != -1
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(IsBalanced(root))
	root2 := &TreeNode{1, &TreeNode{2, &TreeNode{3, &TreeNode{4, nil, nil}, nil}, nil}, &TreeNode{2, nil, nil}}
	fmt.Println(IsBalanced(root2))
}
```
