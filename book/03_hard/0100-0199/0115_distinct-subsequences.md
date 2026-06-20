# 0115 — Distinct Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numDistinct(s string, t string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #115: Distinct Subsequences
// https://leetcode.com/problems/distinct-subsequences/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("115. Distinct Subsequences")
	fmt.Println("rabbbit, rabbit ->", numDistinct("rabbbit", "rabbit"), "(expected 3)")
	fmt.Println("babgbag, bag ->", numDistinct("babgbag", "bag"), "(expected 5)")
	fmt.Println("r, r ->", numDistinct("r", "r"), "(expected 1)")
}

func numDistinct(s string, t string) int {
	m, n := len(s), len(t)
  // Edge case: input kosong — langsung return
	if n == 0 {
		return 1
	}
	if m < n {
		return 0
	}

	// dp[j] = number of distinct subsequences of s[:i] that equal t[:j]
  // Alokasi slice integer
	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= m; i++ {
		prev := dp[0]
		for j := 1; j <= n; j++ {
			curr := dp[j]
			if s[i-1] == t[j-1] {
				dp[j] = prev + dp[j]
			}
			prev = curr
		}
	}

	return dp[n]
}
```
