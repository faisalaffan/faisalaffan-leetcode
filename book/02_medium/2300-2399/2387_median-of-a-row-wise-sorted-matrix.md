# 2387 — Median Of A Row Wise Sorted Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func matrixMedian(grid [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Binary Search

**Waktu:** O(rows * log(cols) * log(max-min))  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2387: Median of a Row Wise Sorted Matrix
// https://leetcode.com/problems/median-of-a-row-wise-sorted-matrix/
// Difficulty: Medium
// Time: O(rows * log(cols) * log(max-min)) | Space: O(1)
// Binary search on value, count elements <= mid.

import "fmt"

func main() {
	fmt.Println(matrixMedian([][]int{{1, 1, 2}, {2, 3, 3}, {1, 3, 4}})) // 2
	fmt.Println(matrixMedian([][]int{{1, 2}, {3, 4}}))                 // 2
}

func matrixMedian(grid [][]int) int {
	r, c := len(grid), len(grid[0])
	target := r*c/2 + 1
	lo, hi := 1, 1000000

	for lo < hi {
		mid := (lo + hi) / 2
		count := 0
		for i := 0; i < r; i++ {
			// binary search in each row for count of elements <= mid
			row := grid[i]
			left, right := 0, c
  // Two-pointer loop
			for left < right {
				m := (left + right) / 2
				if row[m] <= mid {
					left = m + 1
				} else {
					right = m
				}
			}
			count += left
		}
		if count >= target {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo
}
```
