# 3891 — Minimum Increase To Maximize Special Indices

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumIncreaseToMaximizeSpecialIndices(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** O(N)  |  **Ruang:** O(N)

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

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
  // Alokasi slice
	cost := make([]int, n)
	for i := 1; i < n-1; i++ {
		need := max(nums[i-1], nums[i+1]) + 1
		if nums[i] < need {
			cost[i] = need - nums[i]
		}
	}

	maxPeaks := (n - 1) / 2
  // Matriks 2D
	memo := make([][]int, n+1)
  // Range loop
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
