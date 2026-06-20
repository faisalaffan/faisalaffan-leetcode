# 0940 — Distinct Subsequences Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func distinctSubseqII(s string) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #940: Distinct Subsequences II
// https://leetcode.com/problems/distinct-subsequences-ii/
// Difficulty: Hard
// DP with last occurrence tracking.
// dp[i] = 2*dp[i-1] - dp[last[s[i-1]]-1] (mod 1e9+7)
// Result counts non-empty subsequences.

import "fmt"

func distinctSubseqII(s string) int {
	mod := int(1e9 + 7)
	n := len(s)
  // Alokasi slice integer
	dp := make([]int, n+1)
	dp[0] = 1 // empty subsequence

  // Membuat map (HashMap) — pencarian O(1)
	last := make(map[byte]int) // last occurrence index (1-based)

	for i := 1; i <= n; i++ {
		dp[i] = (dp[i-1] * 2) % mod
		c := s[i-1]
		if prev, ok := last[c]; ok {
			dp[i] = (dp[i] - dp[prev-1] + mod) % mod
		}
		last[c] = i
	}

	// Subtract empty subsequence
	return (dp[n] - 1 + mod) % mod
}

func main() {
	fmt.Println(distinctSubseqII("abc")) // Expected: 7
	fmt.Println(distinctSubseqII("aba")) // Expected: 6
	fmt.Println(distinctSubseqII("aaa")) // Expected: 3
	fmt.Println(distinctSubseqII("ab"))  // Expected: 3
	fmt.Println(distinctSubseqII("abca")) // Expected: 14?
}
```
