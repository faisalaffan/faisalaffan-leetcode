# 0516 — Longest Palindromic Subsequence

## Deskripsi

**Soal:** [0516. Longest Palindromic Subsequence](https://leetcode.com/problems/longest-palindromic-subsequence/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(n^2)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

```go
package main

// LeetCode #516: Longest Palindromic Subsequence
// https://leetcode.com/problems/longest-palindromic-subsequence/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(n^2)

import "fmt"

func main() {
	fmt.Println(LongestPalindromicSubsequence("bbbab"))
	fmt.Println(LongestPalindromicSubsequence("cbbd"))
}

func LongestPalindromicSubsequence(s string) int {
	n := len(s)
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, n)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n)
		dp[i][i] = 1
	}

	for length := 2; length <= n; length++ {
		for i := 0; i <= n-length; i++ {
			j := i + length - 1
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1] + 2
			} else {
				if dp[i+1][j] > dp[i][j-1] {
					dp[i][j] = dp[i+1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[0][n-1]
}
```
