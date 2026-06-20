# 0712 — Minimum Ascii Delete Sum For Two Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah string (teks). Tugasmu adalah memanipulasi, mencari pola, atau menghitung sesuatu dari string tersebut.

Ibarat kamu sedang mengedit dokumen teks — kamu perlu mencari kata tertentu, menghitung huruf, atau mengubah format teks. String di Go adalah slice of byte yang immutable (tidak bisa diubah langsung, harus dikonversi ke `[]byte` dulu).

**Konsep kunci:** karakter, substring, prefix/suffix, konversi `string` ↔ `[]byte`.

**Fungsi yang perlu kamu implementasikan:**
```go
func minimumDeleteSum(s1 string, s2 string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(m * n)  
**Kompleksitas Ruang:** O(m * n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #712: Minimum ASCII Delete Sum for Two Strings
// https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/
// Difficulty: Medium
// Time: O(m * n)
// Space: O(m * n)

import "fmt"

func main() {
	fmt.Println(minimumDeleteSum("sea", "eat"))
	fmt.Println(minimumDeleteSum("delete", "leet"))
}

func minimumDeleteSum(s1 string, s2 string) int {
	m, n := len(s1), len(s2)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i <= m; i++ {
		dp[i][0] = dp[i-1][0] + int(s1[i-1])
	}
	for j := 1; j <= n; j++ {
		dp[0][j] = dp[0][j-1] + int(s2[j-1])
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j]+int(s1[i-1]), dp[i][j-1]+int(s2[j-1]))
			}
		}
	}

	return dp[m][n]
}
```
