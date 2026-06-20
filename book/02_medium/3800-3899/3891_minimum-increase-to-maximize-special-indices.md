# 3891 — Minimum Increase To Maximize Special Indices

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan sebuah array (larik) bilangan. Tugasmu adalah mencari elemen atau pola tertentu dalam array tersebut, lalu mengembalikan hasilnya sesuai permintaan soal.

Bayangkan kamu sedang memeriksa daftar nilai ujian — kamu perlu menemukan nilai tertentu atau menghitung sesuatu dari daftar tersebut. Array adalah struktur data paling dasar: kumpulan elemen yang disimpan berurutan di memori.

**Konsep kunci:** indeks (posisi), value (nilai), panjang array (len).

**Fungsi yang perlu kamu implementasikan:**
```go
func MinimumIncreaseToMaximizeSpecialIndices(nums []int) int
```

> **💡 Hint:** DP with memoization. Max special indices = ceil((n-2)/2) = (n-1)/2.

## 🔍 Petunjuk Penyelesaian

**Teknik yang digunakan:** Dynamic Programming

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(N)

> **Untuk fresh graduate:** Kuasai dulu teknik **Dynamic Programming** sebelum lanjut ke solusi. Teknik ini sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3891: Minimum Increase to Maximize Special Indices
// https://leetcode.com/problems/minimum-increase-to-maximize-special-indices/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: DP with memoization. Max special indices = ceil((n-2)/2) = (n-1)/2.
// Find min cost by considering each position as peak or not.

import (
	"fmt"
	"math"
)

func MinimumIncreaseToMaximizeSpecialIndices(nums []int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	// cost to make each position a peak
  // Alokasi slice integer
	cost := make([]int, n)
	for i := 1; i < n-1; i++ {
		need := max(nums[i-1], nums[i+1]) + 1
		if nums[i] < need {
			cost[i] = need - nums[i]
		}
	}

	maxPeaks := (n - 1) / 2
  // Membuat matriks/slice 2D untuk DP
	memo := make([][]int, n+1)
  // Range loop: iterasi dengan indeks + nilai
	for i := range memo {
		memo[i] = make([]int, maxPeaks+1)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	// dp(pos, peaksSelected) = min cost from pos to n-2 (pos starts at 1)
	var dp func(pos int, left int) int
	dp = func(pos int, left int) int {
		if left == 0 {
			return 0
		}
		if pos >= n-1 {
			if left > 0 {
				return math.MaxInt32
			}
			return 0
		}
		if memo[pos][left] != -1 {
			return memo[pos][left]
		}

		// Skip position pos
		best := dp(pos+1, left)

		// Make pos a peak (skip pos+1 since adjacent can't be peak)
		best = min(best, cost[pos]+dp(pos+2, left-1))

		memo[pos][left] = best
		return best
	}

	ans := dp(1, maxPeaks)
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Example 1
	fmt.Println(MinimumIncreaseToMaximizeSpecialIndices([]int{1, 2, 2})) // Expected: 1

	// Example 2
	fmt.Println(MinimumIncreaseToMaximizeSpecialIndices([]int{2, 1, 1, 3})) // Expected: 2

	// Example 3
	fmt.Println(MinimumIncreaseToMaximizeSpecialIndices([]int{5, 2, 1, 4, 3})) // Expected: 4
}
```
