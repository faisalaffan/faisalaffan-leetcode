# 0629 — K Inverse Pairs Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func kInversePairs(n int, k int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #629: K Inverse Pairs Array
// https://leetcode.com/problems/k-inverse-pairs-array/
// Difficulty: Hard

import "fmt"

const mod = 1_000_000_007

func main() {
	// Test cases
	testCases := []struct {
		n    int
		k    int
		want int
	}{
		{3, 0, 1},
		{3, 1, 2},
		{3, 2, 2},
		{3, 3, 1},
		{4, 0, 1},
		{4, 1, 3},
		{4, 2, 5},
		{4, 3, 6},
		{4, 4, 5},
		{4, 5, 3},
		{4, 6, 1},
		{10, 5, 1068},
		{1000, 1000, 663677020},
	}

	for _, tc := range testCases {
		got := kInversePairs(tc.n, tc.k)
		status := "PASS"
		if got != tc.want {
			status = "FAIL"
		}
		fmt.Printf("%s: kInversePairs(%d, %d) = %d (want %d)\n", status, tc.n, tc.k, got, tc.want)
	}
}

func kInversePairs(n int, k int) int {
	// dp[j] = number of arrays of current size with exactly j inverse pairs
  // Alokasi slice integer
	dp := make([]int, k+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		// prefix sum helper
  // Alokasi slice integer
		next := make([]int, k+1)
		prefix := 0
		for j := 0; j <= k; j++ {
			prefix = (prefix + dp[j]) % mod
			if j-i >= 0 {
				prefix = (prefix - dp[j-i] + mod) % mod
			}
			next[j] = prefix
		}
		dp = next
	}

	return dp[k]
}
```
