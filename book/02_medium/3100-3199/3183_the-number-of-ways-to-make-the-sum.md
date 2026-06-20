# 3183 — The Number Of Ways To Make The Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numberOfWays(n int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3183: The Number of Ways to Make the Sum
// https://leetcode.com/problems/the-number-of-ways-to-make-the-sum/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func numberOfWays(n int) int {
	const mod = 1000000007
  // Alokasi slice
	dp := make([]int, n+1)
	dp[0] = 1

	for _, coin := range []int{1, 2, 6} {
		for i := coin; i <= n; i++ {
			dp[i] = (dp[i] + dp[i-coin]) % mod
		}
	}

	ans := dp[n]
	if n-4 >= 0 {
		ans = (ans + dp[n-4]) % mod
	}
	if n-8 >= 0 {
		ans = (ans + dp[n-8]) % mod
	}
	return ans
}

func main() {
	fmt.Println(numberOfWays(4)) // Expected: 5
	fmt.Println(numberOfWays(1)) // Expected: 1
	fmt.Println(numberOfWays(6)) // Expected: 7
}
```
