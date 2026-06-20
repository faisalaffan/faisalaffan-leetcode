# 0062 — Unique Paths

## Deskripsi

**Soal:** [0062. Unique Paths](https://leetcode.com/problems/unique-paths/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(m*n)  
**Kompleksitas Ruang:** O(m*n)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func uniquePaths(m int, n int) int`

## Solusi Go

```go
package main

// LeetCode #62: Unique Paths
// https://leetcode.com/problems/unique-paths/
// Difficulty: Medium

import "fmt"

func uniquePaths(m int, n int) int {
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, m)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][0] = 1
	}
	for j := 0; j < n; j++ {
		dp[0][j] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1]
}

func main() {
	// Test case 1
	fmt.Println(uniquePaths(3, 7)) // 28

	// Test case 2
	fmt.Println(uniquePaths(3, 2)) // 3

	// Test case 3
	fmt.Println(uniquePaths(7, 3)) // 28
}

// Time: O(m*n) | Space: O(m*n)
```
