# 0435 — Non Overlapping Intervals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func eraseOverlapIntervals(intervals [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #435: Non-overlapping Intervals
// https://leetcode.com/problems/non-overlapping-intervals/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	// Sort by end
  // Custom sort
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})

	count := 0
	end := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}))
	// Expected: 1

	// Test case 2
	fmt.Println("Test 2:", eraseOverlapIntervals([][]int{{1, 2}, {1, 2}, {1, 2}}))
	// Expected: 2

	// Test case 3
	fmt.Println("Test 3:", eraseOverlapIntervals([][]int{{1, 2}, {2, 3}}))
	// Expected: 0
}
```
