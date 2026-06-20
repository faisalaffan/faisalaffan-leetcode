# 1586 — Binary Search Tree Iterator Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func ConstructorBST(root *TreeNode) BSTIterator
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search, Stack

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1586: Binary Search Tree Iterator II
// https://leetcode.com/problems/binary-search-tree-iterator-ii/
// Difficulty: Medium [Paid]

import "fmt"

// TreeNode is a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Tree: [7, 3, 15, null, null, 9, 20]
	root := &TreeNode{Val: 7}
	root.Left = &TreeNode{Val: 3}
	root.Right = &TreeNode{Val: 15, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20}}

	it := ConstructorBST(root)
	fmt.Println("Next:", it.Next())      // 3
	fmt.Println("Next:", it.Next())      // 7
	fmt.Println("HasPrev:", it.HasPrev()) // true
	fmt.Println("Prev:", it.Prev())      // 3
	fmt.Println("Next:", it.Next())      // 7
	fmt.Println("Next:", it.Next())      // 9
	fmt.Println("Next:", it.Next())      // 15
	fmt.Println("HasNext:", it.HasNext()) // true
	fmt.Println("Next:", it.Next())      // 20
	fmt.Println("HasNext:", it.HasNext()) // false
}

type BSTIterator struct {
	stack []*TreeNode
	pos   int
	order []int
}

func ConstructorBST(root *TreeNode) BSTIterator {
	return BSTIterator{order: inorder(root)}
}

func inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
  // Alokasi slice integer
	result := make([]int, 0)
	result = append(result, inorder(root.Left)...)
	result = append(result, root.Val)
	result = append(result, inorder(root.Right)...)
	return result
}

func (it *BSTIterator) HasNext() bool {
	return it.pos < len(it.order)
}

func (it *BSTIterator) Next() int {
	val := it.order[it.pos]
	it.pos++
	return val
}

func (it *BSTIterator) HasPrev() bool {
	return it.pos > 1
}

func (it *BSTIterator) Prev() int {
	it.pos--
	return it.order[it.pos-1]
}
```
