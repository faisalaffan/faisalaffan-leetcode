# 0474 — Ones And Zeroes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func OnesAndZeroes(strs []string, m int, n int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(m * n * len(strs))  |  **Ruang:** O(m * n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #474: Ones and Zeroes
// https://leetcode.com/problems/ones-and-zeroes/
// Difficulty: Medium
// Time: O(m * n * len(strs))
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(OnesAndZeroes([]string{"10", "0001", "111001", "1", "0"}, 5, 3))
	fmt.Println(OnesAndZeroes([]string{"10", "0", "1"}, 1, 1))
}

func OnesAndZeroes(strs []string, m int, n int) int {
  // Matriks 2D
	dp := make([][]int, m+1)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, s := range strs {
		zeros, ones := countBits(s)
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				if dp[i-zeros][j-ones]+1 > dp[i][j] {
					dp[i][j] = dp[i-zeros][j-ones] + 1
				}
			}
		}
	}

	return dp[m][n]
}

func countBits(s string) (int, int) {
	zeros, ones := 0, 0
	for _, c := range s {
		if c == '0' {
			zeros++
		} else {
			ones++
		}
	}
	return zeros, ones
}
```
