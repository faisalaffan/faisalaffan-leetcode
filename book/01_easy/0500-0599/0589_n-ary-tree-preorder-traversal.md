# 0589 — N Ary Tree Preorder Traversal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func NAryTreePreorderTraversal(root *Node) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** DFS

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **DFS** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #589: N-ary Tree Preorder Traversal
// https://leetcode.com/problems/n-ary-tree-preorder-traversal/
// Difficulty: Easy

import "fmt"

// Node represents an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

// Time: O(n), Space: O(n)
func NAryTreePreorderTraversal(root *Node) []int {
	var result []int
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		for _, child := range node.Children {
			dfs(child)
		}
	}
	dfs(root)
	return result
}

func main() {
	// Test: [1,null,3,2,4,null,5,6]
	root1 := &Node{
		Val: 1,
		Children: []*Node{
			{Val: 3, Children: []*Node{
				{Val: 5},
				{Val: 6},
			}},
			{Val: 2},
			{Val: 4},
		},
	}
	fmt.Println(NAryTreePreorderTraversal(root1))

	// Test: single node
	root2 := &Node{Val: 1}
	fmt.Println(NAryTreePreorderTraversal(root2))
}
```
