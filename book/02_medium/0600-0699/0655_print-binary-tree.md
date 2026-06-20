# 0655 — Print Binary Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func printTree(root *TreeNode) [][]string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n * h) where h is height  
**Kompleksitas Ruang:** O(n * h)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #655: Print Binary Tree
// https://leetcode.com/problems/print-binary-tree/
// Difficulty: Medium
// Time: O(n * h) where h is height
// Space: O(n * h)

import (
	"fmt"
	"strconv"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Right = &TreeNode{Val: 4}

	fmt.Println(printTree(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func printTree(root *TreeNode) [][]string {
	height := getHeight(root)
	width := (1 << height) - 1
  // Membuat matriks/slice 2D untuk DP
	result := make([][]string, height)
  // Range loop: iterasi dengan indeks + nilai
	for i := range result {
		result[i] = make([]string, width)
	}

	fill(root, result, 0, 0, width-1)
	return result
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getHeight(root.Left), getHeight(root.Right))
}

func fill(root *TreeNode, result [][]string, level int, left int, right int) {
	if root == nil {
		return
	}
	mid := (left + right) / 2
	result[level][mid] = strconv.Itoa(root.Val)
	fill(root.Left, result, level+1, left, mid-1)
	fill(root.Right, result, level+1, mid+1, right)
}
```
