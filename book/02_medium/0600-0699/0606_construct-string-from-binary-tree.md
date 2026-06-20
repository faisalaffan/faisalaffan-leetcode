# 0606 — Construct String From Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func Tree2str(root *TreeNode) string`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(h) where h is tree height


## 💻 Solusi Go

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
