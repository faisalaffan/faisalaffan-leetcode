# 1968 — Array With Elements Not Equal To Average Of Neighbors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func RearrangeArray(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Sorting

**Waktu:** O(n log n), Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1968: Array With Elements Not Equal to Average of Neighbors
// https://leetcode.com/problems/array-with-elements-not-equal-to-average-of-neighbors/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(RearrangeArray([]int{1, 2, 3, 4, 5}))
	fmt.Println(RearrangeArray([]int{6, 2, 0, 9, 7}))
}

// Time: O(n log n), Space: O(n)
func RearrangeArray(nums []int) []int {
  // Sort O(n log n)
	sort.Ints(nums)
	n := len(nums)
  // Alokasi slice
	result := make([]int, n)
	left, right := 0, n-1
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			result[i] = nums[left]
			left++
		} else {
			result[i] = nums[right]
			right--
		}
	}
	return result
}
```
