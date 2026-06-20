# 1659 — Maximize Grid Happiness

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func getMaxGridHappiness(m int, n int, introvertsCount int, extrovertsCount int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
	"math"
)

// LeetCode #1659: Maximize Grid Happiness
// https://leetcode.com/problems/maximize-grid-happiness/
// Difficulty: Hard
//
// Place introverts (I) and extroverts (E) in an m x n grid.
// - Introvert happiness = 120 (base) - 30 per neighbor (any I or E nearby)
// - Extrovert happiness = 40 (base) + 20 per neighbor (any I or E nearby)
// - Empty cell = 0
// Maximize total happiness with at most introvertsCount introverts
// and extrovertsCount extroverts.
//
// DP with state compression: process row by row. For each row, we track
// the placement of the previous row as a 3-base (ternary) number where
// 0 = empty, 1 = introvert, 2 = extrovert. n <= 5 so 3^5 = 243 states.
//
// dp[row][maskPrev][i][e] = max happiness up to row `row` when previous row
// is `maskPrev`, with i introverts and e extroverts used so far.

func getMaxGridHappiness(m int, n int, introvertsCount int, extrovertsCount int) int {
	// Total states for one row: 3^n
	states := int(math.Pow(3, float64(n)))

	// Precompute for each state:
	// - rowScore: total happiness within the row
	// - intro count, extro count
  // Alokasi slice
	rowScore := make([]int, states)
  // Alokasi slice
	introCount := make([]int, states)
  // Alokasi slice
	extroCount := make([]int, states)

	for s := 0; s < states; s++ {
		prev := -1
		tmp := s
		inner := 0
		introCnt := 0
		extroCnt := 0
  // Alokasi slice
		mask := make([]int, n)
		for pos := 0; pos < n; pos++ {
			cell := tmp % 3
			tmp /= 3
			mask[pos] = cell
			if cell == 1 {
				introCnt++
				inner += 120
			} else if cell == 2 {
				extroCnt++
				inner += 40
			}
			// Horizontal interaction with left neighbor
			if prev != -1 {
				if prev == 1 && cell == 1 {
					inner -= 60 // I-I
				} else if prev == 1 && cell == 2 {
					inner -= 10 // I-E (I loses 30, E gains 20)
				} else if prev == 2 && cell == 1 {
					inner -= 10 // E-I (same)
				} else if prev == 2 && cell == 2 {
					inner += 40 // E-E (each gains 20)
				}
			}
			prev = cell
		}
		rowScore[s] = inner
		introCount[s] = introCnt
		extroCount[s] = extroCnt

		// Precompute vertical interaction score with the same state (placeholder)
		// Will compute actual inter-row score on the fly.
	}

	// Precompute vertical interaction between two states (top row and bottom row)
  // Matriks 2D
	vertBetween := make([][]int, states)
	for s1 := 0; s1 < states; s1++ {
		vertBetween[s1] = make([]int, states)
		for s2 := 0; s2 < states; s2++ {
			score := 0
			tmp1, tmp2 := s1, s2
			for pos := 0; pos < n; pos++ {
				top := tmp1 % 3
				bottom := tmp2 % 3
				tmp1 /= 3
				tmp2 /= 3
				if top == 1 && bottom == 1 {
					score -= 60
				} else if top == 1 && bottom == 2 {
					score -= 10
				} else if top == 2 && bottom == 1 {
					score -= 10
				} else if top == 2 && bottom == 2 {
					score += 40
				}
			}
			vertBetween[s1][s2] = score
		}
	}

	// dp[mask][i][e] = max happiness for processed rows with previous row = mask,
	// i introverts used, e extroverts used
  // Matriks 2D
	dp := make([][][]int, states)
	for s := 0; s < states; s++ {
		dp[s] = make([][]int, introvertsCount+1)
		for i := 0; i <= introvertsCount; i++ {
			dp[s][i] = make([]int, extrovertsCount+1)
			for e := 0; e <= extrovertsCount; e++ {
				dp[s][i][e] = math.MinInt32
			}
		}
	}
	dp[0][0][0] = 0 // mask 0 = all empty

	for r := 0; r < m; r++ {
  // Matriks 2D
		ndp := make([][][]int, states)
		for s := 0; s < states; s++ {
			ndp[s] = make([][]int, introvertsCount+1)
			for i := 0; i <= introvertsCount; i++ {
				ndp[s][i] = make([]int, extrovertsCount+1)
				for e := 0; e <= extrovertsCount; e++ {
					ndp[s][i][e] = math.MinInt32
				}
			}
		}

		for prevMask := 0; prevMask < states; prevMask++ {
			for i := 0; i <= introvertsCount; i++ {
				for e := 0; e <= extrovertsCount; e++ {
					cur := dp[prevMask][i][e]
					if cur == math.MinInt32 {
						continue
					}
					// Try every possible state for the current row
					for curMask := 0; curMask < states; curMask++ {
						ni := i + introCount[curMask]
						ne := e + extroCount[curMask]
						if ni > introvertsCount || ne > extrovertsCount {
							continue
						}
						score := cur + rowScore[curMask] + vertBetween[prevMask][curMask]
						if score > ndp[curMask][ni][ne] {
							ndp[curMask][ni][ne] = score
						}
					}
				}
			}
		}
		dp = ndp
	}

	ans := 0
	for mask := 0; mask < states; mask++ {
		for i := 0; i <= introvertsCount; i++ {
			for e := 0; e <= extrovertsCount; e++ {
				if dp[mask][i][e] > ans {
					ans = dp[mask][i][e]
				}
			}
		}
	}
	return ans
}

func main() {
	// Example 1:
	// Input: m = 2, n = 3, introvertsCount = 1, extrovertsCount = 2
	// Output: 240
	fmt.Println(getMaxGridHappiness(2, 3, 1, 2))

	// Example 2:
	// Input: m = 3, n = 1, introvertsCount = 2, extrovertsCount = 1
	// Output: 260
	fmt.Println(getMaxGridHappiness(3, 1, 2, 1))

	// Example 3:
	// Input: m = 2, n = 2, introvertsCount = 4, extrovertsCount = 0
	// Output: 240
	fmt.Println(getMaxGridHappiness(2, 2, 4, 0))
}
```
