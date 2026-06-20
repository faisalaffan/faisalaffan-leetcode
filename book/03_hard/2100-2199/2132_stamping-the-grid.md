# 2132 — Stamping The Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2132: Stamping the Grid
// https://leetcode.com/problems/stamping-the-grid/
// Difficulty: Hard
//
// 2D prefix sum + 2D difference array. First compute prefix sums to query
// empty subgrids. For each possible stamp top-left, if the area is obstacle-free,
// mark it in the diff array. Then reconstruct coverage and verify all empty
// cells are covered by at least one stamp.

import "fmt"

func main() {
	grid1 := [][]int{
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
		{1, 0, 0, 0},
	}
	fmt.Println(possibleToStamp(grid1, 4, 3))

	grid2 := [][]int{
		{1, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 1},
	}
	fmt.Println(possibleToStamp(grid2, 2, 2))

	fmt.Println(possibleToStamp([][]int{{0}}, 1, 1))
}

func possibleToStamp(grid [][]int, stampHeight int, stampWidth int) bool {
	m, n := len(grid), len(grid[0])

	// 2D prefix sum (obstacles = 1)
  // Membuat matriks/slice 2D untuk DP
	prefix := make([][]int, m+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range prefix {
		prefix[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			prefix[i+1][j+1] = prefix[i][j+1] + prefix[i+1][j] - prefix[i][j] + grid[i][j]
		}
	}

	sumRange := func(r1, c1, r2, c2 int) int {
		return prefix[r2+1][c2+1] - prefix[r1][c2+1] - prefix[r2+1][c1] + prefix[r1][c1]
	}

	// 2D difference array
  // Membuat matriks/slice 2D untuk DP
	diff := make([][]int, m+2)
  // Range loop: iterasi dengan indeks + nilai
	for i := range diff {
		diff[i] = make([]int, n+2)
	}

	// Mark valid stamp placements
	for i := 0; i+stampHeight <= m; i++ {
		for j := 0; j+stampWidth <= n; j++ {
			if sumRange(i, j, i+stampHeight-1, j+stampWidth-1) == 0 {
				diff[i][j]++
				diff[i][j+stampWidth]--
				diff[i+stampHeight][j]--
				diff[i+stampHeight][j+stampWidth]++
			}
		}
	}

	// Reconstruct coverage and verify
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				diff[i][j] += diff[i-1][j]
			}
			if j > 0 {
				diff[i][j] += diff[i][j-1]
			}
			if i > 0 && j > 0 {
				diff[i][j] -= diff[i-1][j-1]
			}
			if grid[i][j] == 0 && diff[i][j] == 0 {
				return false
			}
		}
	}

	return true
}
```
