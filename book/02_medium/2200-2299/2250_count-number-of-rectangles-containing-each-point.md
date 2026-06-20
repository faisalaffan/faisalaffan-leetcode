# 2250 — Count Number Of Rectangles Containing Each Point

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func countRectangles(rectangles [][]int, points [][]int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search, Sorting

**Waktu:** O((n + m) log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2250: Count Number of Rectangles Containing Each Point
// https://leetcode.com/problems/count-number-of-rectangles-containing-each-point/
// Difficulty: Medium
// Time: O((n + m) log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func countRectangles(rectangles [][]int, points [][]int) []int {
	// Group rectangles by height
  // Matriks 2D
	byHeight := make([][]int, 101)
	for _, r := range rectangles {
		h := r[1]
		byHeight[h] = append(byHeight[h], r[0])
	}
	for h := 0; h <= 100; h++ {
  // Sort O(n log n)
		sort.Ints(byHeight[h])
	}

  // Alokasi slice
	result := make([]int, len(points))
	for i, p := range points {
		x, y := p[0], p[1]
		count := 0
		for h := y; h <= 100; h++ {
			if len(byHeight[h]) == 0 {
				continue
			}
			// Binary search for first rectangle with width >= x
			idx := sort.SearchInts(byHeight[h], x)
			count += len(byHeight[h]) - idx
		}
		result[i] = count
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println(countRectangles([][]int{{1, 2}, {2, 3}, {2, 5}}, [][]int{{2, 1}, {1, 4}}))
	// Expected: [2, 1]

	// Test case 2
	fmt.Println(countRectangles([][]int{{1, 1}, {2, 2}, {3, 3}}, [][]int{{1, 3}, {1, 1}}))
	// Expected: [1, 3]
}
```
