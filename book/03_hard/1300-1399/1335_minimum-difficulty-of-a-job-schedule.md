# 1335 — Minimum Difficulty Of A Job Schedule

## Deskripsi

**Soal:** [1335. Minimum Difficulty Of A Job Schedule](https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func minDifficulty(jobDifficulty []int, d int) int`

## Solusi Go

```go
package main

// LeetCode #1335: Minimum Difficulty of a Job Schedule
// https://leetcode.com/problems/minimum-difficulty-of-a-job-schedule/
// Difficulty: Hard

import "fmt"

func minDifficulty(jobDifficulty []int, d int) int {
	n := len(jobDifficulty)
	if n < d {
		return -1
	}

	// dp[i][j] = min difficulty to schedule first j jobs in i days
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, d+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n+1)
		for j := range dp[i] {
			dp[i][j] = 1 << 30 // large number
		}
	}
	dp[0][0] = 0

	for i := 1; i <= d; i++ {
		for j := i; j <= n; j++ {
			maxVal := 0
			// k = first job done on day i, so jobs k..j-1 are done on day i
			for k := j; k >= i; k-- {
				if jobDifficulty[k-1] > maxVal {
					maxVal = jobDifficulty[k-1]
				}
				if dp[i-1][k-1]+maxVal < dp[i][j] {
					dp[i][j] = dp[i-1][k-1] + maxVal
				}
			}
		}
	}

	return dp[d][n]
}

func main() {
	// Example 1
	fmt.Println(minDifficulty([]int{6, 5, 4, 3, 2, 1}, 2))
	// Expected: 7

	// Example 2
	fmt.Println(minDifficulty([]int{9, 9, 9}, 4))
	// Expected: -1

	// Example 3
	fmt.Println(minDifficulty([]int{1, 1, 1}, 3))
	// Expected: 3
}
```
