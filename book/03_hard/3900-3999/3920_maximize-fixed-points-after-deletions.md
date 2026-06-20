# 3920 — Maximize Fixed Points After Deletions

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxFixedPoints(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **DP** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3920: Maximize Fixed Points After Deletions
// https://leetcode.com/problems/maximize-fixed-points-after-deletions/
// Difficulty: Hard
//
// Given array nums, you can delete any number of elements. After
// deletions, an element is a "fixed point" if nums[i] == i in the
// resulting array. Maximize the number of fixed points.
//
// Approach: For each element, if nums[i] <= i, we can potentially
// make it a fixed point by deleting i - nums[i] elements before it.
// DP over elements tracking max fixed points achievable.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxFixedPoints([]int{0, 1, 2, 3}))
	// Example 2
	fmt.Println(maxFixedPoints([]int{0, 0, 2, 3}))
	// Edge: no matches
	fmt.Println(maxFixedPoints([]int{5, 5, 5}))
}

func maxFixedPoints(nums []int) int {
	n := len(nums)
	// dp[i] = max fixed points using first i+1 elements
  // Alokasi slice
	dp := make([]int, n+1)

	for i := 0; i < n; i++ {
		// Skip this element
		dp[i+1] = dp[i]
		// Try to make nums[i] a fixed point
		if nums[i] <= i {
			need := i - nums[i]
			candidates := dp[i-need] + 1
			if candidates > dp[i+1] {
				dp[i+1] = candidates
			}
		}
	}
	return dp[n]
}
```
