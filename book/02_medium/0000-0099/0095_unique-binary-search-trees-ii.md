# 0095 — Unique Binary Search Trees Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan pohon (tree). Tugasmu menjelajahi atau memanipulasi struktur pohon.

**Cara berpikir:** TreeNode punya Val, Left, Right. Gunakan DFS rekursif (pre/in/post-order).

**Fungsi Solusi:** `func generateTrees(n int) []*TreeNode`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Binary Search

**Waktu:** O(4^n / n^(3/2))  |  **Ruang:** O(4^n / n^(3/2))

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #95: Unique Binary Search Trees II
// https://leetcode.com/problems/unique-binary-search-trees-ii/
// Difficulty: Medium

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
  // Edge case: input kosong
	if n == 0 {
		return []*TreeNode{}
	}
	return build(1, n)
}

func build(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	result := []*TreeNode{}
	for i := start; i <= end; i++ {
		leftTrees := build(start, i-1)
		rightTrees := build(i+1, end)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				result = append(result, &TreeNode{Val: i, Left: left, Right: right})
			}
		}
	}
	return result
}

func printTreePreorder(root *TreeNode) {
	if root == nil {
		fmt.Print("null ")
		return
	}
	fmt.Printf("%d ", root.Val)
	printTreePreorder(root.Left)
	printTreePreorder(root.Right)
}

func main() {
	// Test case 1
	trees := generateTrees(3)
	fmt.Println(len(trees)) // 5
	for _, t := range trees {
		printTreePreorder(t)
		fmt.Println()
	}

	// Test case 2
	trees = generateTrees(1)
	fmt.Println(len(trees)) // 1
	for _, t := range trees {
		printTreePreorder(t)
		fmt.Println()
	}
}

// Time: O(4^n / n^(3/2)) | Space: O(4^n / n^(3/2))
```
