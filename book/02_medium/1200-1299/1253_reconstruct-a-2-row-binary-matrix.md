# 1253 — Reconstruct A 2 Row Binary Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah pohon (tree) — struktur data hierarkis dengan node (simpul) dan edge (cabang). Tugasmu adalah menjelajahi pohon tersebut (traversal), mencari nilai, atau menghitung properti tertentu.

Bayangkan struktur organisasi perusahaan: ada CEO (root), VP (children), Manager (grandchildren). Setiap node bisa punya 0 atau lebih anak. Pohon di Go direpresentasikan dengan struct yang memiliki pointer ke children (Left, Right untuk binary tree).

**Konsep kunci:** root, leaf, parent, child, depth, traversal (pre-order, in-order, post-order), recursive DFS.

**Fungsi yang perlu kamu implementasikan:**
```go
func reconstructMatrix(upper int, lower int, colsum []int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1253: Reconstruct a 2-Row Binary Matrix
// https://leetcode.com/problems/reconstruct-a-2-row-binary-matrix/
// Difficulty: Medium

// colsum[i] = sum of col i (top + bottom).
// Fill col with 2 first (top=1, bottom=1), then 1s.
// Greedy: use top row capacity first.

// Time: O(n)
// Space: O(n)

func reconstructMatrix(upper int, lower int, colsum []int) [][]int {
	n := len(colsum)
  // Alokasi slice integer
	top := make([]int, n)
  // Alokasi slice integer
	bottom := make([]int, n)

	for i, s := range colsum {
		if s == 2 {
			top[i] = 1
			bottom[i] = 1
			upper--
			lower--
		}
	}

	if upper < 0 || lower < 0 {
		return [][]int{}
	}

	for i, s := range colsum {
		if s == 1 {
			if upper > 0 {
				top[i] = 1
				upper--
			} else if lower > 0 {
				bottom[i] = 1
				lower--
			} else {
				return [][]int{}
			}
		}
	}

	if upper != 0 || lower != 0 {
		return [][]int{}
	}

	return [][]int{top, bottom}
}

func main() {
	fmt.Printf("%v (expected: [[1 1 0 0] [0 0 1 1]])\n",
		reconstructMatrix(2, 2, []int{1, 1, 1, 1}))

	fmt.Printf("%v (expected: [[]])\n",
		reconstructMatrix(2, 1, []int{1, 1, 1}))

	fmt.Printf("%v\n",
		reconstructMatrix(5, 5, []int{2, 1, 2, 0, 1, 2}))
}
```
