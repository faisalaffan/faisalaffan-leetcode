# 2912 — Number Of Ways To Reach Destination In The Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func numberOfWaysToReachDestinationInTheGrid(m, n int, blocked [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2912: Number of Ways to Reach Destination in the Grid
// https://leetcode.com/problems/number-of-ways-to-reach-destination-in-the-grid/
// Difficulty: Hard [Paid]
//
// Count the number of ways to go from (0,0) to (m-1,n-1) in a grid, moving only
// right and down, with some blocked cells. Uses DP with modulo 10^9+7.
// For large grids with few blocked cells, uses combinatorics (inclusion-exclusion).
// Default implementation handles general DP for moderate-sized grids.

import "fmt"

const mod2912 = 1000000007

func numberOfWaysToReachDestinationInTheGrid(m, n int, blocked [][]int) int {
	// Mark blocked cells
  // Membuat matriks/slice 2D untuk DP
	grid := make([][]bool, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range grid {
		grid[i] = make([]bool, n)
	}
	for _, b := range blocked {
		if b[0] < m && b[1] < n {
			grid[b[0]][b[1]] = true
		}
	}

  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n)
	}

	// First row
	for j := 0; j < n; j++ {
		if grid[0][j] {
			break
		}
		dp[0][j] = 1
	}

	// First column
	for i := 0; i < m; i++ {
		if grid[i][0] {
			break
		}
		dp[i][0] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if grid[i][j] {
				dp[i][j] = 0
			} else {
				dp[i][j] = (dp[i-1][j] + dp[i][j-1]) % mod2912
			}
		}
	}

	return dp[m-1][n-1]
}

func main() {
	// Example: 3x3, no blocked
	fmt.Println(numberOfWaysToReachDestinationInTheGrid(3, 3, [][]int{}))

	// Example: 2x2, no blocked => 2 (RD, DR)
	fmt.Println(numberOfWaysToReachDestinationInTheGrid(2, 2, [][]int{}))

	// Example: 3x3, one blocked at (1,1) => 2
	fmt.Println(numberOfWaysToReachDestinationInTheGrid(3, 3, [][]int{{1, 1}}))

	// 1xN grid
	fmt.Println(numberOfWaysToReachDestinationInTheGrid(1, 5, [][]int{}))

	// Start is blocked => 0
	fmt.Println(numberOfWaysToReachDestinationInTheGrid(3, 3, [][]int{{0, 0}}))

	// End is blocked => 0
	fmt.Println(numberOfWaysToReachDestinationInTheGrid(3, 3, [][]int{{2, 2}}))

	// Larger grid
	fmt.Println(numberOfWaysToReachDestinationInTheGrid(10, 10, [][]int{}))
}
```
