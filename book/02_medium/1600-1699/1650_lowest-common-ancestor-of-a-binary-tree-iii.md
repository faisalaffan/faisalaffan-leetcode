# 1650 — Lowest Common Ancestor Of A Binary Tree Iii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func LowestCommonAncestorIII(p *Node, q *Node) *Node
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(H), Space: O(1)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1650: Lowest Common Ancestor of a Binary Tree III
// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-iii/
// Difficulty: Medium [Paid]

import "fmt"

// Node with parent pointer.
type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func main() {
	// Tree: [3,5,1,6,2,0,8,null,null,7,4]
	root := &Node{Val: 3}
	n5 := &Node{Val: 5, Parent: root}
	n1 := &Node{Val: 1, Parent: root}
	root.Left = n5
	root.Right = n1

	n6 := &Node{Val: 6, Parent: n5}
	n2 := &Node{Val: 2, Parent: n5}
	n5.Left = n6
	n5.Right = n2

	n0 := &Node{Val: 0, Parent: n1}
	n8 := &Node{Val: 8, Parent: n1}
	n1.Left = n0
	n1.Right = n8

	n7 := &Node{Val: 7, Parent: n2}
	n4 := &Node{Val: 4, Parent: n2}
	n2.Left = n7
	n2.Right = n4

	p := n5 // 5
	q := n1 // 1
	lca := LowestCommonAncestorIII(p, q)
	fmt.Println("LCA of 5 and 1:", lca.Val) // 3

	p = n5 // 5
	q = n4 // 4
	lca = LowestCommonAncestorIII(p, q)
	fmt.Println("LCA of 5 and 4:", lca.Val) // 5
}

func LowestCommonAncestorIII(p *Node, q *Node) *Node {
	// Time: O(H), Space: O(1)
	// Same as intersection of two linked lists
	a, b := p, q
	for a != b {
		if a == nil {
			a = q
		} else {
			a = a.Parent
		}
		if b == nil {
			b = p
		} else {
			b = b.Parent
		}
	}
	return a
}
```
