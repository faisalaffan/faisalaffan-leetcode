# 2304 — Minimum Path Cost In A Grid

## Deskripsi

**Soal:** [2304. Minimum Path Cost In A Grid](https://leetcode.com/problems/minimum-path-cost-in-a-grid/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m * n^2)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func minPathCost(grid [][]int, moveCost [][]int) int`

## Solusi Go

```go
package main

// LeetCode #2304: Minimum Path Cost in a Grid
// https://leetcode.com/problems/minimum-path-cost-in-a-grid/
// Difficulty: Medium
// Time: O(m * n^2) | Space: O(n)

import "fmt"

func minPathCost(grid [][]int, moveCost [][]int) int {
	m, n := len(grid), len(grid[0])
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n)
	copy(dp, grid[0])

	for r := 1; r < m; r++ {
  // Membuat slice untuk menyimpan hasil
		next := make([]int, n)
		for j := 0; j < n; j++ {
			next[j] = 1 << 30
			for k := 0; k < n; k++ {
				cost := dp[k] + moveCost[grid[r-1][k]][j] + grid[r][j]
				if cost < next[j] {
					next[j] = cost
				}
			}
		}
		dp = next
	}

	minCost := dp[0]
	for _, v := range dp {
		if v < minCost {
			minCost = v
		}
	}
	return minCost
}

func main() {
	// Test case 1
	fmt.Println(minPathCost([][]int{{5, 3}, {4, 0}, {2, 1}}, [][]int{{9, 8}, {1, 5}, {10, 12}, {18, 6}, {2, 4}, {14, 3}}))
	// Expected: 17

	// Test case 2
	fmt.Println(minPathCost([][]int{{5, 1, 2}, {4, 0, 3}}, [][]int{{12, 10, 15}, {20, 23, 8}, {21, 7, 1}, {8, 1, 13}, {9, 10, 25}, {5, 3, 2}}))
	// Expected: 6
}
```
