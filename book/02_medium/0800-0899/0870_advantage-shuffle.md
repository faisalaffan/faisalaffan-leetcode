# 0870 — Advantage Shuffle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func AdvantageShuffle(nums1 []int, nums2 []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #870: Advantage Shuffle
// https://leetcode.com/problems/advantage-shuffle/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(AdvantageShuffle([]int{2, 7, 11, 15}, []int{1, 10, 4, 11}))
	fmt.Println(AdvantageShuffle([]int{12, 24, 8, 32}, []int{13, 25, 32, 11}))
	fmt.Println(AdvantageShuffle([]int{2, 0, 4, 1, 2}, []int{1, 3, 0, 0, 2}))
}

// Time: O(n log n) | Space: O(n)
func AdvantageShuffle(nums1 []int, nums2 []int) []int {
	n := len(nums1)
  // Sort O(n log n)
	sort.Ints(nums1)

  // Alokasi slice
	idx := make([]int, n)
  // Range loop
	for i := range idx {
		idx[i] = i
	}
  // Custom sort
	sort.Slice(idx, func(i, j int) bool {
		return nums2[idx[i]] < nums2[idx[j]]
	})

  // Alokasi slice
	ans := make([]int, n)
	left, right := 0, n-1
	for _, x := range nums1 {
		if x > nums2[idx[left]] {
			ans[idx[left]] = x
			left++
		} else {
			ans[idx[right]] = x
			right--
		}
	}

	return ans
}
```
