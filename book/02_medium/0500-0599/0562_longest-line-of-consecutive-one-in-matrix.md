# 0562 — Longest Line Of Consecutive One In Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func LongestLine(mat [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #562: Longest Line of Consecutive One in Matrix
// https://leetcode.com/problems/longest-line-of-consecutive-one-in-matrix/
// Difficulty: Medium [Paid]
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	mat := [][]int{
		{0, 1, 1, 0},
		{0, 1, 1, 0},
		{0, 0, 0, 1},
	}
	fmt.Println(LongestLine(mat))
}

func LongestLine(mat [][]int) int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return 0
	}
	m, n := len(mat), len(mat[0])
	maxLen := 0

	// dp[i][j][0] = horizontal, dp[i][j][1] = vertical
	// dp[i][j][2] = diagonal, dp[i][j][3] = anti-diagonal
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, 4)
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 1 {
				dp[i][j][0] = 1
				dp[i][j][1] = 1
				dp[i][j][2] = 1
				dp[i][j][3] = 1

				if j > 0 {
					dp[i][j][0] = dp[i][j-1][0] + 1
				}
				if i > 0 {
					dp[i][j][1] = dp[i-1][j][1] + 1
				}
				if i > 0 && j > 0 {
					dp[i][j][2] = dp[i-1][j-1][2] + 1
				}
				if i > 0 && j < n-1 {
					dp[i][j][3] = dp[i-1][j+1][3] + 1
				}

				for k := 0; k < 4; k++ {
					if dp[i][j][k] > maxLen {
						maxLen = dp[i][j][k]
					}
				}
			}
		}
	}

	return maxLen
}
```
