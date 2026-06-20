# 3957 — Maximum Sum Of M Non Overlapping Subarrays Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func maxSum(nums []int, m int, l int, r int) int
```

> **💡 Hint:** DP with prefix sums and gap constraint. dp[i][j] =

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming, Prefix Sum

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3957: Maximum Sum of M Non-Overlapping Subarrays II
// https://leetcode.com/problems/maximum-sum-of-m-non-overlapping-subarrays-ii/
// Difficulty: Hard
//
// Select exactly m non-overlapping subarrays from nums, each with
// length in [l, r]. Maximize total sum. This version adds a
// constraint that subarrays cannot be adjacent (at least one
// element gap).
//
// Approach: DP with prefix sums and gap constraint. dp[i][j] =
// max sum using first i elements, j subarrays. Gap enforced by
// requiring i-k-1 gap before next subarray.

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(maxSum([]int{1, 2, 3, 4, 5}, 2, 1, 2))
	// Example 2
	fmt.Println(maxSum([]int{10, -1, -2, 10}, 2, 1, 2))
	// Edge: single subarray
	fmt.Println(maxSum([]int{5, 3, 7}, 1, 1, 3))
}

func maxSum(nums []int, m int, l int, r int) int {
	n := len(nums)
  // Alokasi slice integer
	pref := make([]int, n+1)
	for i := 0; i < n; i++ {
		pref[i+1] = pref[i] + nums[i]
	}

	// dp[i][j] = max sum using first i elements, j subarrays
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32 / 2
		}
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = 0
	}

	for j := 1; j <= m; j++ {
		for i := 1; i <= n; i++ {
			// Skip nums[i-1]
			dp[i][j] = dp[i-1][j]
			// Take subarray ending at i-1 with gap of at least 1
			for k := l; k <= r && k <= i; k++ {
				prevEnd := i - k
				gapStart := 0
				if prevEnd > 0 {
					gapStart = prevEnd - 1
				}
				if prevEnd > 0 {
					// Need at least 1 gap before this subarray
					if dp[gapStart][j-1] != math.MinInt32/2 {
						sum := pref[i] - pref[i-k]
						if dp[gapStart][j-1]+sum > dp[i][j] {
							dp[i][j] = dp[gapStart][j-1] + sum
						}
					}
				} else if j == 1 {
					// First subarray, no gap needed
					sum := pref[i] - pref[i-k]
					if sum > dp[i][j] {
						dp[i][j] = sum
					}
				}
			}
		}
	}

	if dp[n][m] < 0 {
		return 0
	}
	return dp[n][m]
}
```
