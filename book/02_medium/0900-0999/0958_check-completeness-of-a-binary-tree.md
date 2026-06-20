# 0958 — Check Completeness Of A Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func isCompleteTree(root *TreeNode) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #958: Check Completeness of a Binary Tree
// https://leetcode.com/problems/check-completeness-of-a-binary-tree/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(n) | Space: O(n)
func isCompleteTree(root *TreeNode) bool {
	q := []*TreeNode{root}
	seenNull := false
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			seenNull = true
		} else {
			if seenNull {
				return false
			}
			q = append(q, node.Left, node.Right)
		}
	}
	return true
}

func main() {
	// Complete tree: [1,2,3,4,5,6]
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Left:  &TreeNode{Val: 6},
		},
	}
	fmt.Println(isCompleteTree(root))

	// Not complete: [1,2,3,4,5,null,7]
	root2 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val:   3,
			Right: &TreeNode{Val: 7},
		},
	}
	fmt.Println(isCompleteTree(root2))
}
```
