# 1586 — Binary Search Tree Iterator Ii

## Deskripsi

**Soal:** [1586. Binary Search Tree Iterator Ii](https://leetcode.com/problems/binary-search-tree-iterator-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner)

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
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
