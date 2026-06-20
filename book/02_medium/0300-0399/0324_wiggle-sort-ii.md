# 0324 — Wiggle Sort Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func wiggleSort(nums []int) `

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #324: Wiggle Sort II
// https://leetcode.com/problems/wiggle-sort-ii/
// Difficulty: Medium
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func wiggleSort(nums []int) {
	n := len(nums)
  // Alokasi slice
	sorted := make([]int, n)
	copy(sorted, nums)
  // Sort O(n log n)
	sort.Ints(sorted)

	// Fill from end of sorted array into odd positions first, then even
	mid := (n + 1) / 2
	j, k := mid-1, n-1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			nums[i] = sorted[j]
			j--
		} else {
			nums[i] = sorted[k]
			k--
		}
	}
}

func main() {
	// Test case 1
	nums1 := []int{1, 5, 1, 1, 6, 4}
	wiggleSort(nums1)
	fmt.Println("Test 1:", nums1)

	// Test case 2
	nums2 := []int{1, 3, 2, 2, 3, 1}
	wiggleSort(nums2)
	fmt.Println("Test 2:", nums2)

	// Test case 3
	nums3 := []int{1, 2, 3}
	wiggleSort(nums3)
	fmt.Println("Test 3:", nums3)
}
```
