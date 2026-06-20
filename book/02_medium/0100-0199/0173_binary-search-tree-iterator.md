# 0173 — Binary Search Tree Iterator

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func Constructor(root *TreeNode) BSTIterator`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, Stack

**Waktu:** O(1) average per next/hasNext, Space: O(h)  |  **Ruang:** O(h)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

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
