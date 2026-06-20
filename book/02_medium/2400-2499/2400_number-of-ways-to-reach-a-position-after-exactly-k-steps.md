# 2400 — Number Of Ways To Reach A Position After Exactly K Steps

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan bilangan bulat. Tugasmu menghitung atau menganalisis properti bilangan.

**Cara berpikir:** Modulo `%` ambil digit terakhir. Pembagian `/` buang digit. Untuk reverse: `rev = rev*10 + digit`.

**Fungsi Solusi:** `func numberOfWays(startPos int, endPos int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, DP

**Waktu:** O(k^2)  |  **Ruang:** O(k)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2400: Number of Ways to Reach a Position After Exactly k Steps
// https://leetcode.com/problems/number-of-ways-to-reach-a-position-after-exactly-k-steps/
// Difficulty: Medium
// Time: O(k^2) | Space: O(k)
// DP: dp[s] = ways to be at position s after i steps.
// Constraint: distance d = |startPos - endPos| must have same parity as k and d <= k.

import "fmt"

func main() {
	fmt.Println(numberOfWays(1, 2, 3)) // 3
	fmt.Println(numberOfWays(2, 5, 10)) // 0
	fmt.Println(numberOfWays(1, 1, 1)) // 0
}

const MOD = 1000000007

func numberOfWays(startPos int, endPos int, k int) int {
	d := abs(startPos - endPos)
	if d > k || (k-d)%2 != 0 {
		return 0
	}
	// DP with offset to handle negative positions
	offset := k
	size := 2*k + 1
  // Alokasi slice
	dp := make([]int, size)
	dp[0+offset] = 1

	for step := 0; step < k; step++ {
  // Alokasi slice
		next := make([]int, size)
		for pos := -k; pos <= k; pos++ {
			idx := pos + offset
			if dp[idx] == 0 {
				continue
			}
			// move left
			if pos-1 >= -k {
				next[pos-1+offset] = (next[pos-1+offset] + dp[idx]) % MOD
			}
			// move right
			if pos+1 <= k {
				next[pos+1+offset] = (next[pos+1+offset] + dp[idx]) % MOD
			}
		}
		dp = next
	}

	return dp[endPos-startPos+offset]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
