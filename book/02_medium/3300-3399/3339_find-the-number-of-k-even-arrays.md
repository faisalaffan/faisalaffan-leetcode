# 3339 — Find The Number Of K Even Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func countKEvenArrays(n int, m int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n * k) Space: O(n * k)  
**Kompleksitas Ruang:** O(n * k)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3339: Find the Number of K-Even Arrays
// https://leetcode.com/problems/find-the-number-of-k-even-arrays/
// Difficulty: Medium
// Time: O(n * k) Space: O(n * k)

import (
	"fmt"
)

func main() {
	fmt.Println(countKEvenArrays(3, 4, 2)) // 8
	fmt.Println(countKEvenArrays(5, 1, 0)) // 1
	fmt.Println(countKEvenArrays(7, 7, 5)) // 5832
}

func countKEvenArrays(n int, m int, k int) int {
	const mod = 1_000_000_007

	evens := m / 2
	odds := m - evens

	// dp[pos][pairs][parity] where parity 0=even, 1=odd
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([][]int, k+2)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	// Base: first element (1-indexed)
	dp[1][0][0] = evens
	dp[1][0][1] = odds

	for i := 2; i <= n; i++ {
		for j := 0; j <= k; j++ {
			// Place even: adds pair if prev was even
			if evens > 0 {
				dp[i][j][0] = (dp[i-1][j][1] * evens) % mod
				if j > 0 {
					dp[i][j][0] = (dp[i][j][0] + dp[i-1][j-1][0]*evens) % mod
				}
			}
			// Place odd: never adds a pair
			if odds > 0 {
				dp[i][j][1] = (dp[i-1][j][0] + dp[i-1][j][1]) % mod
				dp[i][j][1] = (dp[i][j][1] * odds) % mod
			}
		}
	}

	return (dp[n][k][0] + dp[n][k][1]) % mod
}
```
