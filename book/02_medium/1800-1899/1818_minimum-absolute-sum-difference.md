# 1818 — Minimum Absolute Sum Difference

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minAbsoluteSumDiff(nums1 []int, nums2 []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1818: Minimum Absolute Sum Difference
// https://leetcode.com/problems/minimum-absolute-sum-difference/
// Difficulty: Medium
// Time: O(n log n), Space: O(n)

import (
	"fmt"
	"sort"
)

const mod = 1_000_000_007

func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	n := len(nums1)
  // Alokasi slice
	sorted := make([]int, n)
	copy(sorted, nums1)
  // Sort O(n log n)
	sort.Ints(sorted)

	total := 0
	maxReduction := 0

	for i := 0; i < n; i++ {
		origDiff := abs(nums1[i] - nums2[i])
		total = (total + origDiff) % mod

		// Find closest value to nums2[i] in sorted nums1
		idx := sort.SearchInts(sorted, nums2[i])
		if idx < n {
			maxReduction = max(maxReduction, origDiff-abs(sorted[idx]-nums2[i]))
		}
		if idx > 0 {
			maxReduction = max(maxReduction, origDiff-abs(sorted[idx-1]-nums2[i]))
		}
	}

	return (total - maxReduction + mod) % mod
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minAbsoluteSumDiff([]int{1, 7, 5}, []int{2, 3, 5})) // Expected: 3
	fmt.Println(minAbsoluteSumDiff([]int{2, 4, 6, 8, 10}, []int{2, 4, 6, 8, 10})) // Expected: 0
	fmt.Println(minAbsoluteSumDiff([]int{1, 10, 4, 4, 2, 7}, []int{9, 3, 5, 1, 7, 4})) // Expected: 20
}
```
