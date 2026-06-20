# 2969 — Minimum Number Of Coins For Fruits Ii

## Deskripsi

**Soal:** [2969. Minimum Number Of Coins For Fruits Ii](https://leetcode.com/problems/minimum-number-of-coins-for-fruits-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func minimumCoins(prices []int) int`

## Solusi Go

```go
package main

// LeetCode #2969: Minimum Number of Coins for Fruits II
// https://leetcode.com/problems/minimum-number-of-coins-for-fruits-ii/
// Difficulty: Hard
//
// You have n types of fruits, each with a price. When you buy fruit i,
// you get fruits i+1, i+2, ..., 2*i+1 for free (or up to n-1).
// Find minimum coins to acquire all fruits.
//
// DP from right to left: dp[i] = min cost to acquire fruits i..n-1.
// dp[i] = prices[i] + min(dp[j]) for j in [i+1, 2*i+2].
// Use monotonic deque to maintain min dp[j] in the range.

import "fmt"

func minimumCoins(prices []int) int {
	n := len(prices)
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, n+1)
	for i := 0; i < n; i++ {
		dp[i] = int(1e9)
	}
	dp[n] = 0

  // Membuat slice untuk menyimpan hasil
	dq := make([]int, 0, n)
	dq = append(dq, n)

	for i := n - 1; i >= 0; i-- {
		// Remove indices that are out of range [i+1, 2*i+2]
		// Since we process right to left, we remove indices > 2*i+2
		for len(dq) > 0 && dq[0] > 2*i+2 {
			dq = dq[1:]
		}
		dp[i] = prices[i] + dp[dq[0]]
		// Maintain increasing order of dp values
		for len(dq) > 0 && dp[dq[len(dq)-1]] >= dp[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
	}
	return dp[0]
}

func main() {
	// Example
	fmt.Println(minimumCoins([]int{3, 1, 2}))
	fmt.Println(minimumCoins([]int{1, 10, 1, 1}))

	// Edge cases
	fmt.Println(minimumCoins([]int{5}))
	fmt.Println(minimumCoins([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}))
}
```
