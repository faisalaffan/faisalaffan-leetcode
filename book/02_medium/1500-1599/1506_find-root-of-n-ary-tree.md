# 1506 — Find Root Of N Ary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func FindRoot(tree []*Node) *Node
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1506: Find Root of N-Ary Tree
// https://leetcode.com/problems/find-root-of-n-ary-tree/
// Difficulty: Medium [Paid]

import "fmt"

// Node is an N-ary tree node.
type Node struct {
	Val      int
	Children []*Node
}

func main() {
	// Build tree: root = 1 -> [2, 3, 4]; 3 -> [5, 6]
	child5 := &Node{Val: 5}
	child6 := &Node{Val: 6}
	child2 := &Node{Val: 2}
	child3 := &Node{Val: 3, Children: []*Node{child5, child6}}
	child4 := &Node{Val: 4}
	root := &Node{Val: 1, Children: []*Node{child2, child3, child4}}

	// All nodes in random order (without knowing root)
	allNodes := []*Node{child2, child5, child3, child4, root, child6}

	found := FindRoot(allNodes)
	fmt.Println("Found root val:", found.Val)
}

func FindRoot(tree []*Node) *Node {
	// Time: O(N), Space: O(1)
	// The root is the only node that is never a child.
	// XOR all node values + all child values. Root value remains.
	var xorSum int
	for _, node := range tree {
		xorSum ^= node.Val
		for _, child := range node.Children {
			xorSum ^= child.Val
		}
	}

	// Find node with matching value
	for _, node := range tree {
		if node.Val == xorSum {
			return node
		}
	}
	return nil
}
```
