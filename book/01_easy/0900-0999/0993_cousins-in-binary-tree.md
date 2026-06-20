# 0993 — Cousins In Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func isCousins(root *TreeNode, x int, y int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DFS

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **DFS** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #993: Cousins in Binary Tree
// https://leetcode.com/problems/cousins-in-binary-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
		},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(isCousins(root, 4, 3)) // false

	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 4},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 5},
		},
	}
	fmt.Println(isCousins(root2, 5, 4)) // true
}

// isCousins checks if two nodes are cousins (same depth, different parent).
// Time: O(n). Space: O(n).
func isCousins(root *TreeNode, x int, y int) bool {
	var xDepth, yDepth int
	var xParent, yParent *TreeNode
	dfsCousins(root, nil, 0, x, y, &xDepth, &yDepth, &xParent, &yParent)
	return xDepth == yDepth && xParent != nil && yParent != nil && xParent != yParent
}

func dfsCousins(node, parent *TreeNode, depth int, x, y int, xDepth, yDepth *int, xParent, yParent **TreeNode) {
	if node == nil {
		return
	}
	if node.Val == x {
		*xDepth = depth
		*xParent = parent
	}
	if node.Val == y {
		*yDepth = depth
		*yParent = parent
	}
	dfsCousins(node.Left, node, depth+1, x, y, xDepth, yDepth, xParent, yParent)
	dfsCousins(node.Right, node, depth+1, x, y, xDepth, yDepth, xParent, yParent)
}
```
