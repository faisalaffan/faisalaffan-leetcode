# 2291 — Maximum Profit From Trading Stocks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maximumProfit(presentValues []int, futureValues []int, budget int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n * budget)  
**Kompleksitas Ruang:** O(budget)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
  // Alokasi slice integer
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
