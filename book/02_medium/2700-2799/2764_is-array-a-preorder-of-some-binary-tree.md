# 2764 — Is Array A Preorder Of Some Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func IsArrayAPreorderOfSomeBinaryTree(nodes [][]int) bool`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Two Pointer, Stack

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2764: Is Array a Preorder of Some Binary Tree
// https://leetcode.com/problems/is-array-a-preorder-of-some-binary-tree/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func IsArrayAPreorderOfSomeBinaryTree(nodes [][]int) bool {
	// nodes[i] = [id, parentId]
  // HashMap: O(1) lookup
	children := make(map[int][]int)
	for _, node := range nodes {
		id, parent := node[0], node[1]
		children[parent] = append(children[parent], id)
	}

	// Simulate DFS preorder
	stack := []int{-1} // root parent
	idx := 0
	n := len(nodes)

	for len(stack) > 0 {
		cur := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		childList := children[cur]
		for i := len(childList) - 1; i >= 0; i-- {
			if idx >= n || childList[i] != nodes[idx][0] {
				return false
			}
			stack = append(stack, childList[i])
			idx++
		}
	}

	return idx == n
}

func main() {
	fmt.Println(IsArrayAPreorderOfSomeBinaryTree([][]int{{0, -1}, {1, 0}, {2, 0}}))
	fmt.Println(IsArrayAPreorderOfSomeBinaryTree([][]int{{0, -1}, {1, 0}, {3, 2}, {2, 1}}))
}
```
