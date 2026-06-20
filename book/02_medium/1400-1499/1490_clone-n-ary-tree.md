# 1490 — Clone N Ary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func CloneTree(root *Node) *Node
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Stack

**Kompleksitas Waktu:** O(N), Space: O(N) (recursion stack)  
**Kompleksitas Ruang:** O(N) (recursion stack)

> **Untuk fresh graduate:** Kuasai dulu teknik **Stack** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1490: Clone N-ary Tree
// https://leetcode.com/problems/clone-n-ary-tree/
// Difficulty: Medium [Paid]

import "fmt"

// Node is an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

func main() {
	// Build tree: root = 1 -> [3, 2, 4]; 3 -> [5, 6]
	root := &Node{Val: 1}
	child3 := &Node{Val: 3}
	child2 := &Node{Val: 2}
	child4 := &Node{Val: 4}
	root.Children = []*Node{child3, child2, child4}
	child3.Children = []*Node{{Val: 5}, {Val: 6}}

	cloned := CloneTree(root)
	fmt.Println("Root cloned:", cloned != nil && cloned != root)
	fmt.Println("Root val:", cloned.Val)
	fmt.Println("Children count:", len(cloned.Children))
	fmt.Println("Deep cloned:", cloned.Children[0].Children[0].Val == 5 && cloned.Children[0] != child3.Children[0])

	// Test nil
	fmt.Println(CloneTree(nil))
}

func CloneTree(root *Node) *Node {
	// Time: O(N), Space: O(N) (recursion stack)
	if root == nil {
		return nil
	}

	clone := &Node{Val: root.Val}
	clone.Children = make([]*Node, len(root.Children))
	for i, child := range root.Children {
		clone.Children[i] = CloneTree(child)
	}
	return clone
}
```
