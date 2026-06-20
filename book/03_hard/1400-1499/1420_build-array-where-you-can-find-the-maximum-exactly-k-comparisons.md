# 1420 — Build Array Where You Can Find The Maximum Exactly K Comparisons

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numOfArrays(n int, m int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1420: Build Array Where You Can Find The Maximum Exactly K Comparisons
// https://leetcode.com/problems/build-array-where-you-can-find-the-maximum-exactly-k-comparisons/
// Difficulty: Hard

import "fmt"

const mod1420 = 1_000_000_007

func numOfArrays(n int, m int, k int) int {
	if k == 0 || k > m {
		return 0
	}
	// dp[i][j][c] = ways for length i, max = j, cost = c
  // Membuat matriks/slice 2D untuk DP
	dp := make([][][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([][]int, m+1)
		for j := range dp[i] {
			dp[i][j] = make([]int, k+1)
		}
	}

	for j := 1; j <= m; j++ {
		dp[1][j][1] = 1
	}

	for i := 2; i <= n; i++ {
		for j := 1; j <= m; j++ {
			for c := 1; c <= k; c++ {
				// Append value <= j: choose any of j values, cost unchanged
				dp[i][j][c] = (dp[i][j][c] + dp[i-1][j][c]*j) % mod1420

				// Append value == j (new max): sum over previous max < j
				if c > 1 {
					for p := 1; p < j; p++ {
						dp[i][j][c] = (dp[i][j][c] + dp[i-1][p][c-1]) % mod1420
					}
				}
			}
		}
	}

	var ans int
	for j := 1; j <= m; j++ {
		ans = (ans + dp[n][j][k]) % mod1420
	}
	return ans
}

func main() {
	// Example: n=2, m=3, k=1 -> 6
	fmt.Println(numOfArrays(2, 3, 1))
}
```
