# 2001 — Number Of Pairs Of Interchangeable Rectangles

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func NumberOfPairsOfInterchangeableRectangles(rectangles [][]int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n log max(w,h)), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2001: Number of Pairs of Interchangeable Rectangles
// https://leetcode.com/problems/number-of-pairs-of-interchangeable-rectangles/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(NumberOfPairsOfInterchangeableRectangles([][]int{{4, 8}, {3, 6}, {10, 20}, {15, 30}}))
	fmt.Println(NumberOfPairsOfInterchangeableRectangles([][]int{{4, 5}, {7, 8}}))
}

// Time: O(n log max(w,h)), Space: O(n)
func NumberOfPairsOfInterchangeableRectangles(rectangles [][]int) int64 {
  // HashMap: O(1) lookup
	cnt := make(map[[2]int]int64)
	for _, r := range rectangles {
		w, h := r[0], r[1]
		g := gcd2001(w, h)
		key := [2]int{w / g, h / g}
		cnt[key]++
	}

	var ans int64
	for _, m := range cnt {
		ans += m * (m - 1) / 2
	}
	return ans
}

func gcd2001(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
```
