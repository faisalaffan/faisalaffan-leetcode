# 2089 — Find Target Indices After Sorting Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindTargetIndicesAfterSortingArray(nums []int, target int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1) ignoring sort  |  **Ruang:** O(1) ignoring sort

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2089: Find Target Indices After Sorting Array
// https://leetcode.com/problems/find-target-indices-after-sorting-array/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(FindTargetIndicesAfterSortingArray([]int{1, 2, 5, 2, 3}, 2)) // [1 2]
	fmt.Println(FindTargetIndicesAfterSortingArray([]int{1, 2, 5, 2, 3}, 3)) // [3]
	fmt.Println(FindTargetIndicesAfterSortingArray([]int{1, 2, 5, 2, 3}, 5)) // [4]
}

// Time: O(n log n), Space: O(1) ignoring sort
func FindTargetIndicesAfterSortingArray(nums []int, target int) []int {
  // Sort O(n log n)
	sort.Ints(nums)
	var result []int
	for i, v := range nums {
		if v == target {
			result = append(result, i)
		}
	}
	return result
}
```
