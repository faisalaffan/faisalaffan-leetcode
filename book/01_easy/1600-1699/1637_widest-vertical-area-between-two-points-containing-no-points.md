# 1637 — Widest Vertical Area Between Two Points Containing No Points

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func MaxWidthOfVerticalArea(points [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(log n) (for sorting)  |  **Ruang:** O(log n) (for sorting)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1637: Widest Vertical Area Between Two Points Containing No Points
// https://leetcode.com/problems/widest-vertical-area-between-two-points-containing-no-points/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(log n) (for sorting)
func MaxWidthOfVerticalArea(points [][]int) int {
  // Alokasi slice
	xs := make([]int, len(points))
	for i, p := range points {
		xs[i] = p[0]
	}
  // Sort O(n log n)
	sort.Ints(xs)
	maxWidth := 0
	for i := 1; i < len(xs); i++ {
		if xs[i]-xs[i-1] > maxWidth {
			maxWidth = xs[i] - xs[i-1]
		}
	}
	return maxWidth
}

func main() {
	fmt.Println(MaxWidthOfVerticalArea([][]int{{8, 7}, {9, 9}, {7, 4}, {9, 7}}))
	fmt.Println(MaxWidthOfVerticalArea([][]int{{3, 1}, {9, 0}, {1, 0}, {1, 4}, {5, 3}, {8, 8}}))
}
```
