# 0100 — Same Tree

## Deskripsi

**Soal:** [0100. Same Tree](https://leetcode.com/problems/same-tree/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h is tree height

**Algoritma:** —

**Fungsi Solusi:** `func IsSameTree(p *TreeNode, q *TreeNode) bool`

## Solusi Go

```go
package main

// LeetCode #100: Same Tree
// https://leetcode.com/problems/same-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h) where h is tree height
func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

func main() {
	t1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	t2 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	fmt.Println(IsSameTree(t1, t2))
	t3 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	fmt.Println(IsSameTree(t1, t3))
}
```
