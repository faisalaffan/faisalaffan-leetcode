# 0662 — Maximum Width Of Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func widthOfBinaryTree(root *TreeNode) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, BFS

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #662: Maximum Width of Binary Tree
// https://leetcode.com/problems/maximum-width-of-binary-tree/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 5}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 9}

	fmt.Println(widthOfBinaryTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}

	type pair struct {
		node *TreeNode
		idx  int
	}

	queue := []pair{{root, 0}}
	maxWidth := 0

	for len(queue) > 0 {
		n := len(queue)
		first := queue[0].idx
		last := queue[n-1].idx
		maxWidth = max(maxWidth, last-first+1)

		for i := 0; i < n; i++ {
			node, idx := queue[i].node, queue[i].idx
			if node.Left != nil {
				queue = append(queue, pair{node.Left, idx * 2})
			}
			if node.Right != nil {
				queue = append(queue, pair{node.Right, idx*2 + 1})
			}
		}
		queue = queue[n:]
	}

	return maxWidth
}
```
