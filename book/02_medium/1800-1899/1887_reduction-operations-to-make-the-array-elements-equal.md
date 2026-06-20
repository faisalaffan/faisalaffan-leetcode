# 1887 — Reduction Operations To Make The Array Elements Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func ReductionOperations(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1887: Reduction Operations to Make the Array Elements Equal
// https://leetcode.com/problems/reduction-operations-to-make-the-array-elements-equal/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(ReductionOperations([]int{5, 1, 3}))
	fmt.Println(ReductionOperations([]int{1, 1, 1}))
	fmt.Println(ReductionOperations([]int{1, 1, 2, 2, 3}))
}

// Time: O(n log n), Space: O(1)
func ReductionOperations(nums []int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	n := len(nums)
	ops := 0
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			ops += n - i
		}
	}
	return ops
}
```
