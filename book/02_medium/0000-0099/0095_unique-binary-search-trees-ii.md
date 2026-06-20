# 0095 — Unique Binary Search Trees Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func generateTrees(n int) []*TreeNode
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Binary Search

**Kompleksitas Waktu:** O(4^n / n^(3/2))  
**Kompleksitas Ruang:** O(4^n / n^(3/2))

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Edge case: input kosong — langsung return
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
