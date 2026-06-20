# 1846 — Maximum Element After Decreasing And Rearranging

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaximumElementAfterDecreasingAndRearranging(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n), Space: O(1) (ignoring sort space)  |  **Ruang:** O(1) (ignoring sort space)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1846: Maximum Element After Decreasing and Rearranging
// https://leetcode.com/problems/maximum-element-after-decreasing-and-rearranging/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximumElementAfterDecreasingAndRearranging([]int{2, 2, 1, 2, 1}))
	fmt.Println(MaximumElementAfterDecreasingAndRearranging([]int{100, 1, 1000}))
	fmt.Println(MaximumElementAfterDecreasingAndRearranging([]int{1, 2, 3, 4, 5}))
}

// Time: O(n log n), Space: O(1) (ignoring sort space)
func MaximumElementAfterDecreasingAndRearranging(arr []int) int {
  // Sort O(n log n)
	sort.Ints(arr)
	arr[0] = 1
	for i := 1; i < len(arr); i++ {
		if arr[i]-arr[i-1] > 1 {
			arr[i] = arr[i-1] + 1
		}
	}
	return arr[len(arr)-1]
}
```
