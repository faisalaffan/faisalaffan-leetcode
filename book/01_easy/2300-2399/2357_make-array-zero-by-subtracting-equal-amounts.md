# 2357 — Make Array Zero By Subtracting Equal Amounts

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MakeArrayZeroBySubtractingEqualAmounts(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2357: Make Array Zero by Subtracting Equal Amounts
// https://leetcode.com/problems/make-array-zero-by-subtracting-equal-amounts/
// Difficulty: Easy
// Time O(n log n) | Space O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MakeArrayZeroBySubtractingEqualAmounts([]int{1, 5, 0, 3, 5})) // 3
	fmt.Println(MakeArrayZeroBySubtractingEqualAmounts([]int{0}))              // 0
}

func MakeArrayZeroBySubtractingEqualAmounts(nums []int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	count := 0
  // Linear scan O(n)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			count++
			sub := nums[i]
			for j := i; j < len(nums); j++ {
				nums[j] -= sub
			}
		}
	}
	return count
}
```
