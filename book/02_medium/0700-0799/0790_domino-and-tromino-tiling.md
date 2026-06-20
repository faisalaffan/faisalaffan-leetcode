# 0790 — Domino And Tromino Tiling

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numTilings(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #790: Domino and Tromino Tiling
// https://leetcode.com/problems/domino-and-tromino-tiling/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(numTilings(3))
	fmt.Println(numTilings(1))
	fmt.Println(numTilings(5))
}

func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	mod := 1000000007
  // Alokasi slice integer
	dp := make([]int, n+1)
  // Alokasi slice integer
	dp2 := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	dp2[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + 2*dp2[i-1]) % mod
		dp2[i] = (dp[i-2] + dp2[i-1]) % mod
	}

	return dp[n]
}
```
