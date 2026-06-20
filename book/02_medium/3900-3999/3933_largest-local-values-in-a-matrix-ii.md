# 3933 — Largest Local Values In A Matrix Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func LargestLocalValuesInAMatrixIi(grid [][]int) int
```

> **💡 Hint:** Precompute 2D prefix sums for each value threshold.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Prefix Sum

**Kompleksitas Waktu:** O(N*M*V)  
**Kompleksitas Ruang:** O(N*M*V) where V = max value <= 200

> **Untuk fresh graduate:** Kuasai dulu teknik **Prefix Sum** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3933: Largest Local Values in a Matrix II
// https://leetcode.com/problems/largest-local-values-in-a-matrix-ii/
// Difficulty: Medium
// Time: O(N*M*V) | Space: O(N*M*V) where V = max value <= 200
// Approach: Precompute 2D prefix sums for each value threshold.
// For each cell (i,j) with value v > 0, check if any cell in the
// (2v+1)x(2v+1) region (minus corners) has value > v.

import "fmt"

func LargestLocalValuesInAMatrixIi(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	// Find max value in grid
	maxVal := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] > maxVal {
				maxVal = grid[i][j]
			}
		}
	}

	// Precompute 2D prefix sums for each threshold t: count of cells >= t
	// pref[t][i+1][j+1] = count of cells with value >= t in rectangle [0..i]x[0..j]
  // Membuat matriks/slice 2D untuk DP
	pref := make([][][]int, maxVal+2)
	for t := 1; t <= maxVal+1; t++ {
  // Membuat matriks/slice 2D untuk DP
		p := make([][]int, m+1)
		for i := 0; i <= m; i++ {
			p[i] = make([]int, n+1)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				val := 0
				if grid[i][j] >= t {
					val = 1
				}
				p[i+1][j+1] = p[i][j+1] + p[i+1][j] - p[i][j] + val
			}
		}
		pref[t] = p
	}

	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			v := grid[i][j]
			if v == 0 {
				continue
			}

			// Define the check region: (2v+1)x(2v+1) minus corners
			r1 := max(0, i-v)
			c1 := max(0, j-v)
			r2 := min(m-1, i+v)
			c2 := min(n-1, j+v)

			// Count cells with value > v in the full rectangle
			p := pref[v+1]
			total := p[r2+1][c2+1] - p[r1][c2+1] - p[r2+1][c1] + p[r1][c1]

			// Subtract corners if they're within bounds and distinct from edges
			corners := [][2]int{
				{i - v, j - v},
				{i - v, j + v},
				{i + v, j - v},
				{i + v, j + v},
			}
			for _, c := range corners {
				cr, cc := c[0], c[1]
				if cr >= 0 && cr < m && cc >= 0 && cc < n {
					if grid[cr][cc] > v {
						total--
					}
				}
			}

			if total == 0 {
				ans++
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1
  // Membuat matriks/slice 2D untuk DP
	grid1 := make([][]int, 7)
	for i := 0; i < 7; i++ {
		grid1[i] = make([]int, 7)
	}
	grid1[3][3] = 2
	fmt.Println(LargestLocalValuesInAMatrixIi(grid1)) // Expected: 1

	// Example 2
	fmt.Println(LargestLocalValuesInAMatrixIi([][]int{{1, 2}, {3, 4}})) // Expected: 1

	// Example 3
	fmt.Println(LargestLocalValuesInAMatrixIi([][]int{{1, 0, 1}, {0, 1, 0}, {1, 0, 1}})) // Expected: 5

	// Example 4
	fmt.Println(LargestLocalValuesInAMatrixIi([][]int{{1, 1}, {1, 1}})) // Expected: 4
}
```
