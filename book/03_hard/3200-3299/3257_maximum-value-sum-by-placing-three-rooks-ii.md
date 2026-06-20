# 3257 — Maximum Value Sum By Placing Three Rooks Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func maximumValueSumII(board [][]int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(m * n), Space: O(m * n)  |  **Ruang:** O(m * n)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3257: Maximum Value Sum by Placing Three Rooks II
// https://leetcode.com/problems/maximum-value-sum-by-placing-three-rooks-ii/
// Difficulty: Hard
//
// Same problem as 3256 but with larger constraints (m, n ≤ 500).
// Approach: Fix the middle row r2, then use prefix/suffix decomposition.
// Precompute top 3 values in rows above and below each row. For each cell
// in the middle row, try all combinations with prefix and suffix candidates.
//
// Time: O(m * n), Space: O(m * n)

import "fmt"

func main() {
	// Example 1: 3x3
	board := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(maximumValueSumII(board))

	// Example 2: 4x4
	board2 := [][]int{{10, 20, 30, 40}, {50, 60, 70, 80}, {90, 100, 110, 120}, {130, 140, 150, 160}}
	fmt.Println(maximumValueSumII(board2))

	// Example 3: 3x4
	board3 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(maximumValueSumII(board3))

	// Example 4: 5x3 with negatives
	board4 := [][]int{{-1, -2, -3}, {-4, -5, -6}, {-7, -8, -9}, {-10, -11, -12}, {-13, -14, -15}}
	fmt.Println(maximumValueSumII(board4))

	// Example 5: 10x10 with random pattern
	board5 := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25, 26, 27, 28, 29, 30},
		{31, 32, 33, 34, 35, 36, 37, 38, 39, 40},
		{41, 42, 43, 44, 45, 46, 47, 48, 49, 50},
		{51, 52, 53, 54, 55, 56, 57, 58, 59, 60},
		{61, 62, 63, 64, 65, 66, 67, 68, 69, 70},
		{71, 72, 73, 74, 75, 76, 77, 78, 79, 80},
		{81, 82, 83, 84, 85, 86, 87, 88, 89, 90},
		{91, 92, 93, 94, 95, 96, 97, 98, 99, 100},
	}
	fmt.Println(maximumValueSumII(board5))
}

type cand struct {
	val int64
	col int
}

func maximumValueSumII(board [][]int) int64 {
	m := len(board)
	if m < 3 {
		return 0
	}
	n := len(board[0])
	if n < 3 {
		return 0
	}

	ng := int64(-1 << 60)

	// Keep top 3 candidates by value, ensuring distinct columns
	update := func(top3 []cand, val int64, col int) []cand {
		for k := 0; k < 3; k++ {
			if val > top3[k].val {
				// Shift and insert
				dup := false
				for p := 0; p < k; p++ {
					if top3[p].col == col {
						dup = true
						break
					}
				}
				if dup {
					continue
				}
				// shift right
				for p := 2; p > k; p-- {
					top3[p] = top3[p-1]
				}
				top3[k] = cand{val, col}
				break
			}
		}
		return top3
	}

	init3 := []cand{{ng, -1}, {ng, -1}, {ng, -1}}

	// pref[i] = top 3 cells in rows [0..i]
  // Matriks 2D
	pref := make([][]cand, m)
	top := make([]cand, 3)
	copy(top, init3)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			top = update(top, int64(board[i][j]), j)
		}
		pref[i] = make([]cand, 3)
		copy(pref[i], top)
	}

	// suff[i] = top 3 cells in rows [i..m-1]
  // Matriks 2D
	suff := make([][]cand, m)
	top = make([]cand, 3)
	copy(top, init3)
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			top = update(top, int64(board[i][j]), j)
		}
		suff[i] = make([]cand, 3)
		copy(suff[i], top)
	}

	var ans int64 = ng

	// Fix middle row r2
	for r2 := 1; r2 < m-1; r2++ {
		// For each column in the middle row
		for c2 := 0; c2 < n; c2++ {
			v2 := int64(board[r2][c2])
			// Try all combinations from pref[r2-1] and suff[r2+1]
			for _, p := range pref[r2-1] {
				if p.col == -1 {
					continue
				}
				for _, s := range suff[r2+1] {
					if s.col == -1 {
						continue
					}
					if p.col != c2 && s.col != c2 && p.col != s.col {
						sum := p.val + v2 + s.val
						if sum > ans {
							ans = sum
						}
					}
				}
			}
		}
	}

	if ans == ng {
		return 0
	}
	return ans
}
```
