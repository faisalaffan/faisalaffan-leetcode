# 3339 — Find The Number Of K Even Arrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func countKEvenArrays(n int, m int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n * k) Space: O(n * k)  |  **Ruang:** O(n * k)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3339: Find the Number of K-Even Arrays
// https://leetcode.com/problems/find-the-number-of-k-even-arrays/
// Difficulty: Medium
// Time: O(n * k) Space: O(n * k)

import (
	"fmt"
)

func main() {
	fmt.Println(countKEvenArrays(3, 4, 2)) // 8
	fmt.Println(countKEvenArrays(5, 1, 0)) // 1
	fmt.Println(countKEvenArrays(7, 7, 5)) // 5832
}

func countKEvenArrays(n int, m int, k int) int {
	const mod = 1_000_000_007

	evens := m / 2
	odds := m - evens

	// dp[pos][pairs][parity] where parity 0=even, 1=odd
  // Matriks 2D
	dp := make([][][]int, n+1)
  // Range loop
	for i := range dp {
		dp[i] = make([][]int, k+2)
		for j := range dp[i] {
			dp[i][j] = make([]int, 2)
		}
	}

	// Base: first element (1-indexed)
	dp[1][0][0] = evens
	dp[1][0][1] = odds

	for i := 2; i <= n; i++ {
		for j := 0; j <= k; j++ {
			// Place even: adds pair if prev was even
			if evens > 0 {
				dp[i][j][0] = (dp[i-1][j][1] * evens) % mod
				if j > 0 {
					dp[i][j][0] = (dp[i][j][0] + dp[i-1][j-1][0]*evens) % mod
				}
			}
			// Place odd: never adds a pair
			if odds > 0 {
				dp[i][j][1] = (dp[i-1][j][0] + dp[i-1][j][1]) % mod
				dp[i][j][1] = (dp[i][j][1] * odds) % mod
			}
		}
	}

	return (dp[n][k][0] + dp[n][k][1]) % mod
}
```
