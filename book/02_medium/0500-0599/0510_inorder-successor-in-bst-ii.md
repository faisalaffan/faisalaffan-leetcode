# 0510 — Inorder Successor In Bst Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func InorderSuccessorInBstIi(node *Node) *Node
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(h) where h is height of tree  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #510: Inorder Successor in BST II
// https://leetcode.com/problems/inorder-successor-in-bst-ii/
// Difficulty: Medium [Paid]
// Time: O(h) where h is height of tree
// Space: O(1)

import "fmt"

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func main() {
	// Build tree: [2,1,3]
	root := &Node{Val: 2}
	root.Left = &Node{Val: 1, Parent: root}
	root.Right = &Node{Val: 3, Parent: root}
	fmt.Println(InorderSuccessorInBstIi(root.Left).Val) // node 1 -> successor 2
	fmt.Println(InorderSuccessorInBstIi(root).Val)       // node 2 -> successor 3

	// For node 3, successor should be nil
	successor := InorderSuccessorInBstIi(root.Right)
	if successor == nil {
		fmt.Println("nil")
	} else {
		fmt.Println(successor.Val)
	}
}

func InorderSuccessorInBstIi(node *Node) *Node {
	if node == nil {
		return nil
	}

	// If right child exists, find leftmost in right subtree
	if node.Right != nil {
		cur := node.Right
		for cur.Left != nil {
			cur = cur.Left
		}
		return cur
	}

	// Otherwise, go up until we find a node that is a left child
	cur := node
	for cur.Parent != nil && cur.Parent.Right == cur {
		cur = cur.Parent
	}

	return cur.Parent
}
```
