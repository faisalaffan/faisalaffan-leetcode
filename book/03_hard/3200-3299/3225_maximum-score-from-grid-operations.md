# 3225 — Maximum Score From Grid Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maximumScore(grid [][]int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP, Prefix Sum

**Waktu:** O(n^3), Space: O(n^2)  |  **Ruang:** O(n^2)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3225: Maximum Score From Grid Operations
// https://leetcode.com/problems/maximum-score-from-grid-operations/
// Difficulty: Hard
//
// DP with two states: for each column we track the "black height" (number of
// black cells from the top). A white cell (i,j) scores if i >= h_j AND
// i < max(h_{j-1}, h_{j+1}). We process columns left-to-right maintaining
// two DP arrays: pick[curr] and skip[curr] for the current column's height.
//
// Time: O(n^3), Space: O(n^2)

import "fmt"

func main() {
	// Example 1: 1x1 grid
	fmt.Println(maximumScore([][]int{{5}}))
	// Example 2: simple 2x2
	fmt.Println(maximumScore([][]int{{1, 2}, {3, 4}}))
	// Example 3: all zeros
	fmt.Println(maximumScore([][]int{{0, 0}, {0, 0}}))
	// Example 4: 3x3
	fmt.Println(maximumScore([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}))
	// Example 5: 4x4
	fmt.Println(maximumScore([][]int{{10, 20, 30, 40}, {50, 60, 70, 80}, {90, 100, 110, 120}, {130, 140, 150, 160}}))
}

func maximumScore(grid [][]int) int64 {
	n := len(grid)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 0
	}

	// prefix[col][row] = sum of grid[0..row-1][col]
  // Matriks 2D
	prefix := make([][]int64, n)
	for c := 0; c < n; c++ {
		prefix[c] = make([]int64, n+1)
		for r := 0; r < n; r++ {
			prefix[c][r+1] = prefix[c][r] + int64(grid[r][c])
		}
	}

  // Alokasi slice
	pick := make([]int64, n+1)
  // Alokasi slice
	skip := make([]int64, n+1)

	for col := 1; col < n; col++ {
  // Alokasi slice
		currPick := make([]int64, n+1)
  // Alokasi slice
		currSkip := make([]int64, n+1)

		for hCurr := 0; hCurr <= n; hCurr++ {
			for hPrev := 0; hPrev <= n; hPrev++ {
				if hCurr > hPrev {
					// Current column extends deeper than previous.
					// White cells in previous column (rows hPrev..hCurr-1) score.
					score := prefix[col-1][hCurr] - prefix[col-1][hPrev]
					if skip[hPrev]+score > currPick[hCurr] {
						currPick[hCurr] = skip[hPrev] + score
					}
					if skip[hPrev]+score > currSkip[hCurr] {
						currSkip[hCurr] = skip[hPrev] + score
					}
				} else {
					// Current column is shallower (or equal).
					// White cells in current column (rows hCurr..hPrev-1) score.
					score := prefix[col][hPrev] - prefix[col][hCurr]
					if pick[hPrev]+score > currPick[hCurr] {
						currPick[hCurr] = pick[hPrev] + score
					}
					if pick[hPrev] > currSkip[hCurr] {
						currSkip[hCurr] = pick[hPrev]
					}
				}
			}
		}

		pick = currPick
		skip = currSkip
	}

	var ans int64
	for _, v := range pick {
		if v > ans {
			ans = v
		}
	}
	return ans
}
```
