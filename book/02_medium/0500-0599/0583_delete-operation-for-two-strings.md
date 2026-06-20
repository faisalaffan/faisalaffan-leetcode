# 0583 — Delete Operation For Two Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func MinDistance(word1 string, word2 string) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(m * n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #583: Delete Operation for Two Strings
// https://leetcode.com/problems/delete-operation-for-two-strings/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(MinDistance("sea", "eat"))
	fmt.Println(MinDistance("leetcode", "etco"))
}

func MinDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
  // Alokasi slice
	dp := make([]int, n+1)

	for i := 1; i <= m; i++ {
		prev := 0
		for j := 1; j <= n; j++ {
			temp := dp[j]
			if word1[i-1] == word2[j-1] {
				dp[j] = prev + 1
			} else {
				if dp[j] > dp[j-1] {
					dp[j] = dp[j]
				} else {
					dp[j] = dp[j-1]
				}
			}
			prev = temp
		}
	}

	lcs := dp[n]
	return (m - lcs) + (n - lcs)
}
```
