# 0235 — Lowest Common Ancestor Of A Binary Search Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(h), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #235: Lowest Common Ancestor of a Binary Search Tree
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// Difficulty: Medium
// Time: O(h), Space: O(1)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val < root.Val && q.Val < root.Val {
			root = root.Left
		} else if p.Val > root.Val && q.Val > root.Val {
			root = root.Right
		} else {
			return root
		}
	}
	return nil
}

func main() {
	root := &TreeNode{6, &TreeNode{2, &TreeNode{0, nil, nil}, &TreeNode{4, &TreeNode{3, nil, nil}, &TreeNode{5, nil, nil}}}, &TreeNode{8, &TreeNode{7, nil, nil}, &TreeNode{9, nil, nil}}}
	fmt.Println(lowestCommonAncestor(root, root.Left, root.Right).Val)

	p := root.Left
	q := root.Left.Right
	fmt.Println(lowestCommonAncestor(root, p, q).Val)

	p2 := root.Left
	q2 := root.Left.Left
	fmt.Println(lowestCommonAncestor(root, p2, q2).Val)
}
```
