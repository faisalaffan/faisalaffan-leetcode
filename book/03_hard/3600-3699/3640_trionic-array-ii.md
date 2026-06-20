# 3640 — Trionic Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxSumTrionic(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3640: Trionic Array II
// https://leetcode.com/problems/trionic-array-ii/
// Difficulty: Hard
//
// A trionic subarray is one that can be split into three parts: strictly
// increasing, strictly decreasing, then strictly increasing.
// Find the maximum sum of any trionic subarray.
//
// Approach: DP tracking states for each phase: increasing, decreasing, increasing.

import "fmt"
import "math"

func main() {
	// Example 1
	fmt.Println(maxSumTrionic([]int{1, 3, 2, 4, 5}))
	// Example 2
	fmt.Println(maxSumTrionic([]int{5, 4, 3, 2, 1}))
	// Example 3
	fmt.Println(maxSumTrionic([]int{1, 2, 3, 2, 1, 2, 3}))
	// Edge: single element
	fmt.Println(maxSumTrionic([]int{10}))
}

func maxSumTrionic(nums []int) int {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// dp[i][state] = max sum of a trionic subarray ending at i with given state
	// state 0: not started / part of increasing (phase 1)
	// state 1: in decreasing (phase 2)
	// state 2: in increasing (phase 3)
	// We need at least 3 elements to form a trionic subarray

  // Alokasi slice
	dp := make([][3]int, n)
  // Range loop
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}

	result := math.MinInt32

	for i := 0; i < n; i++ {
		// Start a new subarray at i (phase 1)
		dp[i][0] = nums[i]

		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				// Extend phase 1 or start phase 1
				if dp[j][0] != math.MinInt32 {
					dp[i][0] = max(dp[i][0], dp[j][0]+nums[i])
				}
				// Transition from phase 2 to phase 3 (increasing again)
				if dp[j][1] != math.MinInt32 {
					dp[i][2] = max(dp[i][2], dp[j][1]+nums[i])
				}
				// Extend phase 3
				if dp[j][2] != math.MinInt32 {
					dp[i][2] = max(dp[i][2], dp[j][2]+nums[i])
				}
			}
			if nums[i] < nums[j] {
				// Transition from phase 1 to phase 2 (decreasing)
				if dp[j][0] != math.MinInt32 {
					dp[i][1] = max(dp[i][1], dp[j][0]+nums[i])
				}
				// Extend phase 2
				if dp[j][1] != math.MinInt32 {
					dp[i][1] = max(dp[i][1], dp[j][1]+nums[i])
				}
			}
		}

		// Result must end in phase 3
		result = max(result, dp[i][2])
	}

	if result == math.MinInt32 {
		return 0
	}
	return result
}
```
