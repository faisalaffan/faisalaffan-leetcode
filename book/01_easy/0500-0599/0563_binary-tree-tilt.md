# 0563 — Binary Tree Tilt

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func BinaryTreeTilt(root *TreeNode) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n), Space: O(h)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #563: Binary Tree Tilt
// https://leetcode.com/problems/binary-tree-tilt/
// Difficulty: Easy

import "fmt"

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n), Space: O(h)
func BinaryTreeTilt(root *TreeNode) int {
	totalTilt := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		tilt := left - right
		if tilt < 0 {
			tilt = -tilt
		}
		totalTilt += tilt
		return left + right + node.Val
	}
	dfs(root)
	return totalTilt
}

func main() {
	// Test: [1,2,3]
	root1 := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 3},
	}
	fmt.Println(BinaryTreeTilt(root1))

	// Test: [4,2,9,3,5,null,7]
	root2 := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 3},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   9,
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(BinaryTreeTilt(root2))
}
```
