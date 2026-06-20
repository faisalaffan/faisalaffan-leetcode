# 0115 — Distinct Subsequences

## Deskripsi

**Soal:** [0115. Distinct Subsequences](https://leetcode.com/problems/distinct-subsequences/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #115: Distinct Subsequences
// https://leetcode.com/problems/distinct-subsequences/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("115. Distinct Subsequences")
	fmt.Println("rabbbit, rabbit ->", numDistinct("rabbbit", "rabbit"), "(expected 3)")
	fmt.Println("babgbag, bag ->", numDistinct("babgbag", "bag"), "(expected 5)")
	fmt.Println("r, r ->", numDistinct("r", "r"), "(expected 1)")
}

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
  // Edge case: input kosong
	if n == 0 {
		return 1
	}
	if m < n {
		return 0
	}

	// dp[j] = number of distinct subsequences of s[:i] that equal t[:j]
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= m; i++ {
		prev := dp[0]
		for j := 1; j <= n; j++ {
			curr := dp[j]
			if s[i-1] == t[j-1] {
				dp[j] = prev + dp[j]
			}
			prev = curr
		}
	}

	return dp[n]
}
```
