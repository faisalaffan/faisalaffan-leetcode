# 0872 — Leaf Similar Trees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n + m). Space: O(n + m).  
**Kompleksitas Ruang:** O(n + m).

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #872: Leaf-Similar Trees
// https://leetcode.com/problems/leaf-similar-trees/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// [3,5,1,6,2,9,8,null,null,7,4]
	root1 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 4}},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 9},
			Right: &TreeNode{Val: 8},
		},
	}
	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   5,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 7},
		},
		Right: &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 8}},
		},
	}
	fmt.Println(leafSimilar(root1, root2)) // true

	root3 := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}
	root4 := &TreeNode{Val: 2, Left: &TreeNode{Val: 2}}
	fmt.Println(leafSimilar(root3, root4)) // false
}

// leafSimilar checks if two trees have the same leaf value sequence.
// Time: O(n + m). Space: O(n + m).
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
  // Alokasi slice integer
	leaves1 := make([]int, 0)
	collectLeaves(root1, &leaves1)
  // Alokasi slice integer
	leaves2 := make([]int, 0)
	collectLeaves(root2, &leaves2)
	if len(leaves1) != len(leaves2) {
		return false
	}
	for i, v := range leaves1 {
		if v != leaves2[i] {
			return false
		}
	}
	return true
}

func collectLeaves(node *TreeNode, leaves *[]int) {
	if node == nil {
		return
	}
	if node.Left == nil && node.Right == nil {
		*leaves = append(*leaves, node.Val)
		return
	}
	collectLeaves(node.Left, leaves)
	collectLeaves(node.Right, leaves)
}
```
