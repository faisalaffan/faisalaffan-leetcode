# 2345 — Finding The Number Of Visible Mountains

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func visibleMountains(mountains [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2345: Finding the Number of Visible Mountains
// https://leetcode.com/problems/finding-the-number-of-visible-mountains/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func visibleMountains(mountains [][]int) int {
	// Each mountain is peak at [x, y], base at [x-y, x+y]
	type interval struct{ left, right int }
  // Alokasi slice
	intervals := make([]interval, len(mountains))
	for i, m := range mountains {
		x, y := m[0], m[1]
		intervals[i] = interval{x - y, x + y}
	}

	// Sort by left ascending, right descending
  // Custom sort
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i].left != intervals[j].left {
			return intervals[i].left < intervals[j].left
		}
		return intervals[i].right > intervals[j].right
	})

	visible := 0
	maxRight := -1 << 30
	i := 0

	for i < len(intervals) {
		if intervals[i].right <= maxRight {
			i++
			continue
		}
		// Check if this mountain is not a duplicate
		if i+1 < len(intervals) && intervals[i].left == intervals[i+1].left && intervals[i].right == intervals[i+1].right {
			// Duplicate - skip all with same interval
			j := i
			for j < len(intervals) && intervals[j].left == intervals[i].left && intervals[j].right == intervals[i].right {
				j++
			}
			maxRight = intervals[i].right
			i = j
			continue
		}
		visible++
		maxRight = intervals[i].right
		i++
	}
	return visible
}

func main() {
	// Test case 1
	fmt.Println(visibleMountains([][]int{{2, 2}, {6, 3}, {5, 4}}))
	// Expected: 2

	// Test case 2
	fmt.Println(visibleMountains([][]int{{1, 2}, {1, 2}, {2, 1}}))
	// Expected: 1
}
```
