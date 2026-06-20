# 1092 — Shortest Common Supersequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func shortestCommonSupersequence(str1 string, str2 string) string`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Backtracking

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1092: Shortest Common Supersequence
// https://leetcode.com/problems/shortest-common-supersequence/
// Difficulty: Hard
//
// Compute LCS via DP, then backtrack to build the SCS by merging str1 and
// str2 while including LCS characters only once.
// SCS length = len(str1) + len(str2) - LCS length.

import "fmt"

func main() {
	fmt.Println(shortestCommonSupersequence("abac", "cab"))
}

func shortestCommonSupersequence(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
  // Matriks 2D
	dp := make([][]int, m+1)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	// Build LCS length table
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else if dp[i-1][j] > dp[i][j-1] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1]
			}
		}
	}

	// Backtrack to build SCS in reverse
	res := make([]byte, 0, m+n-dp[m][n])
	i, j := m, n
	for i > 0 || j > 0 {
		if i == 0 {
			j--
			res = append(res, str2[j])
		} else if j == 0 {
			i--
			res = append(res, str1[i])
		} else if str1[i-1] == str2[j-1] {
			i--
			j--
			res = append(res, str1[i])
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
			res = append(res, str1[i])
		} else {
			j--
			res = append(res, str2[j])
		}
	}

	// Reverse the result
	for l, r := 0, len(res)-1; l < r; l, r = l+1, r-1 {
		res[l], res[r] = res[r], res[l]
	}

	return string(res)
}
```
