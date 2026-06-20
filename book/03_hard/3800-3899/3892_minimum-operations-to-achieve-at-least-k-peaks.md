# 3892 — Minimum Operations To Achieve At Least K Peaks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func minOperations(nums []int, k int) int
```

> **💡 Hint:** For each position, compute the minimum increments

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Two Pointer, Dynamic Programming

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

> **Untuk fresh graduate:** Kuasai dulu teknik **Two Pointer** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3892: Minimum Operations to Achieve At Least K Peaks
// https://leetcode.com/problems/minimum-operations-to-achieve-at-least-k-peaks/
// Difficulty: Hard
//
// In a circular array, a peak is an element strictly greater than
// both neighbors. In one operation, increment any element by 1.
// Find minimum operations to achieve at least k peaks.
//
// Approach: For each position, compute the minimum increments
// needed to make it a peak (exceed both neighbors). Use DP to
// select k non-adjacent positions minimizing total cost.

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(minOperations([]int{1, 2, 3}, 1))
	// Example 2
	fmt.Println(minOperations([]int{1, 2, 1, 2, 1}, 2))
	// Edge: already has k peaks
	fmt.Println(minOperations([]int{5, 3, 5, 3, 5}, 2))
}

func minOperations(nums []int, k int) int {
	n := len(nums)
	if k > (n+1)/2 {
		return -1
	}

	// cost[i] = min increments to make position i a peak
  // Alokasi slice integer
	cost := make([]int, n)
	for i := 0; i < n; i++ {
		left := nums[(i-1+n)%n]
		right := nums[(i+1)%n]
		need := 0
		if nums[i] <= left {
			need = left - nums[i] + 1
		}
		if nums[i] <= right {
			add := right - nums[i] + 1
			if add > need {
				need = add
			}
		}
		cost[i] = need
	}

	// DP[i][j] = min cost for first i positions with j peaks (non-adjacent)
  // Membuat matriks/slice 2D untuk DP
	dp := make([][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range dp {
		dp[i] = make([]int, k+1)
		for j := range dp[i] {
			dp[i][j] = math.MaxInt32 / 2
		}
	}
	dp[0][0] = 0
	dp[1][0] = 0
	if n >= 1 {
		dp[1][1] = cost[0]
	}

	for i := 2; i <= n; i++ {
		for j := 0; j <= k; j++ {
			// Skip position i-1
			dp[i][j] = min(dp[i][j], dp[i-1][j])
			// Take position i-1 as peak (need i-2 to not be peak)
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i-2][j-1]+cost[i-1])
			}
		}
	}

	ans := dp[n][k]
	// Also consider circular wrap: position 0 and n-1 cannot both be peaks
	if k >= 2 {
		// Try excluding both ends from being peaks
  // Membuat matriks/slice 2D untuk DP
		dp2 := make([][]int, n)
  // Range loop: iterasi dengan indeks + nilai
		for i := range dp2 {
			dp2[i] = make([]int, k+1)
			for j := range dp2[i] {
				dp2[i][j] = math.MaxInt32 / 2
			}
		}
		dp2[0][0] = 0
		dp2[0][1] = cost[0]
		for i := 1; i < n; i++ {
			for j := 0; j <= k; j++ {
				dp2[i][j] = min(dp2[i][j], dp2[i-1][j])
				if j > 0 {
					prev := 0
					if i >= 2 {
						prev = dp2[i-2][j-1]
					} else if i == 1 && j == 1 {
						prev = 0
					}
					if prev < math.MaxInt32/2 {
						dp2[i][j] = min(dp2[i][j], prev+cost[i])
					}
				}
			}
		}
		ans = min(ans, dp2[n-1][k])
	}

	if ans >= math.MaxInt32/2 {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Keep compiler happy
var _ = sort.Ints
```
