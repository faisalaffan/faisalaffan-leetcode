# 1062 — Longest Repeating Substring

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func longestRepeatingSubstring(s string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Prefix Sum

**Waktu:** O(n^2)  |  **Ruang:** O(n^2) can be O(n) with optimized DP

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1062: Longest Repeating Substring
// https://leetcode.com/problems/longest-repeating-substring/
// Difficulty: Medium
//
// Approach: DP - find longest common prefix between all pairs of suffixes
// Time: O(n^2)
// Space: O(n^2) can be O(n) with optimized DP

import "fmt"

func main() {
	fmt.Println(longestRepeatingSubstring("abcd"))    // 0
	fmt.Println(longestRepeatingSubstring("abbaba"))  // 2
	fmt.Println(longestRepeatingSubstring("aabcaabdaab")) // 3
}

func longestRepeatingSubstring(s string) int {
	n := len(s)
  // Matriks 2D
	dp := make([][]int, n+1)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	result := 0
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			if s[i-1] == s[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
				if dp[i][j] > result {
					result = dp[i][j]
				}
			}
		}
	}

	return result
}
```
