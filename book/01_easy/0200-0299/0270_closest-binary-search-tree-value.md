# 0270 — Closest Binary Search Tree Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func ClosestValue(root *TreeNode, target float64) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search

**Kompleksitas Waktu:** O(h)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #270: Closest Binary Search Tree Value
// https://leetcode.com/problems/closest-binary-search-tree-value/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time: O(h) | Space: O(1)
func ClosestValue(root *TreeNode, target float64) int {
	closest := root.Val
	for root != nil {
		if math.Abs(float64(root.Val)-target) < math.Abs(float64(closest)-target) {
			closest = root.Val
		}
		if target < float64(root.Val) {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return closest
}

func main() {
	root := &TreeNode{4, &TreeNode{2, &TreeNode{1, nil, nil}, &TreeNode{3, nil, nil}}, &TreeNode{5, nil, nil}}
	fmt.Println(ClosestValue(root, 3.714286))
	fmt.Println(ClosestValue(&TreeNode{1, nil, nil}, 4.428571))
}
```
