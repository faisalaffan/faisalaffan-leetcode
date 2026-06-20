# 0257 — Binary Tree Paths

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func BinaryTreePaths(root *TreeNode) []string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(n)  |  **Ruang:** O(h)

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #257: Binary Tree Paths
// https://leetcode.com/problems/binary-tree-paths/
// Difficulty: Easy

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(h)
func BinaryTreePaths(root *TreeNode) []string {
	var res []string
	var dfs func(*TreeNode, string)
	dfs = func(node *TreeNode, path string) {
		if node == nil {
			return
		}
		path += strconv.Itoa(node.Val)
		if node.Left == nil && node.Right == nil {
			res = append(res, path)
			return
		}
		path += "->"
		dfs(node.Left, path)
		dfs(node.Right, path)
	}
	dfs(root, "")
	return res
}

func main() {
	root := &TreeNode{1, &TreeNode{2, nil, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	fmt.Println(BinaryTreePaths(root))
	fmt.Println(BinaryTreePaths(nil))
}
```
