# 0101 — Symmetric Tree

## Deskripsi

**Soal:** [0101. Symmetric Tree](https://leetcode.com/problems/symmetric-tree/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h)

**Algoritma:** —

**Fungsi Solusi:** `func IsSymmetric(root *TreeNode) bool`

## Solusi Go

```go
package main

// LeetCode #101: Symmetric Tree
// https://leetcode.com/problems/symmetric-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func IsSymmetric(root *TreeNode) bool {
	var check func(*TreeNode, *TreeNode) bool
	check = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil || a.Val != b.Val {
			return false
		}
		return check(a.Left, b.Right) && check(a.Right, b.Left)
	}
	return check(root, root)
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{3, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{3, nil, nil}}}
	fmt.Println(IsSymmetric(root))
	root2 := &TreeNode{1, &TreeNode{2, nil, &TreeNode{3, nil, nil}}, &TreeNode{2, nil, &TreeNode{3, nil, nil}}}
	fmt.Println(IsSymmetric(root2))
}
```
