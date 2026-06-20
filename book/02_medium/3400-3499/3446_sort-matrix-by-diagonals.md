# 3446 — Sort Matrix By Diagonals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan data yang perlu diurutkan dengan aturan tertentu. Tugasmu adalah mengurutkan data tersebut dan mungkin melakukan operasi tambahan setelah terurut.

Mengurutkan data adalah operasi fundamental di computer science. Go menyediakan `sort.Ints()` untuk integer, `sort.Strings()` untuk string, dan `sort.Slice()` untuk custom sorting dengan closure.

**Konsep kunci:** comparator, ascending/descending, stable sort, custom sort key.

**Fungsi yang perlu kamu implementasikan:**
```go
func sortMatrix(grid [][]int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer

**Kompleksitas Waktu:** O(n^2 log n) Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3446: Sort Matrix by Diagonals
// https://leetcode.com/problems/sort-matrix-by-diagonals/
// Difficulty: Medium
// Time: O(n^2 log n) Space: O(n)

import (
	"fmt"
	"slices"
)

func sortMatrix(grid [][]int) [][]int {
	n := len(grid)
	if n <= 1 {
		return grid
	}

	// Bottom-left (including main diagonal) — descending
	for i := 0; i < n; i++ {
		r, c := i, 0
		length := n - i
  // Alokasi slice integer
		diag := make([]int, length)
		for j := 0; j < length; j++ {
			diag[j] = grid[r+j][c+j]
		}
		slices.SortFunc(diag, func(a, b int) int { return b - a })
		for j := 0; j < length; j++ {
			grid[r+j][c+j] = diag[j]
		}
	}

	// Top-right (excluding main diagonal) — ascending
	for j := 1; j < n; j++ {
		r, c := 0, j
		length := n - j
  // Alokasi slice integer
		diag := make([]int, length)
		for k := 0; k < length; k++ {
			diag[k] = grid[r+k][c+k]
		}
		slices.Sort(diag)
		for k := 0; k < length; k++ {
			grid[r+k][c+k] = diag[k]
		}
	}
	return grid
}

func main() {
	fmt.Println(sortMatrix([][]int{{1, 7, 3}, {9, 8, 2}, {4, 5, 6}}))
	// [[8 2 3] [9 6 7] [4 5 1]]
	fmt.Println(sortMatrix([][]int{{0, 1}, {2, 3}}))
	// [[2 1] [2 3]]
}
```
