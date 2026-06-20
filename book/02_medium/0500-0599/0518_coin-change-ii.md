# 0518 — Coin Change Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func CoinChangeIi(amount int, coins []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(amount * n) where n = len(coins)  
**Kompleksitas Ruang:** O(amount)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #518: Coin Change II
// https://leetcode.com/problems/coin-change-ii/
// Difficulty: Medium
// Time: O(amount * n) where n = len(coins)
// Space: O(amount)

import "fmt"

func main() {
	fmt.Println(CoinChangeIi(5, []int{1, 2, 5}))
	fmt.Println(CoinChangeIi(3, []int{2}))
	fmt.Println(CoinChangeIi(10, []int{10}))
}

func CoinChangeIi(amount int, coins []int) int {
  // Alokasi slice integer
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}
```
