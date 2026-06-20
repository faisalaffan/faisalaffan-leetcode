# 0652 — Find Duplicate Subtrees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func findDuplicateSubtrees(root *TreeNode) []*TreeNode`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer, DFS

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #652: Find Duplicate Subtrees
// https://leetcode.com/problems/find-duplicate-subtrees/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import (
	"fmt"
	"strconv"
)

func main() {
	// Test case: root = [1,2,3,4,null,2,4,null,null,4]
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Left.Left = &TreeNode{Val: 4}
	root.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Left.Left = &TreeNode{Val: 4}
	root.Right.Right = &TreeNode{Val: 4}

	result := findDuplicateSubtrees(root)
	for _, node := range result {
		fmt.Println(node.Val)
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
  // HashMap: O(1) lookup
	seen := make(map[string]int)
	result := make([]*TreeNode, 0)

	var dfs func(node *TreeNode) string
	dfs = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}

		left := dfs(node.Left)
		right := dfs(node.Right)
		key := strconv.Itoa(node.Val) + "," + left + "," + right

		seen[key]++
		if seen[key] == 2 {
			result = append(result, node)
		}

		return key
	}

	dfs(root)
	return result
}
```
