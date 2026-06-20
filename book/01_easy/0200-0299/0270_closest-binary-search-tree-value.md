# 0270 — Closest Binary Search Tree Value

## Deskripsi

**Soal:** [0270. Closest Binary Search Tree Value](https://leetcode.com/problems/closest-binary-search-tree-value/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(h)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Binary Search (pencarian biner)

**Fungsi Solusi:** `func ClosestValue(root *TreeNode, target float64) int`

## Solusi Go

```go
package main

// LeetCode #270: Closest Binary Search Tree Value
// https://leetcode.com/problems/closest-binary-search-tree-value/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(h) | Space: O(1)
func ClosestValue(root *TreeNode, target float64) int {
	closest := root.Val
	for root != nil {
		if math.Abs(float64(root.Val)-target) < math.Abs(float64(closest)-target) {
			closest = root.Val
		}
		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return closest
}

func main() {
	root := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{5, nil, nil}}
	fmt.Println(ClosestValue(root, 3.714286))
	fmt.Println(ClosestValue(&TreeNode{1, nil, nil}, 4.428571))
}
```
