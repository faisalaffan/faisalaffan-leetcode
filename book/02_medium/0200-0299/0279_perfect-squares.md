# 0279 — Perfect Squares

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func numSquares(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n * sqrt(n)), Space: O(n)  
**Kompleksitas Ruang:** O(n)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #279: Perfect Squares
// https://leetcode.com/problems/perfect-squares/
// Difficulty: Medium
// Time: O(n * sqrt(n)), Space: O(n)

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
  // Alokasi slice integer
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			sq := j * j
			if 1+dp[i-sq] < dp[i] {
				dp[i] = 1 + dp[i-sq]
			}
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(1))
}
```
