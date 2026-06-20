# 0122 — Best Time To Buy And Sell Stock Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxProfit(prices []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #122: Best Time to Buy and Sell Stock II
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/
// Difficulty: Medium

import "fmt"

func maxProfit(prices []int) int {
	profit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			profit += prices[i] - prices[i-1]
		}
	}
	return profit
}

func main() {
	// Test case 1
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 7

	// Test case 2
	fmt.Println(maxProfit([]int{1, 2, 3, 4, 5})) // 4

	// Test case 3
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1})) // 0
}

// Time: O(n) | Space: O(1)
```
