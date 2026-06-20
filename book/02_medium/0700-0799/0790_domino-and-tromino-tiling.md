# 0790 — Domino And Tromino Tiling

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numTilings(n int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #790: Domino and Tromino Tiling
// https://leetcode.com/problems/domino-and-tromino-tiling/
// Difficulty: Medium
// Time: O(n)
// Space: O(n)

import "fmt"

func main() {
	fmt.Println(numTilings(3))
	fmt.Println(numTilings(1))
	fmt.Println(numTilings(5))
}

func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	mod := 1000000007
  // Alokasi slice
	dp := make([]int, n+1)
  // Alokasi slice
	dp2 := make([]int, n+1)

	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	dp2[2] = 1

	for i := 3; i <= n; i++ {
		dp[i] = (dp[i-1] + dp[i-2] + 2*dp2[i-1]) % mod
		dp2[i] = (dp[i-2] + dp2[i-1]) % mod
	}

	return dp[n]
}
```
