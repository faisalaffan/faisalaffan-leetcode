# 1262 — Greatest Sum Divisible By Three

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sekumpulan bilangan dan diminta untuk menghitung penjumlahan dengan aturan tertentu. Tugasmu adalah menemukan kombinasi, subset, atau urutan yang memenuhi target penjumlahan.

Seperti menghitung kembalian belanja — kamu perlu kombinasi pecahan uang yang tepat. Soal penjumlahan seringnya diselesaikan dengan HashMap (two-sum pattern) atau Prefix Sum (jumlah kumulatif).

**Konsep kunci:** target sum, complement (pelengkap), prefix sum, cumulative sum.

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSumDivThree(nums []int) int
```

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

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
