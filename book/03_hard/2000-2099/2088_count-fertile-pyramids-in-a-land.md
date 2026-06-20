# 2088 — Count Fertile Pyramids In A Land

## Deskripsi

**Soal:** [2088. Count Fertile Pyramids In A Land](https://leetcode.com/problems/count-fertile-pyramids-in-a-land/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #2088: Count Fertile Pyramids in a Land
// https://leetcode.com/problems/count-fertile-pyramids-in-a-land/
// Difficulty: Hard
//
// DP approach: dp[i][j] = max pyramid height with (i,j) as the top.
// Regular pyramid (top down): dp[i][j] = 1 + min(dp[i+1][j-1], dp[i+1][j], dp[i+1][j+1])
// Inverted pyramid (top up):  dp[i][j] = 1 + min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1])
// Sum (dp[i][j] - 1) over all cells where dp[i][j] > 1.

import "fmt"

func main() {
	fmt.Println(countPyramids([][]int{{0, 1, 1, 0}, {1, 1, 1, 1}}))
	fmt.Println(countPyramids([][]int{{1, 1, 1}, {1, 1, 1}}))
	fmt.Println(countPyramids([][]int{{1}}))
	fmt.Println(countPyramids([][]int{{1, 1}}))
	fmt.Println(countPyramids([][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}}))
}

func countPyramids(grid [][]int) int {
	m, n := len(grid), len(grid[0])
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, m)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n)
	}

	total := 0

	// Regular pyramids (top pointing down): bottom-up DP
	for i := m - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				dp[i][j] = 0
				continue
			}
			if i == m-1 || j == 0 || j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 1 + min(dp[i+1][j-1], min(dp[i+1][j], dp[i+1][j+1]))
			}
			if dp[i][j] > 1 {
				total += dp[i][j] - 1
			}
		}
	}

	// Inverted pyramids (top pointing up): top-down DP
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				dp[i][j] = 0
				continue
			}
			if i == 0 || j == 0 || j == n-1 {
				dp[i][j] = 1
			} else {
				dp[i][j] = 1 + min(dp[i-1][j-1], min(dp[i-1][j], dp[i-1][j+1]))
			}
			if dp[i][j] > 1 {
				total += dp[i][j] - 1
			}
		}
	}

	return total
}
```
