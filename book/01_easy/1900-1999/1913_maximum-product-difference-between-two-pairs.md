# 1913 — Maximum Product Difference Between Two Pairs

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaximumProductDifferenceBetweenTwoPairs(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1) ignoring sort  |  **Ruang:** O(1) ignoring sort

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1913: Maximum Product Difference Between Two Pairs
// https://leetcode.com/problems/maximum-product-difference-between-two-pairs/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximumProductDifferenceBetweenTwoPairs([]int{5, 6, 2, 7, 4}))      // 34
	fmt.Println(MaximumProductDifferenceBetweenTwoPairs([]int{4, 2, 5, 9, 7, 4, 8})) // 64
}

// Time: O(n log n), Space: O(1) ignoring sort
func MaximumProductDifferenceBetweenTwoPairs(nums []int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	n := len(nums)
	return nums[n-1]*nums[n-2] - nums[0]*nums[1]
}
```
