# 0910 — Smallest Range Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SmallestRangeIi(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(log n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #910: Smallest Range II
// https://leetcode.com/problems/smallest-range-ii/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SmallestRangeIi([]int{1}, 0))
	fmt.Println(SmallestRangeIi([]int{0, 10}, 2))
	fmt.Println(SmallestRangeIi([]int{1, 3, 6}, 3))
}

// Time: O(n log n) | Space: O(log n)
func SmallestRangeIi(nums []int, k int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	n := len(nums)
	ans := nums[n-1] - nums[0]

	for i := 1; i < n; i++ {
		mx := max(nums[i-1]+k, nums[n-1]-k)
		mn := min(nums[0]+k, nums[i]-k)
		ans = min(ans, mx-mn)
	}

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
```
