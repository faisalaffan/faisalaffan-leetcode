# 1420 — Build Array Where You Can Find The Maximum Exactly K Comparisons

## Deskripsi

**Soal:** [1420. Build Array Where You Can Find The Maximum Exactly K Comparisons](https://leetcode.com/problems/build-array-where-you-can-find-the-maximum-exactly-k-comparisons/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func numOfArrays(n int, m int, k int) int`

## Solusi Go

```go
package main

// LeetCode #1420: Build Array Where You Can Find The Maximum Exactly K Comparisons
// https://leetcode.com/problems/build-array-where-you-can-find-the-maximum-exactly-k-comparisons/
// Difficulty: Hard

import "fmt"

const mod1420 = 1_000_000_007

func numOfArrays(n int, m int, k int) int {
	if k == 0 || k > m {
		return 0
	}
	// dp[i][j][c] = ways for length i, max = j, cost = c
  // Membuat slice 2D untuk DP/tabel
	dp := make([][][]int, n+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}

	for j := 1; j <= m; j++ {
		dp[1][j][1] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for c := 1; c <= k; c++ {
				// Append value <= j: choose any of j values, cost unchanged
				dp[i][j][c] = (dp[i][j][c] + dp[i-1][j][c]*j) % mod1420

				// Append value == j (new max): sum over previous max < j
				if c > 1 {
					for p := 1; p < j; p++ {
						dp[i][j][c] = (dp[i][j][c] + dp[i-1][p][c-1]) % mod1420
					}
				}
			}
		}
	}

	var ans int
	for j := 1; j <= m; j++ {
		ans = (ans + dp[n][j][k]) % mod1420
	}
	return ans
}

func main() {
	// Example: n=2, m=3, k=1 -> 6
	fmt.Println(numOfArrays(2, 3, 1))
}
```
