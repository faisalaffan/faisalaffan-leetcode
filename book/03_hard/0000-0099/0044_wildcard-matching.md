# 0044 — Wildcard Matching

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func isMatchWildcard(s string, p string) bool
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #44: Wildcard Matching
// https://leetcode.com/problems/wildcard-matching/
// Difficulty: Hard

import "fmt"

// isMatchWildcard checks if string s matches pattern p.
// '?' matches any single character.
// '*' matches any sequence of characters (including empty).
//
// Complexity: O(m*n) time, O(m*n) space where m = len(s), n = len(p)
func isMatchWildcard(s string, p string) bool {
	m, n := len(s), len(p)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]bool, m+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// empty string matches empty pattern
	dp[0][0] = true

	// handle patterns starting with '*' matching empty string
	for j := 1; j <= n; j++ {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-1]
		}
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if p[j-1] == '?' || p[j-1] == s[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else if p[j-1] == '*' {
				// '*' matches empty (dp[i][j-1]) or one/more chars (dp[i-1][j])
				dp[i][j] = dp[i][j-1] || dp[i-1][j]
			}
		}
	}

	return dp[m][n]
}

func main() {
	// Test cases from LeetCode
	fmt.Println("Test 1: aa, a ->", isMatchWildcard("aa", "a"))     // false
	fmt.Println("Test 2: aa, * ->", isMatchWildcard("aa", "*"))     // true
	fmt.Println("Test 3: cb, ?a ->", isMatchWildcard("cb", "?a"))   // false
	fmt.Println("Test 4: adceb, *a*b ->", isMatchWildcard("adceb", "*a*b")) // true
	fmt.Println("Test 5: acdcb, a*c?b ->", isMatchWildcard("acdcb", "a*c?b")) // false
}
```
