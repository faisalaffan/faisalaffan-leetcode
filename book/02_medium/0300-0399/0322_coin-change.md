# 0322 — Coin Change

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func coinChange(coins []int, amount int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n * amount)  |  **Ruang:** O(amount)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #322: Coin Change
// https://leetcode.com/problems/coin-change/
// Difficulty: Medium
// Time: O(n * amount) | Space: O(amount)

import "fmt"

func coinChange(coins []int, amount int) int {
  // Alokasi slice
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if c <= i && dp[i-c]+1 < dp[i] {
				dp[i] = dp[i-c] + 1
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", coinChange([]int{1, 2, 5}, 11))
	// Expected: 3 (5+5+1)

	// Test case 2
	fmt.Println("Test 2:", coinChange([]int{2}, 3))
	// Expected: -1

	// Test case 3
	fmt.Println("Test 3:", coinChange([]int{1}, 0))
	// Expected: 0
}
```
