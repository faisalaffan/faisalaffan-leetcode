# 2291 — Maximum Profit From Trading Stocks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumProfit(presentValues []int, futureValues []int, budget int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n * budget)  |  **Ruang:** O(budget)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2291: Maximum Profit From Trading Stocks
// https://leetcode.com/problems/maximum-profit-from-trading-stocks/
// Difficulty: Medium [Paid]
// Time: O(n * budget) | Space: O(budget)

import "fmt"

func maximumProfit(presentValues []int, futureValues []int, budget int) int {
	n := len(presentValues)
	// dp[b] = max profit with budget b
  // Alokasi slice
	dp := make([]int, budget+1)

	for i := 0; i < n; i++ {
		profit := futureValues[i] - presentValues[i]
		if profit <= 0 {
			continue
		}
		cost := presentValues[i]
		for b := budget; b >= cost; b-- {
			if dp[b-cost]+profit > dp[b] {
				dp[b] = dp[b-cost] + profit
			}
		}
	}
	return dp[budget]
}

func main() {
	// Test case 1
	fmt.Println(maximumProfit([]int{5, 4, 6, 2}, []int{6, 5, 7, 3}, 6))
	// Expected: 2

	// Test case 2
	fmt.Println(maximumProfit([]int{5, 4, 6, 2}, []int{4, 5, 3, 3}, 5))
	// Expected: 0
}
```
