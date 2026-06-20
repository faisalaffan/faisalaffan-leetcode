# 3567 — Minimum Absolute Difference In Sliding Submatrix

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan matriks 2D (grid) — array dua dimensi dengan baris dan kolom. Tugasmu adalah menjelajahi, memanipulasi, atau menghitung properti matriks tersebut.

Bayangkan spreadsheet Excel: ada baris (row) dan kolom (column). Setiap sel punya nilai. Kamu perlu mengolah data di dalam grid tersebut. Matriks di Go adalah `[][]int` (slice of slice).

**Konsep kunci:** baris (row), kolom (col), boundary check, arah gerak (atas/bawah/kiri/kanan), prefix sum 2D.

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumAbsoluteDifferenceInSlidingSubmatrix(grid [][]int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Sliding Window

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Sliding Window** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3567: Minimum Absolute Difference in Sliding Submatrix
// https://leetcode.com/problems/minimum-absolute-difference-in-sliding-submatrix/
// Difficulty: Medium
// Complexity: O(n*m*k^2) time, O(k^2) space

import "fmt"

func main() {
	// Test case 1
	grid := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	k := 2
	fmt.Println("Test 1:", MinimumAbsoluteDifferenceInSlidingSubmatrix(grid, k))
	// Test case 2
	grid2 := [][]int{{1, 1}, {1, 1}}
	k2 := 2
	fmt.Println("Test 2:", MinimumAbsoluteDifferenceInSlidingSubmatrix(grid2, k2))
	// Test case 3
	grid3 := [][]int{{5}}
	k3 := 1
	fmt.Println("Test 3:", MinimumAbsoluteDifferenceInSlidingSubmatrix(grid3, k3))
}

func MinimumAbsoluteDifferenceInSlidingSubmatrix(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	minDiff := -1
	for i := 0; i <= m-k; i++ {
		for j := 0; j <= n-k; j++ {
			// Find min and max in submatrix
			minVal, maxVal := grid[i][j], grid[i][j]
			for r := i; r < i+k; r++ {
				for c := j; c < j+k; c++ {
					if grid[r][c] < minVal {
						minVal = grid[r][c]
					}
					if grid[r][c] > maxVal {
						maxVal = grid[r][c]
					}
				}
			}
			diff := maxVal - minVal
			if minDiff == -1 || diff < minDiff {
				minDiff = diff
			}
		}
	}
	return minDiff
}
```
