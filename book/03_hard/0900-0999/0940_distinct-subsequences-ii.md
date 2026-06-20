# 0940 — Distinct Subsequences Ii

## Deskripsi

**Soal:** [0940. Distinct Subsequences Ii](https://leetcode.com/problems/distinct-subsequences-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func distinctSubseqII(s string) int`

## Solusi Go

```go
package main

// LeetCode #940: Distinct Subsequences II
// https://leetcode.com/problems/distinct-subsequences-ii/
// Difficulty: Hard
// DP with last occurrence tracking.
// dp[i] = 2*dp[i-1] - dp[last[s[i-1]]-1] (mod 1e9+7)
// Result counts non-empty subsequences.

import "fmt"

func distinctSubseqII(s string) int {
	mod := int(1e9 + 7)
	n := len(s)
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n+1)
	dp[0] = 1 // empty subsequence

  // Membuat map untuk pencarian O(1): key → value
	last := make(map[byte]int) // last occurrence index (1-based)

	for i := 1; i <= n; i++ {
		dp[i] = (dp[i-1] * 2) % mod
		c := s[i-1]
		if prev, ok := last[c]; ok {
			dp[i] = (dp[i] - dp[prev-1] + mod) % mod
		}
		last[c] = i
	}

	// Subtract empty subsequence
	return (dp[n] - 1 + mod) % mod
}

func main() {
	fmt.Println(distinctSubseqII("abc")) // Expected: 7
	fmt.Println(distinctSubseqII("aba")) // Expected: 6
	fmt.Println(distinctSubseqII("aaa")) // Expected: 3
	fmt.Println(distinctSubseqII("ab"))  // Expected: 3
	fmt.Println(distinctSubseqII("abca")) // Expected: 14?
}
```
