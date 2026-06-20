# 0056 — Merge Intervals

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func merge(intervals [][]int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #56: Merge Intervals
// https://leetcode.com/problems/merge-intervals/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

  // Custom sort
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{intervals[0]}

	for i := 1; i < len(intervals); i++ {
		last := result[len(result)-1]
		if intervals[i][0] <= last[1] {
			if intervals[i][1] > last[1] {
				last[1] = intervals[i][1]
			}
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
	// [[1 6] [8 10] [15 18]]

	// Test case 2
	fmt.Println(merge([][]int{{1, 4}, {4, 5}})) // [[1 5]]

	// Test case 3
	fmt.Println(merge([][]int{{1, 4}, {2, 3}})) // [[1 4]]
}

// Time: O(n log n) | Space: O(n)
```
