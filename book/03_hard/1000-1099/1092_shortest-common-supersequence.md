# 1092 — Shortest Common Supersequence

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func shortestCommonSupersequence(str1 string, str2 string) string
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Backtracking

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, m+1)
  // Range loop: iterasi dengan indeks + nilai
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
