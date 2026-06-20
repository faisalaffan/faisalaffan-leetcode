# 0965 — Univalued Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func isUnivalTree(root *TreeNode) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #965: Univalued Binary Tree
// https://leetcode.com/problems/univalued-binary-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}},
	}
	fmt.Println(isUnivalTree(root)) // true

	root2 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isUnivalTree(root2)) // false
}

// isUnivalTree checks if all nodes in the tree have the same value.
// Time: O(n). Space: O(n).
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsUni(root, root.Val)
}

func dfsUni(node *TreeNode, val int) bool {
	if node == nil {
		return true
	}
	if node.Val != val {
		return false
	}
	return dfsUni(node.Left, val) && dfsUni(node.Right, val)
}
```
