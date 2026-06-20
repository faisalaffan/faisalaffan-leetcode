# 1038 — Binary Search Tree To Greater Sum Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func bstToGst(root *TreeNode) *TreeNode`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, DFS, BFS

**Waktu:** O(n)  |  **Ruang:** O(h) where h is tree height

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1038: Binary Search Tree to Greater Sum Tree
// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/
// Difficulty: Medium
//
// Approach: Reverse inorder (right -> root -> left) accumulating sum
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
		Val: 4,
		Left: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 0, Left: nil, Right: nil},
			Right: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Left: nil, Right: nil}},
		},
		Right: &TreeNode{
			Val:   6,
			Left:  &TreeNode{Val: 5, Left: nil, Right: nil},
			Right: &TreeNode{Val: 7, Right: &TreeNode{Val: 8, Left: nil, Right: nil}},
		},
	}
	result := bstToGst(root)
	printTree(result)
	fmt.Println()
}

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}

func printTree(root *TreeNode) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			fmt.Print("null ")
			continue
		}
		fmt.Printf("%d ", node.Val)
		queue = append(queue, node.Left, node.Right)
	}
}
```
