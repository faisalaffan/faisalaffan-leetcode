# 1608 — Special Array With X Elements Greater Than Or Equal X

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SpecialArray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(log n) (for sorting)  |  **Ruang:** O(log n) (for sorting)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1608: Special Array With X Elements Greater Than or Equal X
// https://leetcode.com/problems/special-array-with-x-elements-greater-than-or-equal-x/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

// Time: O(n log n), Space: O(log n) (for sorting)
func SpecialArray(nums []int) int {
  // Custom sort
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

  // Linear scan O(n)
	for i := 0; i < len(nums); i++ {
		if nums[i] >= i+1 {
			if i == len(nums)-1 || nums[i+1] < i+1 {
				return i + 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(SpecialArray([]int{3, 5}))
	fmt.Println(SpecialArray([]int{0, 0}))
	fmt.Println(SpecialArray([]int{0, 4, 3, 0, 4}))
}
```
