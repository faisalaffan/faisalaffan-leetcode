# 1594 — Maximum Non Negative Product In A Matrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func MaxProductPath(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(R*C), Space: O(C)  
**Kompleksitas Ruang:** O(C)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1594: Maximum Non Negative Product in a Matrix
// https://leetcode.com/problems/maximum-non-negative-product-in-a-matrix/
// Difficulty: Medium

import "fmt"

func main() {
	grid1 := [][]int{{-1, -2, -3}, {-2, -3, -3}, {-3, -3, -2}}
	fmt.Println(MaxProductPath(grid1))

	grid2 := [][]int{{1, -2, 1}, {1, -2, 1}, {3, -4, 1}}
	fmt.Println(MaxProductPath(grid2))

	grid3 := [][]int{{1, 3}, {0, -4}}
	fmt.Println(MaxProductPath(grid3))
}

func MaxProductPath(grid [][]int) int {
	// Time: O(R*C), Space: O(C)
	const mod = 1_000_000_007

	rows, cols := len(grid), len(grid[0])
	if rows == 0 || cols == 0 {
		return -1
	}

	// minDP[r][c] = minimum product to reach (r,c)
	// maxDP[r][c] = maximum product to reach (r,c)
  // Membuat matriks/slice 2D untuk DP
	minDP := make([][]int64, rows)
  // Membuat matriks/slice 2D untuk DP
	maxDP := make([][]int64, rows)
	for i := 0; i < rows; i++ {
		minDP[i] = make([]int64, cols)
		maxDP[i] = make([]int64, cols)
	}

	maxDP[0][0] = int64(grid[0][0])
	minDP[0][0] = int64(grid[0][0])

	// First row
	for c := 1; c < cols; c++ {
		val := int64(grid[0][c])
		maxDP[0][c] = maxDP[0][c-1] * val
		minDP[0][c] = maxDP[0][c-1] * val
	}

	// First column
	for r := 1; r < rows; r++ {
		val := int64(grid[r][0])
		maxDP[r][0] = maxDP[r-1][0] * val
		minDP[r][0] = maxDP[r-1][0] * val
	}

	for r := 1; r < rows; r++ {
		for c := 1; c < cols; c++ {
			val := int64(grid[r][c])

			options := []int64{
				maxDP[r-1][c] * val,
				maxDP[r][c-1] * val,
				minDP[r-1][c] * val,
				minDP[r][c-1] * val,
			}

			maxVal := options[0]
			minVal := options[0]
			for _, opt := range options {
				if opt > maxVal {
					maxVal = opt
				}
				if opt < minVal {
					minVal = opt
				}
			}

			maxDP[r][c] = maxVal
			minDP[r][c] = minVal
		}
	}

	if maxDP[rows-1][cols-1] < 0 {
		return -1
	}

	return int(maxDP[rows-1][cols-1] % mod)
}
```
