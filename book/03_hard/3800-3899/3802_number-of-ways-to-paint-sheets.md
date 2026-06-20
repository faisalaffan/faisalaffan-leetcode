# 3802 — Number Of Ways To Paint Sheets

## Deskripsi

**Soal:** [3802. Number Of Ways To Paint Sheets](https://leetcode.com/problems/number-of-ways-to-paint-sheets/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

> **Ide Kunci:** Combinatorial DP. dp[i][c] = ways to paint i sheets

## Solusi Go

```go
package main

// LeetCode #3802: Number of Ways to Paint Sheets [Paid]
// https://leetcode.com/problems/number-of-ways-to-paint-sheets/
// Difficulty: Hard
//
// Count ways to paint n sheets with given color limits per color.
//
// Approach: Combinatorial DP. dp[i][c] = ways to paint i sheets
// with first c colors satisfying limits.

import "fmt"

func main() {
	// Example 1
	fmt.Println(numberOfWays(3, []int{2, 2}))
	// Example 2
	fmt.Println(numberOfWays(5, []int{3, 2, 1}))
	// Edge: single color
	fmt.Println(numberOfWays(2, []int{5}))
	// Edge: impossible
	fmt.Println(numberOfWays(3, []int{1, 1}))
}

func numberOfWays(n int, limit []int) int {
	const mod = 1000000007
	m := len(limit)

	// dp[i][j] = ways to paint i sheets using first j colors
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, n+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	dp[0][0] = 1

	for j := 1; j <= m; j++ {
		maxUse := limit[j-1]
		for i := 0; i <= n; i++ {
			// Use k sheets of color j
			for k := 0; k <= maxUse && k <= i; k++ {
				dp[i][j] = (dp[i][j] + dp[i-k][j-1]) % mod
			}
		}
	}

	return dp[n][m]
}
```
