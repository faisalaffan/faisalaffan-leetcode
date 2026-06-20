# 2466 — Count Ways To Build Good Strings

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan string. Tugasmu memanipulasi atau mencari pola dalam teks.

**Cara berpikir:** String immutable di Go — konversi ke `[]byte` untuk modifikasi. Operasi: iterasi karakter, substring `s[i:j]`.

**Fungsi Solusi:** `func countGoodStrings(low int, high int, zero int, one int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(high)  |  **Ruang:** O(high)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2466: Count Ways To Build Good Strings
// https://leetcode.com/problems/count-ways-to-build-good-strings/
// Difficulty: Medium
// Time: O(high) | Space: O(high)
// DP: dp[i] = ways to build string of length i.
// dp[i] = dp[i-zero] + dp[i-one] (if i >= zero/one).

import "fmt"

func main() {
	fmt.Println(countGoodStrings(3, 3, 1, 1)) // 8
	fmt.Println(countGoodStrings(2, 3, 1, 2)) // 5
}

const MOD = 1000000007

func countGoodStrings(low int, high int, zero int, one int) int {
  // Alokasi slice
	dp := make([]int, high+1)
	dp[0] = 1
	ans := 0

	for i := 1; i <= high; i++ {
		if i >= zero {
			dp[i] = (dp[i] + dp[i-zero]) % MOD
		}
		if i >= one {
			dp[i] = (dp[i] + dp[i-one]) % MOD
		}
		if i >= low {
			ans = (ans + dp[i]) % MOD
		}
	}
	return ans
}
```
