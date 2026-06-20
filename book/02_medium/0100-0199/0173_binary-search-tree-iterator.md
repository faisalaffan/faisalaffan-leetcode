# 0173 — Binary Search Tree Iterator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func Constructor(root *TreeNode) BSTIterator
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, Stack

**Kompleksitas Waktu:** O(1) average per next/hasNext, Space: O(h)  
**Kompleksitas Ruang:** O(h)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #173: Binary Search Tree Iterator
// https://leetcode.com/problems/binary-search-tree-iterator/
// Difficulty: Medium
// Time: O(1) average per next/hasNext, Space: O(h)

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	iter := BSTIterator{}
	iter.pushLeft(root)
	return iter
}

func (this *BSTIterator) pushLeft(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}

func (this *BSTIterator) Next() int {
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.pushLeft(node.Right)
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) > 0
}

func main() {
	root := &TreeNode{7, &TreeNode{3, nil, nil}, &TreeNode{15, &TreeNode{9, nil, nil}, &TreeNode{20, nil, nil}}}
	iter := Constructor(root)
	fmt.Println(iter.Next())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
	fmt.Println(iter.Next())
	fmt.Println(iter.HasNext())
}
```
