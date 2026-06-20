# 1866 — Number Of Ways To Rearrange Sticks With K Sticks Visible

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numberOfWaysToRearrangeSticks(n int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1866: Number of Ways to Rearrange Sticks With K Sticks Visible
// https://leetcode.com/problems/number-of-ways-to-rearrange-sticks-with-k-sticks-visible/
// Difficulty: Hard

import "fmt"

func numberOfWaysToRearrangeSticks(n int, k int) int {
	const mod = 1_000_000_007
	// dp[i][j] = ways to arrange i sticks so exactly j are visible
	// Recurrence: dp[i][j] = dp[i-1][j-1] + (i-1)*dp[i-1][j]
	// This is the unsigned Stirling numbers of the first kind
  // Matriks 2D
	dp := make([][]int, n+1)
  // Range loop
	for i := range dp {
		dp[i] = make([]int, k+1)
	}
	dp[0][0] = 1
	for i := 1; i <= n; i++ {
		for j := 1; j <= min(i, k); j++ {
			dp[i][j] = (dp[i-1][j-1] + (i-1)*dp[i-1][j]%mod) % mod
		}
	}
	return dp[n][k]
}

func main() {
	// Example: n=3, k=2 -> 3
	// [1,3,2], [2,3,1], [2,1,3]
	fmt.Println(numberOfWaysToRearrangeSticks(3, 2))

	// Additional test
	fmt.Println(numberOfWaysToRearrangeSticks(5, 3))
}
```
