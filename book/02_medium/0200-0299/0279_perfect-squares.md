# 0279 — Perfect Squares

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numSquares(n int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n * sqrt(n)), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #279: Perfect Squares
// https://leetcode.com/problems/perfect-squares/
// Difficulty: Medium
// Time: O(n * sqrt(n)), Space: O(n)

import (
	"fmt"
	"math"
)

func numSquares(n int) int {
  // Alokasi slice
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = math.MaxInt32
	}

	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			sq := j * j
			if 1+dp[i-sq] < dp[i] {
				dp[i] = 1 + dp[i-sq]
			}
		}
	}

	return dp[n]
}

func main() {
	fmt.Println(numSquares(12))
	fmt.Println(numSquares(13))
	fmt.Println(numSquares(1))
}
```
