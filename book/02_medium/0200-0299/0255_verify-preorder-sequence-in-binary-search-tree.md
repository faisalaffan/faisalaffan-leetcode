# 0255 — Verify Preorder Sequence In Binary Search Tree

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func verifyPreorder(preorder []int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Binary Search, Stack

**Kompleksitas Waktu:** O(n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Binary Search** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #255: Verify Preorder Sequence in Binary Search Tree
// https://leetcode.com/problems/verify-preorder-sequence-in-binary-search-tree/
// Difficulty: Medium [Paid]
// Time: O(n), Space: O(n)

import "fmt"

func verifyPreorder(preorder []int) bool {
	stack := []int{}
	lower := ^int(^uint(0) >> 1) // math.MinInt

	for _, val := range preorder {
		if val < lower {
			return false
		}
		for len(stack) > 0 && val > stack[len(stack)-1] {
			lower = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, val)
	}

	return true
}

func main() {
	fmt.Println(verifyPreorder([]int{5, 2, 1, 3, 6}))
	fmt.Println(verifyPreorder([]int{5, 2, 6, 1, 3}))
	fmt.Println(verifyPreorder([]int{1, 2, 3}))
}
```
