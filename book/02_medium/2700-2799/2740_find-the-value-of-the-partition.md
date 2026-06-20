# 2740 — Find The Value Of The Partition

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindTheValueOfThePartition(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2740: Find the Value of the Partition
// https://leetcode.com/problems/find-the-value-of-the-partition/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func FindTheValueOfThePartition(nums []int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	minDiff := nums[1] - nums[0]
	for i := 2; i < len(nums); i++ {
		if diff := nums[i] - nums[i-1]; diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}

func main() {
	fmt.Println(FindTheValueOfThePartition([]int{1, 3, 2, 4}))
	fmt.Println(FindTheValueOfThePartition([]int{100, 1, 10}))
}
```
