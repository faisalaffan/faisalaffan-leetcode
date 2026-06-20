# 1269 — Number Of Ways To Stay In The Same Place After Some Steps

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numWays(steps int, arrLen int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1269: Number of Ways to Stay in the Same Place After Some Steps
// https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps/
// Difficulty: Hard

import "fmt"

func main() {
	fmt.Println("1269. Number of Ways to Stay in the Same Place After Some Steps")
	fmt.Println("steps=3, arrLen=3:", numWays(3, 3), "(expected 4)")
	fmt.Println("steps=2, arrLen=4:", numWays(2, 4), "(expected 2)")
	fmt.Println("steps=4, arrLen=2:", numWays(4, 2), "(expected 8)")
}

func numWays(steps int, arrLen int) int {
	const MOD = 1000000007

	// Max reachable position is min(steps, arrLen-1).
	maxPos := steps
	if arrLen-1 < maxPos {
		maxPos = arrLen - 1
	}

  // Alokasi slice
	dp := make([]int, maxPos+1)
	dp[0] = 1

	for s := 1; s <= steps; s++ {
  // Alokasi slice
		ndp := make([]int, maxPos+1)
		for pos := 0; pos <= maxPos; pos++ {
			ways := dp[pos] // stay
			if pos > 0 {
				ways = (ways + dp[pos-1]) % MOD // move right (from pos-1)
			}
			if pos < maxPos {
				ways = (ways + dp[pos+1]) % MOD // move left (from pos+1)
			}
			ndp[pos] = ways
		}
		dp = ndp
	}

	return dp[0]
}
```
