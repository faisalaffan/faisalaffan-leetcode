# 0935 — Knight Dialer

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func knightDialer(n int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #935: Knight Dialer
// https://leetcode.com/problems/knight-dialer/
// Difficulty: Medium

import "fmt"

const mod = 1_000_000_007

// Time: O(n) | Space: O(1)
func knightDialer(n int) int {
	dp := [10]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	for i := 1; i < n; i++ {
		dp = [10]int{
			(dp[4] + dp[6]) % mod,
			(dp[6] + dp[8]) % mod,
			(dp[7] + dp[9]) % mod,
			(dp[4] + dp[8]) % mod,
			(dp[0] + dp[3] + dp[9]) % mod,
			0,
			(dp[0] + dp[1] + dp[7]) % mod,
			(dp[2] + dp[6]) % mod,
			(dp[1] + dp[3]) % mod,
			(dp[2] + dp[4]) % mod,
		}
	}
	sum := 0
	for _, v := range dp {
		sum = (sum + v) % mod
	}
	return sum
}

func main() {
	fmt.Println(knightDialer(1))
	fmt.Println(knightDialer(2))
	fmt.Println(knightDialer(3131))
}
```
