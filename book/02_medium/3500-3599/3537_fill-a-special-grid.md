# 3537 — Fill A Special Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func FillASpecialGrid(grid [][]int) 
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #3537: Fill a Special Grid
// https://leetcode.com/problems/fill-a-special-grid/
// Difficulty: Medium
// Complexity: O(n*m) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	grid := [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	FillASpecialGrid(grid)
	fmt.Println("Test 1:", grid)
	// Test case 2
	grid2 := [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}}
	FillASpecialGrid(grid2)
	fmt.Println("Test 2:", grid2)
	// Test case 3
	grid3 := [][]int{{0}}
	FillASpecialGrid(grid3)
	fmt.Println("Test 3:", grid3)
}

func FillASpecialGrid(grid [][]int) {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return
	}
	m, n := len(grid), len(grid[0])
	// Fill each cell with the sum of its row and column index
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				grid[i][j] = i + j
			}
		}
	}
}
```
