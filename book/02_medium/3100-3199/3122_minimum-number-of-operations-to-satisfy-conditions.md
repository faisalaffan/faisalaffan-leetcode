# 3122 — Minimum Number Of Operations To Satisfy Conditions

## Deskripsi

**Soal:** [3122. Minimum Number Of Operations To Satisfy Conditions](https://leetcode.com/problems/minimum-number-of-operations-to-satisfy-conditions/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n * m * 10)  
**Kompleksitas Ruang:** O(m * 10)

**Algoritma:** —

**Fungsi Solusi:** `func minimumOperations(grid [][]int) int`

## Solusi Go

```go
package main

// LeetCode #3122: Minimum Number of Operations to Satisfy Conditions
// https://leetcode.com/problems/minimum-number-of-operations-to-satisfy-conditions/
// Difficulty: Medium
// Time: O(n * m * 10) | Space: O(m * 10)

import (
	"fmt"
	"math"
)

func minimumOperations(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

  // Membuat slice untuk menyimpan hasil
	cost := make([][10]int, n)
	for j := 0; j < n; j++ {
		for d := 0; d < 10; d++ {
			cnt := 0
			for i := 0; i < m; i++ {
				if grid[i][j] != d {
					cnt++
				}
			}
			cost[j][d] = cnt
		}
	}

  // Membuat slice untuk menyimpan hasil
	dp := make([][10]int, n)
	for j := 0; j < n; j++ {
		for d := 0; d < 10; d++ {
			dp[j][d] = math.MaxInt32
		}
	}

	for d := 0; d < 10; d++ {
		dp[0][d] = cost[0][d]
	}

	for j := 1; j < n; j++ {
		for d := 0; d < 10; d++ {
			for pd := 0; pd < 10; pd++ {
				if pd != d {
					dp[j][d] = min(dp[j][d], dp[j-1][pd]+cost[j][d])
				}
			}
		}
	}

	ans := math.MaxInt32
	for d := 0; d < 10; d++ {
		ans = min(ans, dp[n-1][d])
	}
	return ans
}

func main() {
	fmt.Println(minimumOperations([][]int{{1, 0, 2}, {1, 0, 2}})) // Expected: 0
	fmt.Println(minimumOperations([][]int{{1, 1, 1}, {0, 0, 0}})) // Expected: 3
}
```
