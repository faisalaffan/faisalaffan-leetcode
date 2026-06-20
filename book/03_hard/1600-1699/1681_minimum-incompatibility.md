# 1681 — Minimum Incompatibility

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumIncompatibility(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP, Bitmask

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1681: Minimum Incompatibility
// https://leetcode.com/problems/minimum-incompatibility/
// Difficulty: Hard
// Strategy: DP over bitmask. Precompute valid subsets of size n/k.

import (
	"fmt"
	"math"
)

func minimumIncompatibility(nums []int, k int) int {
	n := len(nums)
	subsetSize := n / k
	if subsetSize == 1 {
		return 0
	}

	// Precompute incompatibility for each valid subset of size subsetSize
  // Alokasi slice
	incomp := make([]int, 1<<n)
  // Range loop
	for i := range incomp {
		incomp[i] = -1
	}
	for mask := 1; mask < 1<<n; mask++ {
		if bitsCount(mask) != subsetSize {
			continue
		}
		// Check for duplicates and find min/max
  // HashMap: O(1) lookup
		seen := make(map[int]bool)
		minVal, maxVal := math.MaxInt32, math.MinInt32
		valid := true
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				if seen[nums[i]] {
					valid = false
					break
				}
				seen[nums[i]] = true
				if nums[i] < minVal {
					minVal = nums[i]
				}
				if nums[i] > maxVal {
					maxVal = nums[i]
				}
			}
		}
		if valid {
			incomp[mask] = maxVal - minVal
		}
	}

	// DP: dp[mask] = min incompatibility for chosen elements in mask
  // Alokasi slice
	dp := make([]int, 1<<n)
  // Range loop
	for i := range dp {
		dp[i] = -1
	}
	dp[0] = 0

	for mask := 0; mask < 1<<n; mask++ {
		if dp[mask] == -1 {
			continue
		}
		// Remaining elements
		remaining := ((1 << n) - 1) ^ mask
		if remaining == 0 {
			continue
		}
		// Pick a valid subset from remaining
		sub := remaining
		for sub > 0 {
			if incomp[sub] != -1 {
				newMask := mask | sub
				newVal := dp[mask] + incomp[sub]
				if dp[newMask] == -1 || newVal < dp[newMask] {
					dp[newMask] = newVal
				}
			}
			sub = (sub - 1) & remaining
		}
	}

	return dp[(1<<n)-1]
}

func bitsCount(x int) int {
	count := 0
	for x > 0 {
		count += x & 1
		x >>= 1
	}
	return count
}

func main() {
	// Example 1: [1,2,1,4], k=2 -> 4
	nums1 := []int{1, 2, 1, 4}
	k1 := 2
	fmt.Printf("minimumIncompatibility(%v, %d) = %d (expected 4)\n", nums1, k1, minimumIncompatibility(nums1, k1))

	// Example 2: [6,3,8,1,3,1,2,2], k=4 -> 6
	nums2 := []int{6, 3, 8, 1, 3, 1, 2, 2}
	k2 := 4
	fmt.Printf("minimumIncompatibility(%v, %d) = %d (expected 6)\n", nums2, k2, minimumIncompatibility(nums2, k2))

	// Example 3: [5,3,3,6,3,3], k=3 -> -1
	nums3 := []int{5, 3, 3, 6, 3, 3}
	k3 := 3
	fmt.Printf("minimumIncompatibility(%v, %d) = %d (expected -1)\n", nums3, k3, minimumIncompatibility(nums3, k3))

	// Single subset size
	nums4 := []int{1, 2, 3, 4}
	k4 := 4
	fmt.Printf("minimumIncompatibility(%v, %d) = %d (expected 0)\n", nums4, k4, minimumIncompatibility(nums4, k4))
}
```
