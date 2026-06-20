# 0518 — Coin Change Ii

## Deskripsi

**Soal:** [0518. Coin Change Ii](https://leetcode.com/problems/coin-change-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(amount * n) where n = len(coins)  
**Kompleksitas Ruang:** O(amount)

**Algoritma:** Dynamic Programming (DP)

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
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
