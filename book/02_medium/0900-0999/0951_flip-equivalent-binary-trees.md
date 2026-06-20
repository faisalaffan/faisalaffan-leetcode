# 0951 — Flip Equivalent Binary Trees

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #951: Flip Equivalent Binary Trees
// https://leetcode.com/problems/flip-equivalent-binary-trees/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(n)
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}
	if root1 == nil || root2 == nil || root1.Val != root2.Val {
		return false
	}
	return (flipEquiv(root1.Left, root2.Left) && flipEquiv(root1.Right, root2.Right)) ||
		(flipEquiv(root1.Left, root2.Right) && flipEquiv(root1.Right, root2.Left))
}

func main() {
	// [1,2,3,4,5,6,null,null,null,7,8]
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{
				Val: 5,
				Left:  &TreeNode{Val: 7},
				Right: &TreeNode{Val: 8},
			},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
		},
	}
	// [1,3,2,null,6,4,5,null,null,null,null,8,7]
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 6},
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Right: &TreeNode{Val: 8},
			},
			Right: &TreeNode{
				Val:   5,
				Left:  &TreeNode{Val: 7},
			},
		},
	}
	fmt.Println(flipEquiv(root1, root2))

	// Trivial: both nil
	fmt.Println(flipEquiv(nil, nil))

	// Trivial: one nil
	fmt.Println(flipEquiv(&TreeNode{Val: 1}, nil))
}
```
