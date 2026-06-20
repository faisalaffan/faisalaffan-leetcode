# 0188 — Best Time To Buy And Sell Stock Iv

## Deskripsi

**Soal:** [0188. Best Time To Buy And Sell Stock Iv](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Greedy (pemilihan optimal lokal)

**Fungsi Solusi:** `func maxProfit(k int, prices []int) int`

## Solusi Go

```go
package main

// LeetCode #188: Best Time to Buy and Sell Stock IV
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iv/
// Difficulty: Hard

import (
	"fmt"
)

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}

	// If k >= n/2, we can do unlimited transactions (greedy)
	if k >= n/2 {
		profit := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				profit += prices[i] - prices[i-1]
			}
		}
		return profit
	}

	// dp[i][j] = max profit with at most i transactions up to day j
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int, k+1)
  // Iterasi seluruh elemen
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 1; i <= k; i++ {
		maxDiff := -prices[0] // max(dp[i-1][t] - prices[t]) for t < j
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], prices[j]+maxDiff)
			maxDiff = max(maxDiff, dp[i-1][j]-prices[j])
		}
	}

	return dp[k][n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	k := 2
	prices := []int{3, 2, 6, 5, 0, 3}
	result := maxProfit(k, prices)
	expected := 7

	fmt.Printf("maxProfit(%d, %v) = %d\n", k, prices, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```
