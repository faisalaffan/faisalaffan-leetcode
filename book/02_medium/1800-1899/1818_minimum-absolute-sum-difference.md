# 1818 — Minimum Absolute Sum Difference

## Deskripsi

**Soal:** [1818. Minimum Absolute Sum Difference](https://leetcode.com/problems/minimum-absolute-sum-difference/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func minAbsoluteSumDiff(nums1 []int, nums2 []int) int`

## Solusi Go

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
  // Membuat slice untuk menyimpan hasil
	sorted := make([]int, n)
	copy(sorted, nums1)
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
