# 0111 — Minimum Depth Of Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func MinDepth(root *TreeNode) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(h)


## 💻 Solusi Go

```go
package main

// LeetCode #111: Minimum Depth of Binary Tree
// https://leetcode.com/problems/minimum-depth-of-binary-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return MinDepth(root.Right) + 1
	}
	if root.Right == nil {
		return MinDepth(root.Left) + 1
	}
	l, r := MinDepth(root.Left), MinDepth(root.Right)
	if l < r {
		return l + 1
	}
	return r + 1
}

func main() {
	root := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	fmt.Println(MinDepth(root))
	root2 := &TreeNode{2, nil, &TreeNode{3, nil, &TreeNode{4, nil, &TreeNode{5, nil, &TreeNode{6, nil, nil}}}}}
	fmt.Println(MinDepth(root2))
}
```
