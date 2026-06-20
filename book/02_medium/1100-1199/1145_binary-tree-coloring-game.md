# 1145 — Binary Tree Coloring Game

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func btreeGameWinningMove(root *TreeNode, n int, x int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(h) where h is tree height (recursion)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1145: Binary Tree Coloring Game
// https://leetcode.com/problems/binary-tree-coloring-game/
// Difficulty: Medium

// Two players color a binary tree. Player 1 colors node x blue.
// Player 2 can choose any other node as red. Then they alternate.
// A move = color an uncolored neighbor of your colored node.
// Player 2 wins if they can color more nodes.

// Player 1's initial node x splits tree into 3 components:
// left subtree, right subtree, and rest of tree.
// Player 2 should choose the root of the largest component.
// Player 2 wins if max(leftSize, rightSize, restSize) > n/2.

// Time: O(n)
// Space: O(h) where h is tree height (recursion)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	var xNode *TreeNode
	var findX func(node *TreeNode)
	findX = func(node *TreeNode) {
		if node == nil || xNode != nil {
			return
		}
		if node.Val == x {
			xNode = node
			return
		}
		findX(node.Left)
		findX(node.Right)
	}
	findX(root)

	var count func(node *TreeNode) int
	count = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		return 1 + count(node.Left) + count(node.Right)
	}

	leftSize := count(xNode.Left)
	rightSize := count(xNode.Right)
	restSize := n - leftSize - rightSize - 1

	maxSize := leftSize
	if rightSize > maxSize {
		maxSize = rightSize
	}
	if restSize > maxSize {
		maxSize = restSize
	}

	return maxSize > n/2
}

func main() {
	// Test: root = [1,2,3,4,5,6,7,8,9,10,11], n = 11, x = 3
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{Val: 4,
				Left:  &TreeNode{Val: 8},
				Right: &TreeNode{Val: 9}},
			Right: &TreeNode{Val: 5,
				Left: &TreeNode{Val: 10},
				Right: &TreeNode{Val: 11}}},
		Right: &TreeNode{
			Val: 3,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7}},
	}
	fmt.Printf("%t (expected: true)\n", btreeGameWinningMove(root, 11, 3))

	// Simple tree: [1,2,3], n=3, x=1 -> left(1), right(1), rest(0) -> max=1 <= 1.5 -> false
	root2 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	fmt.Printf("%t (expected: false)\n", btreeGameWinningMove(root2, 3, 1))

	// Single node
	root3 := &TreeNode{Val: 1}
	fmt.Printf("%t (expected: false)\n", btreeGameWinningMove(root3, 1, 1))
}
```
