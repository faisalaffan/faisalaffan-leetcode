# 0298 — Binary Tree Longest Consecutive Sequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func longestConsecutive(root *TreeNode) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DFS

**Waktu:** O(n), Space: O(h)  |  **Ruang:** O(h)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #298: Binary Tree Longest Consecutive Sequence
// https://leetcode.com/problems/binary-tree-longest-consecutive-sequence/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestConsecutive(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLen := 1
	var dfs func(node *TreeNode, parentVal int, length int)
	dfs = func(node *TreeNode, parentVal int, length int) {
		if node == nil {
			return
		}

		if node.Val == parentVal+1 {
			length++
		} else {
			length = 1
		}

		if length > maxLen {
			maxLen = length
		}

		dfs(node.Left, node.Val, length)
		dfs(node.Right, node.Val, length)
	}

	dfs(root, root.Val, 1)
	return maxLen
}

func main() {
	root := &TreeNode{1, nil, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, &TreeNode{5, nil, nil}}}}
	fmt.Println(longestConsecutive(root))

	root2 := &TreeNode{2, nil, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, nil}}
	fmt.Println(longestConsecutive(root2))

	fmt.Println(longestConsecutive(nil))
}
```
