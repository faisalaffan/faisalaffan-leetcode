# 1372 — Longest Zigzag Path In A Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func longestZigZag(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS, Stack

**Kompleksitas Waktu:** O(n) where n = number of nodes  
**Kompleksitas Ruang:** O(h) for recursion stack

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1372: Longest ZigZag Path in a Binary Tree
// https://leetcode.com/problems/longest-zigzag-path-in-a-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Test case 1
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val: 1,
				Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}},
				Right: &TreeNode{Val: 1},
			},
		},
	}
	fmt.Println(longestZigZag(root)) // 3

	// Test case 2: single node
	fmt.Println(longestZigZag(&TreeNode{Val: 1})) // 0

	// Test case 3
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
	}
	fmt.Println(longestZigZag(root2)) // 2
}

// Time: O(n) where n = number of nodes
// Space: O(h) for recursion stack
func longestZigZag(root *TreeNode) int {
	maxLen := 0
	var dfs func(*TreeNode, int, bool) // node, current length, isLeft (true = coming from left)
	dfs = func(node *TreeNode, length int, fromLeft bool) {
		if node == nil {
			return
		}
		if length > maxLen {
			maxLen = length
		}
		// Going to left child
		if fromLeft {
			// Continuing zigzag: was going left, now going left = reset
			dfs(node.Left, 1, true)
			// Change direction: was going left, now going right = continue zigzag
			dfs(node.Right, length+1, false)
		} else {
			// Change direction: was going right, now going left = continue zigzag
			dfs(node.Left, length+1, true)
			// Continuing same direction: was going right, now going right = reset
			dfs(node.Right, 1, false)
		}
	}

	dfs(root.Left, 1, true)
	dfs(root.Right, 1, false)

	return maxLen
}
```
