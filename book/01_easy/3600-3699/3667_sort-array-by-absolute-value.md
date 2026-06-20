# 3667 — Sort Array By Absolute Value

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SortArrayByAbsoluteValue(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3667: Sort Array By Absolute Value
// https://leetcode.com/problems/sort-array-by-absolute-value/
// Difficulty: Easy [Paid]

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortArrayByAbsoluteValue([]int{3, -1, -4, 1, 5}))
}

// Time: O(n log n)
// Space: O(n)
func SortArrayByAbsoluteValue(nums []int) []int {
  // Custom sort
	sort.Slice(nums, func(i, j int) bool {
		return abs(nums[i]) < abs(nums[j])
	})
	return nums
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
