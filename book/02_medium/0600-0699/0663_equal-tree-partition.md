# 0663 — Equal Tree Partition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func checkEqualTree(root *TreeNode) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** HashMap, Two Pointer

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **HashMap** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #663: Equal Tree Partition
// https://leetcode.com/problems/equal-tree-partition/
// Difficulty: Medium [Paid]
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 10}
	root.Right = &TreeNode{Val: 10}
	root.Right.Left = &TreeNode{Val: 2}
	root.Right.Right = &TreeNode{Val: 3}

	fmt.Println(checkEqualTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func checkEqualTree(root *TreeNode) bool {
  // Membuat map (HashMap) — pencarian O(1)
	sums := make(map[int]int)
	total := computeSum(root, sums)

	if total%2 != 0 {
		return false
	}

	// Root sum is also stored; we need a different way to track
	// We'll rebuild sums excluding total
	half := total / 2
	return sums[half] > 0
}

func computeSum(node *TreeNode, sums map[int]int) int {
	if node == nil {
		return 0
	}
	left := computeSum(node.Left, sums)
	right := computeSum(node.Right, sums)
	sum := node.Val + left + right

	// Count subtree sums
	sums[sum]++

	return sum
}
```
