# 2773 — Height Of Special Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func HeightOfSpecialBinaryTree(root *SpecialTreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
