# 0462 — Minimum Moves To Equal Array Elements Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumMovesToEqualArrayElementsIi(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n) for sorting, O(n) for QuickSelect  |  **Ruang:** O(log n) for sorting, O(1) for QuickSelect

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #462: Minimum Moves to Equal Array Elements II
// https://leetcode.com/problems/minimum-moves-to-equal-array-elements-ii/
// Difficulty: Medium
// Time: O(n log n) for sorting, O(n) for QuickSelect
// Space: O(log n) for sorting, O(1) for QuickSelect

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinimumMovesToEqualArrayElementsIi([]int{1, 2, 3}))
	fmt.Println(MinimumMovesToEqualArrayElementsIi([]int{1, 10, 2, 9}))
	fmt.Println(MinimumMovesToEqualArrayElementsIi([]int{1, 0, 0, 8, 6}))
}

func MinimumMovesToEqualArrayElementsIi(nums []int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	median := nums[len(nums)/2]
	moves := 0
	for _, num := range nums {
		diff := num - median
		if diff < 0 {
			diff = -diff
		}
		moves += diff
	}
	return moves
}
```
