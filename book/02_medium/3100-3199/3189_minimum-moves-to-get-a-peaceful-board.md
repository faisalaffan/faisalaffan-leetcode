# 3189 — Minimum Moves To Get A Peaceful Board

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func minMoves(rooks [][]int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3189: Minimum Moves to Get a Peaceful Board
// https://leetcode.com/problems/minimum-moves-to-get-a-peaceful-board/
// Difficulty: Medium [Paid]
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func minMoves(rooks [][]int) int {
	n := len(rooks)
  // Alokasi slice
	rows := make([]int, n)
  // Alokasi slice
	cols := make([]int, n)
	for i, r := range rooks {
		rows[i] = r[0]
		cols[i] = r[1]
	}

  // Sort O(n log n)
	sort.Ints(rows)
  // Sort O(n log n)
	sort.Ints(cols)

	moves := 0
	for i := 0; i < n; i++ {
		moves += abs(rows[i] - i)
		moves += abs(cols[i] - i)
	}
	return moves
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minMoves([][]int{{0, 0}, {1, 1}, {2, 2}}))    // Expected: 0
	fmt.Println(minMoves([][]int{{0, 0}, {0, 2}, {2, 0}}))    // Expected: 2
	fmt.Println(minMoves([][]int{{2, 2}, {0, 0}, {1, 1}}))    // Expected: 0
}
```
