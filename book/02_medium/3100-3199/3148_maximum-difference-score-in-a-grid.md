# 3148 — Maximum Difference Score In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxScore(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3148: Maximum Difference Score in a Grid
// https://leetcode.com/problems/maximum-difference-score-in-a-grid/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m * n)

import (
	"fmt"
	"math"
)

func maxScore(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

  // Membuat matriks/slice 2D untuk DP
	minVal := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range minVal {
		minVal[i] = make([]int, n)
	}

	ans := math.MinInt32

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prevMin := math.MaxInt32
			if i > 0 {
				prevMin = min(prevMin, minVal[i-1][j])
			}
			if j > 0 {
				prevMin = min(prevMin, minVal[i][j-1])
			}

			if i > 0 || j > 0 {
				ans = max(ans, grid[i][j]-prevMin)
			}

			minVal[i][j] = grid[i][j]
			if i > 0 {
				minVal[i][j] = min(minVal[i][j], minVal[i-1][j])
			}
			if j > 0 {
				minVal[i][j] = min(minVal[i][j], minVal[i][j-1])
			}
		}
	}

	return ans
}

func main() {
	fmt.Println(maxScore([][]int{{9, 5, 7, 3}, {8, 9, 6, 1}, {6, 7, 14, 3}, {2, 5, 3, 1}})) // Expected: 9
	fmt.Println(maxScore([][]int{{4, 3, 2}, {3, 2, 1}}))                                      // Expected: -1
}
```
