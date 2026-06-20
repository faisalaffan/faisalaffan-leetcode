# 1403 — Minimum Subsequence In Non Increasing Order

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minSubsequence(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1) excluding output  |  **Ruang:** O(1) excluding output

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1403: Minimum Subsequence in Non-Increasing Order
// https://leetcode.com/problems/minimum-subsequence-in-non-increasing-order/
// Difficulty: Easy
//
// LeetCode submission: func minSubsequence(nums []int) []int

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumSubsequenceInNonIncreasingOrder([]int{4, 3, 10, 9, 8})) // [10 9]
	fmt.Println(MinimumSubsequenceInNonIncreasingOrder([]int{4, 4, 7, 6, 7}))  // [7 7 6]
}

// Time: O(n log n), Space: O(1) excluding output
func MinimumSubsequenceInNonIncreasingOrder(nums []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	total := 0
	for _, v := range nums {
		total += v
	}
	sum := 0
	for i, v := range nums {
		sum += v
		if sum > total-sum {
			return nums[:i+1]
		}
	}
	return nums
}
```
