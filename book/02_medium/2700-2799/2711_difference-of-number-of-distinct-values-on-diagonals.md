# 2711 — Difference Of Number Of Distinct Values On Diagonals

## Deskripsi

**Soal:** [2711. Difference Of Number Of Distinct Values On Diagonals](https://leetcode.com/problems/difference-of-number-of-distinct-values-on-diagonals/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n*(m+n))  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func differenceOfDistinctValues(grid [][]int) [][]int`

## Solusi Go

```go
package main

// LeetCode #2711: Difference of Number of Distinct Values on Diagonals
// https://leetcode.com/problems/difference-of-number-of-distinct-values-on-diagonals/
// Difficulty: Medium
// Time: O(m*n*(m+n)) | Space: O(1)

import "fmt"

func differenceOfDistinctValues(grid [][]int) [][]int {
	m, n := len(grid), len(grid[0])
  // Membuat slice 2D untuk DP/tabel
	ans := make([][]int, m)
  // Iterasi seluruh elemen
	for i := range ans {
		ans[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			// Count distinct values above-left diagonal
  // Membuat map untuk pencarian O(1): key → value
			aboveLeft := make(map[int]bool)
			r, c := i-1, j-1
			for r >= 0 && c >= 0 {
				aboveLeft[grid[r][c]] = true
				r--
				c--
			}

			// Count distinct values below-right diagonal
  // Membuat map untuk pencarian O(1): key → value
			belowRight := make(map[int]bool)
			r, c = i+1, j+1
			for r < m && c < n {
				belowRight[grid[r][c]] = true
				r++
				c++
			}

			diff := len(aboveLeft) - len(belowRight)
			if diff < 0 {
				diff = -diff
			}
			ans[i][j] = diff
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", differenceOfDistinctValues([][]int{{1, 2, 3}, {3, 1, 5}, {3, 2, 1}}))
	// Expected: [[1,1,0],[1,0,1],[0,1,1]]

	// Test case 2
	fmt.Println("Test 2:", differenceOfDistinctValues([][]int{{1}}))
	// Expected: [[0]]

	// Test case 3
	fmt.Println("Test 3:", differenceOfDistinctValues([][]int{{1, 2}, {3, 4}}))
	// Expected: [[1,0],[0,1]]
}
```
