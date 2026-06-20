# 2773 — Height Of Special Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func HeightOfSpecialBinaryTree(root *SpecialTreeNode) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(h)


## 💻 Solusi Go

```go
package main

// LeetCode #2773: Height of Special Binary Tree
// https://leetcode.com/problems/height-of-special-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(h)

import "fmt"

type SpecialTreeNode struct {
	Val       int
	Left      *SpecialTreeNode
	Right     *SpecialTreeNode
	IsSpecial bool
}

func HeightOfSpecialBinaryTree(root *SpecialTreeNode) int {
	var dfs func(*SpecialTreeNode) int
	dfs = func(node *SpecialTreeNode) int {
		if node == nil || node.IsSpecial {
			return 0
		}
		leftH := dfs(node.Left)
		rightH := dfs(node.Right)
		if leftH > rightH {
			return leftH + 1
		}
		return rightH + 1
	}
	return dfs(root)
}

func main() {
	// Tree: 1(not special) -> 2(special), 3(not special)
	root := &SpecialTreeNode{
		Val: 1,
		Left: &SpecialTreeNode{
			Val:       2,
			IsSpecial: true,
		},
		Right: &SpecialTreeNode{
			Val: 3,
			Left: &SpecialTreeNode{
				Val: 4,
			},
		},
	}
	fmt.Println(HeightOfSpecialBinaryTree(root))
	fmt.Println(HeightOfSpecialBinaryTree(nil))
}
```
