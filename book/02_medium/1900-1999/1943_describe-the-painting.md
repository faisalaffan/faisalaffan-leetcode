# 1943 — Describe The Painting

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func SplitPainting(segments [][]int) [][]int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1943: Describe the Painting
// https://leetcode.com/problems/describe-the-painting/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SplitPainting([][]int{{1, 4, 5}, {4, 7, 7}, {1, 7, 9}}))
	fmt.Println(SplitPainting([][]int{{1, 7, 9}, {6, 8, 15}, {8, 10, 7}}))
	fmt.Println(SplitPainting([][]int{{1, 4, 5}, {1, 4, 7}, {4, 7, 1}, {4, 7, 11}}))
}

// Time: O(n log n), Space: O(n)
func SplitPainting(segments [][]int) [][]int64 {
  // HashMap: O(1) lookup
	diff := make(map[int]int64)
  // HashMap: O(1) lookup
	endpoints := make(map[int]bool)

	for _, seg := range segments {
		start, end, color := seg[0], seg[1], seg[2]
		diff[start] += int64(color)
		diff[end] -= int64(color)
		endpoints[start] = true
		endpoints[end] = true
	}

	// Sort unique endpoints
  // Alokasi slice
	points := make([]int, 0, len(endpoints))
	for p := range endpoints {
		points = append(points, p)
	}
  // Sort O(n log n)
	sort.Ints(points)

  // Matriks 2D
	result := make([][]int64, 0)
	var sum int64 = 0
  // Linear scan O(n)
	for i := 0; i < len(points)-1; i++ {
		sum += diff[points[i]]
		if sum != 0 {
			result = append(result, []int64{int64(points[i]), int64(points[i+1]), sum})
		}
	}
	return result
}
```
