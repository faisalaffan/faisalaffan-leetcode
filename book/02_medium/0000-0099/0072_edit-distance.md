# 0072 — Edit Distance

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func minDistance(word1 string, word2 string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(m*n)  |  **Ruang:** O(m*n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #72: Edit Distance
// https://leetcode.com/problems/edit-distance/
// Difficulty: Medium

import "fmt"

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)

  // Matriks 2D
	dp := make([][]int, m+1)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n+1)
		dp[i][0] = i
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = 1 + min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1])
			}
		}
	}

	return dp[m][n]
}

func main() {
	// Test case 1
	fmt.Println(minDistance("horse", "ros")) // 3

	// Test case 2
	fmt.Println(minDistance("intention", "execution")) // 5

	// Test case 3
	fmt.Println(minDistance("", "a")) // 1
}

// Time: O(m*n) | Space: O(m*n)
```
