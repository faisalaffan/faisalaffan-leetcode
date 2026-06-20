# 1977 — Number Of Ways To Separate Numbers

## Deskripsi

**Soal:** [1977. Number Of Ways To Separate Numbers](https://leetcode.com/problems/number-of-ways-to-separate-numbers/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func numberOfWaysToSeparateNumbers(s string) int`

> **Ide Kunci:** DP + LCP (Longest Common Prefix)

## Solusi Go

```go
package main

// LeetCode #1977: Number of Ways to Separate Numbers
// https://leetcode.com/problems/number-of-ways-to-separate-numbers/
// Difficulty: Hard
// Approach: DP + LCP (Longest Common Prefix)
// dp[i] = number of ways to split suffix s[i:]
// pref[i] = sum_{k >= i} dp[k]
// For each i, try each j > i. Compare s[i:j] with s[j:j+len] using LCP.
// If s[i:j] < s[j:j+len], contribution = pref[j+len] (all splits where next number length >= len)
// else contribution = pref[j+len+1] (all splits where next number length > len)

import "fmt"

const MOD1977 = 1000000007

func numberOfWaysToSeparateNumbers(s string) int {
	n := len(s)

	// LCP[i][j] = longest common prefix of s[i:] and s[j:]
  // Membuat slice 2D untuk DP/tabel
	lcp := make([][]int, n+1)
  // Iterasi seluruh elemen
	for i := range lcp {
		lcp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if s[i] == s[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}

  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n+2)  // dp[n] = 1 (empty suffix)
  // Membuat slice untuk menyimpan hasil
	pref := make([]int, n+2) // pref[i] = sum_{k >= i} dp[k]
	dp[n] = 1
	pref[n] = 1

	for i := n - 1; i >= 0; i-- {
		if s[i] == '0' {
			pref[i] = pref[i+1]
			continue
		}
		total := 0
		for j := i + 1; j <= n; j++ {
			curLen := j - i
			if j == n {
				total = (total + 1) % MOD1977
				break
			}
			if s[j] == '0' {
				continue
			}
			if n-j < curLen {
				continue
			}
			l := lcp[i][j]
			if l >= curLen || s[i+l] < s[j+l] {
				total = (total + pref[j+curLen]) % MOD1977
			} else {
				total = (total + pref[j+curLen+1]) % MOD1977
			}
		}
		dp[i] = total
		pref[i] = (dp[i] + pref[i+1]) % MOD1977
	}

	return dp[0]
}

func main() {
	// Example: "327" -> 2
	fmt.Println(numberOfWaysToSeparateNumbers("327"))

	// Additional tests
	fmt.Println(numberOfWaysToSeparateNumbers("123"))
	fmt.Println(numberOfWaysToSeparateNumbers("100"))
	fmt.Println(numberOfWaysToSeparateNumbers("1"))
}
```
