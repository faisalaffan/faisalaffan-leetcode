# 0801 — Minimum Swaps To Make Sequences Increasing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minSwap(nums1 []int, nums2 []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #801: Minimum Swaps To Make Sequences Increasing
// https://leetcode.com/problems/minimum-swaps-to-make-sequences-increasing/
// Difficulty: Hard
//
// Given two integer arrays nums1 and nums2 of the same length, you may swap
// nums1[i] and nums2[i] at any index. Find the minimum number of swaps so that
// both sequences are strictly increasing.
//
// Approach: DP with two states at each position:
//   - dp0: min swaps for prefix up to i WITHOUT swapping at i
//   - dp1: min swaps for prefix up to i WITH swapping at i
// Transition depends on whether the natural order and/or cross order is
// strictly increasing.

import (
	"fmt"
	"math"
)

func main() {
	// Example: [1,3,5,4], [1,2,3,7] → 1
	fmt.Println(minSwap([]int{1, 3, 5, 4}, []int{1, 2, 3, 7}))
	// Already increasing: 0
	fmt.Println(minSwap([]int{1, 2, 3}, []int{4, 5, 6}))
	// Need swap at every position: [1,2,3],[3,2,1] → ?
	fmt.Println(minSwap([]int{1, 2, 3}, []int{3, 2, 1}))
	// Single element: 0
	fmt.Println(minSwap([]int{1}, []int{2}))
	// Two elements, need one swap
	fmt.Println(minSwap([]int{1, 4}, []int{2, 3}))
}

func minSwap(nums1 []int, nums2 []int) int {
	n := len(nums1)
	// dp0 = min swaps for first element without swapping
	// dp1 = min swaps for first element with swapping
	dp0, dp1 := 0, 1

	for i := 1; i < n; i++ {
		ndp0, ndp1 := math.MaxInt32, math.MaxInt32

		// Both sequences are naturally increasing (no swap at i, no swap at i-1)
		if nums1[i] > nums1[i-1] && nums2[i] > nums2[i-1] {
			ndp0 = min(ndp0, dp0)
			ndp1 = min(ndp1, dp1+1)
		}

		// Crossing works: swapping i but not i-1, or vice versa
		if nums1[i] > nums2[i-1] && nums2[i] > nums1[i-1] {
			ndp0 = min(ndp0, dp1)
			ndp1 = min(ndp1, dp0+1)
		}

		dp0, dp1 = ndp0, ndp1
	}

	return min(dp0, dp1)
}
```
