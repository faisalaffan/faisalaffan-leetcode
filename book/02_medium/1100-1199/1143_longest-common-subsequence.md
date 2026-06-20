# 1143 — Longest Common Subsequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** —

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func longestCommonSubsequence(text1 string, text2 string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1143: Longest Common Subsequence
// https://leetcode.com/problems/longest-common-subsequence/

// Given two strings text1 and text2, return the length of their
// longest common subsequence. If there is no common subsequence, return 0.

// Solution uses classic 2D DP:
// dp[i][j] = LCS of text1[:i] and text2[:j]
// If text1[i-1] == text2[j-1]: dp[i][j] = dp[i-1][j-1] + 1
// Else: dp[i][j] = max(dp[i-1][j], dp[i][j-1])

// Time complexity: O(m*n)
// Space complexity: O(m*n), can be optimized to O(min(m,n))

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1), len(text2)
  // Matriks 2D
	dp := make([][]int, m+1)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[i-1] == text2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				if dp[i-1][j] > dp[i][j-1] {
					dp[i][j] = dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-1]
				}
			}
		}
	}

	return dp[m][n]
}

func main() {
	// Test case 1
	fmt.Printf("longestCommonSubsequence(%q, %q) = %d (expected: 3)\n",
		"abcde", "ace", longestCommonSubsequence("abcde", "ace"))

	// Test case 2
	fmt.Printf("longestCommonSubsequence(%q, %q) = %d (expected: 3)\n",
		"abc", "abc", longestCommonSubsequence("abc", "abc"))

	// Test case 3: No common subsequence
	fmt.Printf("longestCommonSubsequence(%q, %q) = %d (expected: 0)\n",
		"abc", "def", longestCommonSubsequence("abc", "def"))
}
```
