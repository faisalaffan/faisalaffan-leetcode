# 3643 — Flip Square Submatrix Vertically

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func FlipSquareSubmatrixVertically(grid [][]int, x, y, k int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(k^2)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3643: Flip Square Submatrix Vertically
// https://leetcode.com/problems/flip-square-submatrix-vertically/
// Difficulty: Easy

import "fmt"

func main() {
	grid := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}
	fmt.Println(FlipSquareSubmatrixVertically(grid, 1, 1, 3))
}

// Time: O(k^2)
// Space: O(1)
func FlipSquareSubmatrixVertically(grid [][]int, x, y, k int) [][]int {
	for i := x; i < x+k/2; i++ {
		i2 := x + k - 1 - (i - x)
		for j := y; j < y+k; j++ {
			grid[i][j], grid[i2][j] = grid[i2][j], grid[i][j]
		}
	}
	return grid
}
```
