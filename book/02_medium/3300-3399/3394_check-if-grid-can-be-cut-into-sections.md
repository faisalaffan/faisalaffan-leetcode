# 3394 — Check If Grid Can Be Cut Into Sections

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func check(intervals []pair) bool`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n log n) Space: O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #3394: Check if Grid can be Cut into Sections
// https://leetcode.com/problems/check-if-grid-can-be-cut-into-sections/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"slices"
)

type pair struct{ l, r int }

func check(intervals []pair) bool {
	slices.SortFunc(intervals, func(a, b pair) int { return a.l - b.l })
	cnt, maxR := 0, 0
	for _, p := range intervals {
		if p.l >= maxR {
			cnt++
		}
		if p.r > maxR {
			maxR = p.r
		}
	}
	return cnt >= 3
}

func checkValidCuts(_ int, rectangles [][]int) bool {
	n := len(rectangles)
	a := make([]pair, n)
	b := make([]pair, n)
	for i, rect := range rectangles {
		a[i] = pair{rect[0], rect[2]}
		b[i] = pair{rect[1], rect[3]}
	}
	return check(a) || check(b)
}

func main() {
	fmt.Println(checkValidCuts(5, [][]int{{1, 0, 5, 2}, {0, 2, 2, 4}, {3, 2, 5, 3}, {0, 4, 4, 5}})) // true
	fmt.Println(checkValidCuts(4, [][]int{{0, 0, 1, 1}, {2, 0, 3, 4}, {0, 2, 2, 3}, {3, 0, 4, 3}})) // true
	fmt.Println(checkValidCuts(4, [][]int{{0, 2, 2, 4}, {1, 0, 3, 2}, {2, 2, 3, 4}, {3, 0, 4, 2}, {3, 2, 4, 4}})) // false
}
```
