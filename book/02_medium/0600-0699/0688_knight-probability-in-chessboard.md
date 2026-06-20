# 0688 — Knight Probability In Chessboard

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid). Tugasmu menjelajahi atau memanipulasi grid.

**Cara berpikir:** `grid[row][col]`. 4 arah: atas/bawah/kiri/kanan. Selalu cek boundary.

**Fungsi Solusi:** `func knightProbability(n int, k int, row int, column int) float64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(K * N^2)  |  **Ruang:** O(N^2)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #688: Knight Probability in Chessboard
// https://leetcode.com/problems/knight-probability-in-chessboard/
// Difficulty: Medium
// Time: O(K * N^2)
// Space: O(N^2)

import "fmt"

func main() {
	fmt.Println(knightProbability(3, 2, 0, 0))
	fmt.Println(knightProbability(1, 0, 0, 0))
}

func knightProbability(n int, k int, row int, column int) float64 {
	dirs := [][2]int{{-2, -1}, {-2, 1}, {-1, -2}, {-1, 2}, {1, -2}, {1, 2}, {2, -1}, {2, 1}}
  // Matriks 2D
	dp := make([][]float64, n)
  // Range loop
	for i := range dp {
		dp[i] = make([]float64, n)
	}
	dp[row][column] = 1.0

	for step := 0; step < k; step++ {
  // Matriks 2D
		next := make([][]float64, n)
  // Range loop
		for i := range next {
			next[i] = make([]float64, n)
		}
		for r := 0; r < n; r++ {
			for c := 0; c < n; c++ {
				if dp[r][c] == 0 {
					continue
				}
				for _, d := range dirs {
					nr, nc := r+d[0], c+d[1]
					if nr >= 0 && nr < n && nc >= 0 && nc < n {
						next[nr][nc] += dp[r][c] / 8.0
					}
				}
			}
		}
		dp = next
	}

	result := 0.0
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			result += dp[r][c]
		}
	}
	return result
}
```
