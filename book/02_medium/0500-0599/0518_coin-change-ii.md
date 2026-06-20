# 0518 — Coin Change Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func CoinChangeIi(amount int, coins []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(amount * n) where n = len(coins)  |  **Ruang:** O(amount)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

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
  // Alokasi slice
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
