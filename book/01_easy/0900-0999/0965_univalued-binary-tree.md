# 0965 — Univalued Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func isUnivalTree(root *TreeNode) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n). Space: O(n).  
**Kompleksitas Ruang:** O(n).

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #965: Univalued Binary Tree
// https://leetcode.com/problems/univalued-binary-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 1}},
		Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 1}},
	}
	fmt.Println(isUnivalTree(root)) // true

	root2 := &TreeNode{
		Val:   2,
		Left:  &TreeNode{Val: 2, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 5}},
		Right: &TreeNode{Val: 2},
	}
	fmt.Println(isUnivalTree(root2)) // false
}

// isUnivalTree checks if all nodes in the tree have the same value.
// Time: O(n). Space: O(n).
func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfsUni(root, root.Val)
}

func dfsUni(node *TreeNode, val int) bool {
	if node == nil {
		return true
	}
	if node.Val != val {
		return false
	}
	return dfsUni(node.Left, val) && dfsUni(node.Right, val)
}
```
