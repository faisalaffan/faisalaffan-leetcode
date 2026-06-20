# 2684 — Maximum Number Of Moves In A Grid

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxMoves(grid [][]int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2684: Maximum Number of Moves in a Grid
// https://leetcode.com/problems/maximum-number-of-moves-in-a-grid/
// Difficulty: Medium
// Time: O(m*n) | Space: O(m*n)

import "fmt"

func maxMoves(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	// dp[row][col] = max moves from (row, col)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n)
	}

	ans := 0
	for col := n - 2; col >= 0; col-- {
		for row := 0; row < m; row++ {
			val := grid[row][col]
			maxNext := 0
			// Can move to row-1, col+1
			if row > 0 && grid[row-1][col+1] > val {
				if 1+dp[row-1][col+1] > maxNext {
					maxNext = 1 + dp[row-1][col+1]
				}
			}
			// Can move to row, col+1
			if grid[row][col+1] > val {
				if 1+dp[row][col+1] > maxNext {
					maxNext = 1 + dp[row][col+1]
				}
			}
			// Can move to row+1, col+1
			if row < m-1 && grid[row+1][col+1] > val {
				if 1+dp[row+1][col+1] > maxNext {
					maxNext = 1 + dp[row+1][col+1]
				}
			}
			dp[row][col] = maxNext
			if col == 0 && maxNext > ans {
				ans = maxNext
			}
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxMoves([][]int{{2, 4, 3, 5}, {5, 4, 9, 3}, {3, 4, 2, 11}, {10, 9, 13, 15}}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", maxMoves([][]int{{3, 2, 4}, {2, 1, 9}, {1, 1, 7}}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", maxMoves([][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}))
	// Expected: 2
}
```
