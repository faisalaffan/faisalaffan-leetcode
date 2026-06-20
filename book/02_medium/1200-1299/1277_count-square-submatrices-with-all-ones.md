# 1277 — Count Square Submatrices With All Ones

## Deskripsi

**Soal:** [1277. Count Square Submatrices With All Ones](https://leetcode.com/problems/count-square-submatrices-with-all-ones/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func countSquares(matrix [][]int) int`

## Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1277: Count Square Submatrices with All Ones
// https://leetcode.com/problems/count-square-submatrices-with-all-ones/
// Difficulty: Medium

// dp[i][j] = side length of largest square ending at (i,j).
// If grid[i][j]==1: dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
// Sum all dp[i][j] = total squares.

// Time: O(m*n)
// Space: O(m*n)

func countSquares(matrix [][]int) int {
	m, n := len(matrix), len(matrix[0])
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, m)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n)
	}

	total := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 1 {
				if i == 0 || j == 0 {
					dp[i][j] = 1
				} else {
					dp[i][j] = 1 + min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1]))
				}
				total += dp[i][j]
			}
		}
	}

	return total
}

func main() {
	fmt.Printf("%d (expected: 15)\n",
		countSquares([][]int{
			{0, 1, 1, 1},
			{1, 1, 1, 1},
			{0, 1, 1, 1},
		}))

	fmt.Printf("%d (expected: 7)\n",
		countSquares([][]int{
			{1, 0, 1},
			{1, 1, 0},
			{1, 1, 0},
		}))
}
```
