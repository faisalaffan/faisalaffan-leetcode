# 1877 — Minimize Maximum Pair Sum In Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinPairSum(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1877: Minimize Maximum Pair Sum in Array
// https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinPairSum([]int{3, 5, 2, 3}))
	fmt.Println(MinPairSum([]int{3, 5, 4, 2, 4, 6}))
}

// Time: O(n log n), Space: O(1)
func MinPairSum(nums []int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	maxSum := 0
	n := len(nums)
	for i := 0; i < n/2; i++ {
		pairSum := nums[i] + nums[n-1-i]
		if pairSum > maxSum {
			maxSum = pairSum
		}
	}
	return maxSum
}
```
