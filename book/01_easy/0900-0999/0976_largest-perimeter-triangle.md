# 0976 — Largest Perimeter Triangle

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func largestPerimeter(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n). Space: O(log n).  |  **Ruang:** O(log n).

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #976: Largest Perimeter Triangle
// https://leetcode.com/problems/largest-perimeter-triangle/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(largestPerimeter([]int{2, 1, 2}))          // 5
	fmt.Println(largestPerimeter([]int{1, 2, 1}))          // 0
	fmt.Println(largestPerimeter([]int{3, 6, 2, 3}))       // 8
}

// largestPerimeter finds the largest perimeter of a triangle from the given side lengths.
// Time: O(n log n). Space: O(log n).
func largestPerimeter(nums []int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	for i := len(nums) - 1; i >= 2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}
	return 0
}
```
