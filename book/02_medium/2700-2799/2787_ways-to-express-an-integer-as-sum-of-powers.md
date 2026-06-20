# 2787 — Ways To Express An Integer As Sum Of Powers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func WaysToExpressAnIntegerAsSumOfPowers(n int, x int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2787: Ways to Express an Integer as Sum of Powers
// https://leetcode.com/problems/ways-to-express-an-integer-as-sum-of-powers/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import "fmt"

func WaysToExpressAnIntegerAsSumOfPowers(n int, x int) int {
	const mod = 1_000_000_007

	// Generate powers
  // Alokasi slice
	powers := make([]int, 0)
	for i := 1; ; i++ {
		p := 1
		for j := 0; j < x; j++ {
			p *= i
		}
		if p > n {
			break
		}
		powers = append(powers, p)
	}

  // Alokasi slice
	dp := make([]int, n+1)
	dp[0] = 1

	for _, p := range powers {
		for s := n; s >= p; s-- {
			dp[s] = (dp[s] + dp[s-p]) % mod
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(WaysToExpressAnIntegerAsSumOfPowers(10, 2))
	fmt.Println(WaysToExpressAnIntegerAsSumOfPowers(4, 1))
}
```
