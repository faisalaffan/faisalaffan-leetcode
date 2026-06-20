# 1026 — Maximum Difference Between Node And Ancestor

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func maxAncestorDiff(root *TreeNode) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(n)  |  **Ruang:** O(h) where h is tree height

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1026: Maximum Difference Between Node and Ancestor
// https://leetcode.com/problems/maximum-difference-between-node-and-ancestor/
// Difficulty: Medium
//
// Approach: DFS tracking min and max values on path from root to leaf
// Time: O(n)
// Space: O(h) where h is tree height

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 8,
		Left: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 1, Left: nil, Right: nil},
			Right: &TreeNode{Val: 6, Left: &TreeNode{Val: 4, Left: nil, Right: nil}, Right: &TreeNode{Val: 7, Left: nil, Right: nil}},
		},
		Right: &TreeNode{
			Val:   10,
			Left:  nil,
			Right: &TreeNode{Val: 14, Left: &TreeNode{Val: 13, Left: nil, Right: nil}, Right: nil},
		},
	}
	fmt.Println(maxAncestorDiff(root)) // 7

	root2 := &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 0, Left: &TreeNode{Val: 3, Left: nil, Right: nil}, Right: nil}}}
	fmt.Println(maxAncestorDiff(root2)) // 3
}

func maxAncestorDiff(root *TreeNode) int {
	result := 0
	dfs(root, root.Val, root.Val, &result)
	return result
}

func dfs(node *TreeNode, minVal, maxVal int, result *int) {
	if node == nil {
		return
	}

	diff := node.Val - minVal
	if diff < 0 {
		diff = -diff
	}
	if diff > *result {
		*result = diff
	}
	diff = node.Val - maxVal
	if diff < 0 {
		diff = -diff
	}
	if diff > *result {
		*result = diff
	}

	if node.Val < minVal {
		minVal = node.Val
	}
	if node.Val > maxVal {
		maxVal = node.Val
	}

	dfs(node.Left, minVal, maxVal, result)
	dfs(node.Right, minVal, maxVal, result)
}
```
