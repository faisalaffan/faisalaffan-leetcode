# 3107 — Minimum Operations To Make Median Of Array Equal To K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minOperationsToMakeMedianK(nums []int, k int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(log n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3107: Minimum Operations to Make Median of Array Equal to K
// https://leetcode.com/problems/minimum-operations-to-make-median-of-array-equal-to-k/
// Difficulty: Medium
// Time: O(n log n) | Space: O(log n)

import (
	"fmt"
	"sort"
)

func minOperationsToMakeMedianK(nums []int, k int) int64 {
  // Sort O(n log n)
	sort.Ints(nums)
	n := len(nums)
	mid := n / 2
	var ans int64

	ans += abs(int64(nums[mid] - k))
	nums[mid] = k

	for i := mid - 1; i >= 0 && nums[i] > k; i-- {
		ans += abs(int64(nums[i] - k))
		nums[i] = k
	}

	for i := mid + 1; i < n && nums[i] < k; i++ {
		ans += abs(int64(nums[i] - k))
		nums[i] = k
	}

	return ans
}

func abs(x int64) int64 {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	fmt.Println(minOperationsToMakeMedianK([]int{2, 5, 6, 8, 5}, 4)) // Expected: 2
	fmt.Println(minOperationsToMakeMedianK([]int{2, 5, 6, 8, 5}, 7)) // Expected: 3
	fmt.Println(minOperationsToMakeMedianK([]int{1, 2, 3, 4, 5, 6}, 4)) // Expected: 0
}
```
