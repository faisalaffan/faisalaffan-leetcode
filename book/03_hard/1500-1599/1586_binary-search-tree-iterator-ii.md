# 1586 — Binary Search Tree Iterator Ii

## Deskripsi

**Soal:** [1586. Binary Search Tree Iterator Ii](https://leetcode.com/problems/binary-search-tree-iterator-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner)

**Fungsi Solusi:** `func Constructor(root *TreeNode) *BSTIteratorII`

> **Ide Kunci:** Perform an inorder traversal to collect all node values,

## Solusi Go

```go
package main

import "fmt"

// LeetCode #1586: Binary Search Tree Iterator II
// https://leetcode.com/problems/binary-search-tree-iterator-ii/
// Difficulty: Hard
//
// BSTIteratorII provides bidirectional traversal of a BST:
// - hasNext() / next()   — forward
// - hasPrev() / prev()   — backward
//
// Approach: Perform an inorder traversal to collect all node values,
// then use a cursor to navigate. This gives O(1) for all operations
// at the cost of O(n) memory.

// TreeNode represents a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BSTIteratorII is the bidirectional BST iterator.
type BSTIteratorII struct {
	vals    []int
	cursor  int // index of the last returned value; -1 means before start
}

// Constructor initializes the iterator with the root of a BST.
func Constructor(root *TreeNode) *BSTIteratorII {
  // Membuat slice untuk menyimpan hasil
	vals := make([]int, 0)
	inorder(root, &vals)
	return &BSTIteratorII{vals: vals, cursor: -1}
}

func inorder(node *TreeNode, vals *[]int) {
	if node == nil {
		return
	}
	inorder(node.Left, vals)
	*vals = append(*vals, node.Val)
	inorder(node.Right, vals)
}

// HasNext returns true if there is a next element.
func (it *BSTIteratorII) HasNext() bool {
	return it.cursor+1 < len(it.vals)
}

// Next returns the next element. It advances the cursor.
func (it *BSTIteratorII) Next() int {
	it.cursor++
	return it.vals[it.cursor]
}

// HasPrev returns true if there is a previous element.
func (it *BSTIteratorII) HasPrev() bool {
	return it.cursor > 0
}

// Prev returns the previous element. It retreats the cursor.
func (it *BSTIteratorII) Prev() int {
	it.cursor--
	return it.vals[it.cursor]
}

// BuildBST builds a BST from a level-order slice (null represented by -1).
func BuildBST(vals []int) *TreeNode {
	if len(vals) == 0 || vals[0] == -1 {
		return nil
	}
	root := &TreeNode{Val: vals[0]}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		node := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != -1 {
			node.Left = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(vals) && vals[i] != -1 {
			node.Right = &TreeNode{Val: vals[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func main() {
	// Example:
	// BST:
	//       7
	//      / \
	//     3   15
	//        /  \
	//       9   20
	//
	// Inorder: [3, 7, 9, 15, 20]

	root := BuildBST([]int{7, 3, 15, -1, -1, 9, 20})
	it := Constructor(root)

	fmt.Println("HasNext:", it.HasNext())      // true
	fmt.Println("Next:", it.Next())            // 3
	fmt.Println("HasPrev:", it.HasPrev())      // false (cursor at 0)
	fmt.Println("HasNext:", it.HasNext())      // true
	fmt.Println("Next:", it.Next())            // 7
	fmt.Println("Prev:", it.Prev())            // 3 (back to 3)
	fmt.Println("Next:", it.Next())            // 7 (forward again)
	fmt.Println("Next:", it.Next())            // 9
	fmt.Println("Next:", it.Next())            // 15
	fmt.Println("HasNext:", it.HasNext())      // true
	fmt.Println("Next:", it.Next())            // 20
	fmt.Println("HasNext:", it.HasNext())      // false
	fmt.Println("HasPrev:", it.HasPrev())      // true
	fmt.Println("Prev:", it.Prev())            // 15
	fmt.Println("Prev:", it.Prev())            // 9
	fmt.Println("Prev:", it.Prev())            // 7
	fmt.Println("Prev:", it.Prev())            // 3
	fmt.Println("HasPrev:", it.HasPrev())      // false
}
```
