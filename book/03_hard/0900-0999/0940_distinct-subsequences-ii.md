# 0940 — Distinct Subsequences Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func distinctSubseqII(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

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
  // Alokasi slice
	dp := make([]int, n+1)
	dp[0] = 1 // empty subsequence

  // HashMap: O(1) lookup
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
