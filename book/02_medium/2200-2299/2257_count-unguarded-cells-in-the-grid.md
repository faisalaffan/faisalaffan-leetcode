# 2257 — Count Unguarded Cells In The Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func countUnguarded(m int, n int, guards [][]int, walls [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** — (analisis sendiri ☕)

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Coba pahami dulu input/output sebelum melihat kode. Gambar di kertas kalau perlu!

## 💻 Solusi Go

```go
package main

// LeetCode #2257: Count Unguarded Cells in the Grid
// https://leetcode.com/problems/count-unguarded-cells-in-the-grid/
// Difficulty: Medium
// Time: O(m * n) | Space: O(m * n)

import "fmt"

func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
  // Membuat matriks/slice 2D untuk DP
	grid := make([][]int, m)
	for i := 0; i < m; i++ {
		grid[i] = make([]int, n)
	}
	// 1 = wall, 2 = guard

	for _, w := range walls {
		grid[w[0]][w[1]] = 1
	}
	for _, g := range guards {
		grid[g[0]][g[1]] = 2
	}

	dirs := [][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, g := range guards {
		gr, gc := g[0], g[1]
		for _, d := range dirs {
			r, c := gr+d[0], gc+d[1]
			for r >= 0 && r < m && c >= 0 && c < n && grid[r][c] != 1 && grid[r][c] != 2 {
				if grid[r][c] == 0 {
					grid[r][c] = 3 // guarded
				}
				r += d[0]
				c += d[1]
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				count++
			}
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println(countUnguarded(4, 6, [][]int{{0, 0}, {1, 1}, {2, 3}}, [][]int{{0, 1}, {2, 2}, {1, 4}}))
	// Expected: 7

	// Test case 2
	fmt.Println(countUnguarded(3, 3, [][]int{{1, 1}}, [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}}))
	// Expected: 4
}
```
