# 0230 — Kth Smallest Element In A Bst

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func kthSmallest(root *TreeNode, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Stack

**Kompleksitas Waktu:** O(h+k), Space: O(h)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #230: Kth Smallest Element in a BST
// https://leetcode.com/problems/kth-smallest-element-in-a-bst/
// Difficulty: Medium
// Time: O(h+k), Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	node := root

	for {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}

		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		k--

		if k == 0 {
			return node.Val
		}

		node = node.Right
	}
}

func main() {
	root := &TreeNode{3, &TreeNode{1, nil, &TreeNode{2, nil, nil}}, &TreeNode{4, nil, nil}}
	fmt.Println(kthSmallest(root, 1))

	root2 := &TreeNode{5, &TreeNode{3, &TreeNode{2, &TreeNode{1, nil, nil}, nil}, &TreeNode{4, nil, nil}}, &TreeNode{6, nil, nil}}
	fmt.Println(kthSmallest(root2, 3))
}
```
