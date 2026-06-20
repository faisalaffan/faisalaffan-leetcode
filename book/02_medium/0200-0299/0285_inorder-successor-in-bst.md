# 0285 — Inorder Successor In Bst

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(h), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #285: Inorder Successor in BST
// https://leetcode.com/problems/inorder-successor-in-bst/
// Difficulty: Medium [Paid]
// Time: O(h), Space: O(1)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var successor *TreeNode

	for root != nil {
		if p.Val < root.Val {
			successor = root
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return successor
}

func main() {
	root := &TreeNode{5, &TreeNode{3, &TreeNode{2, nil, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, &TreeNode{7, nil, nil}}}
	p := root.Left
	fmt.Println(inorderSuccessor(root, p).Val)

	p2 := root.Right
	fmt.Println(inorderSuccessor(root, p2))

	p3 := root.Left.Right
	fmt.Println(inorderSuccessor(root, p3).Val)
}
```
