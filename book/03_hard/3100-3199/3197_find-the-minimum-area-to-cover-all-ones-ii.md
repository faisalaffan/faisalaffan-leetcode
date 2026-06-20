# 3197 — Find The Minimum Area To Cover All Ones Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minAreaRect(grid [][]int, top, bottom, left, right int) (int, bool)`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3197: Find the Minimum Area to Cover All Ones II
// https://leetcode.com/problems/find-the-minimum-area-to-cover-all-ones-ii/
// Difficulty: Hard
//
// Given a binary matrix, find the minimum total area of up to two
// non-overlapping axis-aligned rectangles whose union contains every 1-cell.
//
// Approach: try all horizontal and vertical splits. For each side compute the
// smallest rectangle covering all 1's in that side. Take the minimum sum.

import (
	"fmt"
	"math"
)

func minAreaRect(grid [][]int, top, bottom, left, right int) (int, bool) {
	// Returns area of rectangle covering all 1's in the bounded region, and
	// whether any 1 exists.
	minR, maxR := math.MaxInt32, -1
	minC, maxC := math.MaxInt32, -1
	found := false

	for r := top; r <= bottom; r++ {
		for c := left; c <= right; c++ {
			if grid[r][c] == 1 {
				found = true
				if r < minR {
					minR = r
				}
				if r > maxR {
					maxR = r
				}
				if c < minC {
					minC = c
				}
				if c > maxC {
					maxC = c
				}
			}
		}
	}
	if !found {
		return 0, false
	}
	return (maxR - minR + 1) * (maxC - minC + 1), true
}

func minAreaCoverOnesIi(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	best := math.MaxInt32

	// 1) Try all horizontal splits.
	for split := 0; split < rows-1; split++ {
		a1, ok1 := minAreaRect(grid, 0, split, 0, cols-1)
		a2, ok2 := minAreaRect(grid, split+1, rows-1, 0, cols-1)
		if ok1 && ok2 && a1+a2 < best {
			best = a1 + a2
		}
	}

	// 2) Try all vertical splits.
	for split := 0; split < cols-1; split++ {
		a1, ok1 := minAreaRect(grid, 0, rows-1, 0, split)
		a2, ok2 := minAreaRect(grid, 0, rows-1, split+1, cols-1)
		if ok1 && ok2 && a1+a2 < best {
			best = a1 + a2
		}
	}

	// 3) Also consider single rectangle covering everything (if only one
	// rectangle is needed, the second can be empty/degenerate).
	if area, ok := minAreaRect(grid, 0, rows-1, 0, cols-1); ok && area < best {
		best = area
	}

	if best == math.MaxInt32 {
		return 0
	}
	return best
}

func main() {
	grid := [][]int{
		{1, 0},
		{0, 1},
	}
	fmt.Println(minAreaCoverOnesIi(grid)) // expect 2 (two 1x1 rects)
}
```
