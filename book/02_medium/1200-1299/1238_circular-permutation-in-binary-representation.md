# 1238 — Circular Permutation In Binary Representation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func circularPermutation(n int, start int) []int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Fenwick Tree (BIT)

**Kompleksitas Waktu:** O(2^n)  
**Kompleksitas Ruang:** O(2^n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Fenwick Tree (BIT)** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1238: Circular Permutation in Binary Representation
// https://leetcode.com/problems/circular-permutation-in-binary-representation/
// Difficulty: Medium

// Generate Gray code sequence starting with start.
// Gray code: consecutive numbers differ by exactly 1 bit.

// Time: O(2^n)
// Space: O(2^n)

func circularPermutation(n int, start int) []int {
	size := 1 << n
  // Alokasi slice integer
	result := make([]int, size)

	for i := 0; i < size; i++ {
		result[i] = start ^ i ^ (i >> 1)
	}

	return result
}

func main() {
	fmt.Printf("%v (expected: [3 2 0 1] or similar)\n", circularPermutation(2, 3))
	fmt.Printf("%v\n", circularPermutation(3, 2))
}
```
