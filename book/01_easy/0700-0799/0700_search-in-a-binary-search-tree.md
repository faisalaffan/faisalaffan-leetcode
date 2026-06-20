# 0700 — Search In A Binary Search Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func searchBST(root *TreeNode, val int) *TreeNode`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search

**Waktu:** O(log n) average, O(n) worst. Space: O(1).  |  **Ruang:** O(1).

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #700: Search in a Binary Search Tree
// https://leetcode.com/problems/search-in-a-binary-search-tree/
// Difficulty: Easy

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	// Example: [4,2,7,1,3], val=2
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{Val: 7},
	}
	fmt.Println(searchBST(root, 2))  // node with value 2
	fmt.Println(searchBST(root, 5))  // nil
}

// searchBST searches for a node with the given value in a BST.
// Time: O(log n) average, O(n) worst. Space: O(1).
func searchBST(root *TreeNode, val int) *TreeNode {
	curr := root
	for curr != nil {
		if val == curr.Val {
			return curr
		} else if val < curr.Val {
			curr = curr.Left
		} else {
			curr = curr.Right
		}
	}
	return nil
}
```
