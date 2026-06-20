# 2768 — Number Of Black Blocks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func NumberOfBlackBlocks(m int, n int, coordinates [][]int) []int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2768: Number of Black Blocks
// https://leetcode.com/problems/number-of-black-blocks/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func NumberOfBlackBlocks(m int, n int, coordinates [][]int) []int64 {
  // HashMap: O(1) lookup
	blackCells := make(map[[2]int]bool)
	for _, c := range coordinates {
		blackCells[[2]int{c[0], c[1]}] = true
	}

  // HashMap: O(1) lookup
	blockCount := make(map[int]int64) // count of black cells in 2x2 -> number of blocks
	for _, c := range coordinates {
		r, c2 := c[0], c[1]
		// Check 4 possible top-left corners
		for _, dr := range []int{-1, 0} {
			for _, dc := range []int{-1, 0} {
				tr, tc := r+dr, c2+dc
				if tr < 0 || tc < 0 || tr >= m-1 || tc >= n-1 {
					continue
				}
				cnt := 0
				if blackCells[[2]int{tr, tc}] {
					cnt++
				}
				if blackCells[[2]int{tr, tc + 1}] {
					cnt++
				}
				if blackCells[[2]int{tr + 1, tc}] {
					cnt++
				}
				if blackCells[[2]int{tr + 1, tc + 1}] {
					cnt++
				}
				blockCount[cnt]++
			}
		}
	}

  // Alokasi slice
	result := make([]int64, 5)
	totalBlocks := int64(m-1) * int64(n-1)
	var counted int64
	for k, v := range blockCount {
		result[k] = v
		counted += v
	}
	result[0] = totalBlocks - counted
	return result
}

func main() {
	fmt.Println(NumberOfBlackBlocks(3, 3, [][]int{{0, 0}}))
	fmt.Println(NumberOfBlackBlocks(2, 2, [][]int{{0, 0}, {1, 1}}))
}
```
