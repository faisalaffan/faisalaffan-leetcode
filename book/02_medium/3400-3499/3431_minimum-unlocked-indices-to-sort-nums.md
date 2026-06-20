# 3431 — Minimum Unlocked Indices To Sort Nums

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minUnlockedIndices(nums []int, locked []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n^2) Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3431: Minimum Unlocked Indices to Sort Nums
// https://leetcode.com/problems/minimum-unlocked-indices-to-sort-nums/
// Difficulty: Medium [Paid]
// Time: O(n^2) Space: O(n)

import (
	"fmt"
	"sort"
)

func minUnlockedIndices(nums []int, locked []int) int {
	n := len(nums)
  // Alokasi slice
	sorted := make([]int, n)
	copy(sorted, nums)
  // Sort O(n log n)
	sort.Ints(sorted)

	unlockCost := 0
	for i := 0; i < n; i++ {
		if nums[i] != sorted[i] && locked[i] == 1 {
			unlockCost++
		}
	}
	return unlockCost
}

func main() {
	fmt.Println(minUnlockedIndices([]int{1, 3, 2, 4}, []int{1, 1, 0, 1})) // 1
	fmt.Println(minUnlockedIndices([]int{1, 2, 3}, []int{1, 1, 1}))       // 0
}
```
