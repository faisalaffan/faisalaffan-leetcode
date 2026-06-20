# 0156 — Binary Tree Upside Down

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func upsideDownBinaryTree(root *TreeNode) *TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n), Space: O(h) for recursion  
**Kompleksitas Ruang:** O(h) for recursion

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #156: Binary Tree Upside Down
// https://leetcode.com/problems/binary-tree-upside-down/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(h) for recursion

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}

	newRoot := upsideDownBinaryTree(root.Left)
	root.Left.Left = root.Right
	root.Left.Right = root
	root.Left = nil
	root.Right = nil

	return newRoot
}

func main() {
	root := &TreeNode{1, &TreeNode{2, &TreeNode{4, nil, nil}, &TreeNode{5, nil, nil}}, &TreeNode{3, nil, nil}}
	newRoot := upsideDownBinaryTree(root)
	fmt.Println(newRoot.Val)

	root2 := &TreeNode{1, &TreeNode{2, nil, nil}, nil}
	newRoot2 := upsideDownBinaryTree(root2)
	fmt.Println(newRoot2.Val)

	var root3 *TreeNode
	newRoot3 := upsideDownBinaryTree(root3)
	fmt.Println(newRoot3)
}
```
