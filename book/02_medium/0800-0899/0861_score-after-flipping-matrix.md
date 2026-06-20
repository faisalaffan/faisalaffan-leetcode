# 0861 — Score After Flipping Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func ScoreAfterFlippingMatrix(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #861: Score After Flipping Matrix
// https://leetcode.com/problems/score-after-flipping-matrix/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(ScoreAfterFlippingMatrix([][]int{{0, 0, 1, 1}, {1, 0, 1, 0}, {1, 1, 0, 0}}))
	fmt.Println(ScoreAfterFlippingMatrix([][]int{{0}}))
	fmt.Println(ScoreAfterFlippingMatrix([][]int{{1, 1}, {1, 1}}))
}

// Time: O(m * n) | Space: O(1)
func ScoreAfterFlippingMatrix(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	// Ensure first column is all 1s
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			for j := 0; j < n; j++ {
				grid[i][j] ^= 1
			}
		}
	}

	ans := 0
	for j := 0; j < n; j++ {
		ones := 0
		for i := 0; i < m; i++ {
			ones += grid[i][j]
		}
		if ones < m-ones {
			ones = m - ones
		}
		ans += ones * (1 << (n - 1 - j))
	}

	return ans
}
```
