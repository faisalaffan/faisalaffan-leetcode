# 1262 — Greatest Sum Divisible By Three

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxSumDivThree(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1262: Greatest Sum Divisible by Three
// https://leetcode.com/problems/greatest-sum-divisible-by-three/
// Difficulty: Medium

// DP with 3 states: max sum with remainder 0, 1, 2.

// Time: O(n)
// Space: O(1)

func maxSumDivThree(nums []int) int {
	dp := [3]int{0, -1, -1}

	for _, v := range nums {
		next := dp
		for r := 0; r < 3; r++ {
			if dp[r] != -1 {
				nr := (r + v) % 3
				if dp[r]+v > next[nr] {
					next[nr] = dp[r] + v
				}
			}
		}
		dp = next
	}

	return dp[0]
}

func main() {
	fmt.Printf("%d (expected: 18)\n", maxSumDivThree([]int{3, 6, 5, 1, 8}))
	fmt.Printf("%d (expected: 0)\n", maxSumDivThree([]int{4}))
	fmt.Printf("%d (expected: 12)\n", maxSumDivThree([]int{1, 2, 3, 4, 4}))
}
```
