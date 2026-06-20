# 3111 — Minimum Rectangles To Cover Points

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minRectanglesToCoverPoints(points [][]int, w int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(log n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3111: Minimum Rectangles to Cover Points
// https://leetcode.com/problems/minimum-rectangles-to-cover-points/
// Difficulty: Medium
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func minRectanglesToCoverPoints(points [][]int, w int) int {
  // Alokasi slice
	xs := make([]int, len(points))
	for i, p := range points {
		xs[i] = p[0]
	}
  // Sort O(n log n)
	sort.Ints(xs)

	ans := 0
	i := 0
	for i < len(xs) {
		ans++
		end := xs[i] + w
		for i < len(xs) && xs[i] <= end {
			i++
		}
	}
	return ans
}

func main() {
	fmt.Println(minRectanglesToCoverPoints([][]int{{2, 1}, {1, 0}, {1, 4}, {1, 8}, {3, 5}, {4, 6}}, 1)) // Expected: 2
	fmt.Println(minRectanglesToCoverPoints([][]int{{0, 0}, {1, 1}, {2, 2}, {3, 3}}, 2)) // Expected: 2
}
```
