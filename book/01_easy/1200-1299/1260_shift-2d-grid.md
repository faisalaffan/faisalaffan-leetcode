# 1260 — Shift 2D Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func shiftGrid(grid [][]int, k int) [][]int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #1260: Shift 2D Grid
// https://leetcode.com/problems/shift-2d-grid/
// Difficulty: Easy
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func main() {
	fmt.Println(shiftGrid([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1))
	// [[9,1,2],[3,4,5],[6,7,8]]
	fmt.Println(shiftGrid([][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4))
	// [[12,0,21,13],[3,8,1,9],[19,7,2,5],[4,6,11,10]]
}

// LeetCode submission: shiftGrid
func shiftGrid(grid [][]int, k int) [][]int {
	m, n := len(grid), len(grid[0])
  // Membuat matriks/slice 2D untuk DP
	ans := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range ans {
		ans[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			idx := (i*n + j + k) % (m * n)
			ni, nj := idx/n, idx%n
			ans[ni][nj] = grid[i][j]
		}
	}
	return ans
}
```
