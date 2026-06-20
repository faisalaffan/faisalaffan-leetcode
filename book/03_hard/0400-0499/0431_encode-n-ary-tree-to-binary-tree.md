# 0431 — Encode N Ary Tree To Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import "fmt"

// LeetCode #431: Encode N-ary Tree to Binary Tree
// https://leetcode.com/problems/encode-n-ary-tree-to-binary-tree/
// Difficulty: Hard
//
// Left pointer = first child, Right pointer = next sibling.
// Encode: first child goes to Left, subsequent children chain via Right.
// Decode: collect Left and its Right chain as children.

func main() {
	// Build N-ary tree: 1 -> [3,2,4]; 3 -> [5,6]
	root := &NNode{Val: 1}
	n3 := &NNode{Val: 3}
	n2 := &NNode{Val: 2}
	n4 := &NNode{Val: 4}
	n5 := &NNode{Val: 5}
	n6 := &NNode{Val: 6}
	root.Children = []*NNode{n3, n2, n4}
	n3.Children = []*NNode{n5, n6}

	codec := &NaryCodec{}
	bt := codec.encode(root)
	fmt.Println("Encode root:", bt.Val)

	decoded := codec.decode(bt)
	fmt.Println("Decode root:", decoded.Val)
	fmt.Println("Decode children:", len(decoded.Children))

	// Nil tree
	fmt.Println("Encode nil:", codec.encode(nil))
	fmt.Println("Decode nil:", codec.decode(nil))

	// Leaf node
	leaf := &NNode{Val: 7}
	btLeaf := codec.encode(leaf)
	decodedLeaf := codec.decode(btLeaf)
	fmt.Println("Leaf decoded val:", decodedLeaf.Val)
	fmt.Println("Leaf children:", len(decodedLeaf.Children))
}

// NNode is an N-ary tree node.
type NNode struct {
	Val      int
	Children []*NNode
}

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NaryCodec encodes/decodes N-ary <-> Binary trees.
type NaryCodec struct{}

func (c *NaryCodec) encode(root *NNode) *TreeNode {
	if root == nil {
		return nil
	}
	node := &TreeNode{Val: root.Val}
	if len(root.Children) > 0 {
		node.Left = c.encode(root.Children[0])
	}
	cur := node.Left
	for i := 1; i < len(root.Children); i++ {
		cur.Right = c.encode(root.Children[i])
		cur = cur.Right
	}
	return node
}

func (c *NaryCodec) decode(root *TreeNode) *NNode {
	if root == nil {
		return nil
	}
	node := &NNode{Val: root.Val}
	cur := root.Left
	for cur != nil {
		node.Children = append(node.Children, c.decode(cur))
		cur = cur.Right
	}
	if node.Children == nil {
		node.Children = []*NNode{}
	}
	return node
}
```
