# 2435 — Paths In Matrix Whose Sum Is Divisible By K

## Deskripsi

**Soal:** [2435. Paths In Matrix Whose Sum Is Divisible By K](https://leetcode.com/problems/paths-in-matrix-whose-sum-is-divisible-by-k/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func numberOfPaths(grid [][]int, k int) int`

> **Ide Kunci:** 3D DP.

## Solusi Go

```go
package main

// LeetCode #2435: Paths in Matrix Whose Sum Is Divisible by K
// https://leetcode.com/problems/paths-in-matrix-whose-sum-is-divisible-by-k/
// Difficulty: Hard
//
// Given a grid of size m x n and an integer k, count the number of paths
// from (0,0) to (m-1,n-1) moving only right or down, such that the sum of
// values along the path is divisible by k.
//
// Approach: 3D DP.
// dp[i][j][mod] = number of ways to reach (i,j) with sum % k == mod.
// Transition from top (i-1,j) and left (i,j-1).
// mod = (prevMod + grid[i][j]) % k.

import "fmt"

func numberOfPaths(grid [][]int, k int) int {
	const mod = 1_000_000_007
	m := len(grid)
	n := len(grid[0])

	// dp[i][j][r] = ways to reach (i,j) with sum % k == r
  // Membuat slice 2D untuk DP/tabel
	dp := make([][][]int, m)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([][]int, n)
		for j := range dp[i] {
			dp[i][j] = make([]int, k)
		}
	}

	// Initialize start position
	dp[0][0][grid[0][0]%k] = 1

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			val := grid[i][j]
			for r := 0; r < k; r++ {
				ways := 0
				if i > 0 {
					ways = (ways + dp[i-1][j][r]) % mod
				}
				if j > 0 {
					ways = (ways + dp[i][j-1][r]) % mod
				}
				newR := (r + val) % k
				dp[i][j][newR] = (dp[i][j][newR] + ways) % mod
			}
		}
	}

	return dp[m-1][n-1][0]
}

func main() {
	// Example 1
	fmt.Println(numberOfPaths([][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}}, 3))
	// Example 2
	fmt.Println(numberOfPaths([][]int{{0, 0}}, 5))
	// Example 3
	fmt.Println(numberOfPaths([][]int{{7, 3, 4, 9}, {2, 7, 6, 0}}, 1))
	// Single cell
	fmt.Println(numberOfPaths([][]int{{10}}, 2))
	// 1x3
	fmt.Println(numberOfPaths([][]int{{1, 2, 3}}, 3))
}
```
