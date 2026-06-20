# 1337 — The K Weakest Rows In A Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func kWeakestRows(mat [][]int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(m * n + m log m), Space: O(m)  |  **Ruang:** O(m)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1337: The K Weakest Rows in a Matrix
// https://leetcode.com/problems/the-k-weakest-rows-in-a-matrix/
// Difficulty: Easy
//
// LeetCode submission: func kWeakestRows(mat [][]int, k int) []int

import (
	"fmt"
	"sort"
)

func main() {
	mat1 := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	fmt.Println(TheKWeakestRowsInAMatrix(mat1, 3)) // [2 0 3]

	mat2 := [][]int{
		{1, 0, 0, 0},
		{1, 1, 1, 1},
		{1, 0, 0, 0},
		{1, 1, 0, 0},
	}
	fmt.Println(TheKWeakestRowsInAMatrix(mat2, 2)) // [0 2]
}

// Time: O(m * n + m log m), Space: O(m)
func TheKWeakestRowsInAMatrix(mat [][]int, k int) []int {
  // Alokasi slice
	strength := make([][2]int, len(mat))
	for i, row := range mat {
		s := 0
		for _, v := range row {
			if v == 0 {
				break
			}
			s++
		}
		strength[i] = [2]int{s, i}
	}

  // Custom sort
	sort.Slice(strength, func(i, j int) bool {
		if strength[i][0] != strength[j][0] {
			return strength[i][0] < strength[j][0]
		}
		return strength[i][1] < strength[j][1]
	})

  // Alokasi slice
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = strength[i][1]
	}
	return res
}
```
