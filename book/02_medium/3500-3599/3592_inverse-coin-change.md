# 3592 — Inverse Coin Change

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func InverseCoinChange(coins []int, amount int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3592: Inverse Coin Change
// https://leetcode.com/problems/inverse-coin-change/
// Difficulty: Medium
// Complexity: O(amount * n) time, O(amount) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", InverseCoinChange([]int{1, 2, 5}, 11))
	// Test case 2
	fmt.Println("Test 2:", InverseCoinChange([]int{2}, 3))
	// Test case 3
	fmt.Println("Test 3:", InverseCoinChange([]int{1}, 0))
}

func InverseCoinChange(coins []int, amount int) int {
	// Minimum number of coins to make amount
  // Alokasi slice integer
	dp := make([]int, amount+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i && dp[i-coin]+1 < dp[i] {
				dp[i] = dp[i-coin] + 1
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
```
