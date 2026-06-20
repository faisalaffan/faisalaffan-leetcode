# 0090 — Subsets Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func subsetsWithDup(nums []int) [][]int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n * 2^n)  |  **Ruang:** O(n * 2^n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #90: Subsets II
// https://leetcode.com/problems/subsets-ii/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
  // Sort O(n log n)
	sort.Ints(nums)
	result := [][]int{{}}
	start := 0

  // Linear scan O(n)
	for i := 0; i < len(nums); i++ {
		n := len(result)
		begin := 0
		if i > 0 && nums[i] == nums[i-1] {
			begin = start
		}
		start = n
		for j := begin; j < n; j++ {
  // Alokasi slice
			newSubset := make([]int, len(result[j])+1)
			copy(newSubset, result[j])
			newSubset[len(result[j])] = nums[i]
			result = append(result, newSubset)
		}
	}

	return result
}

func main() {
	// Test case 1
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
	// [[] [1] [2] [1 2] [2 2] [1 2 2]]

	// Test case 2
	fmt.Println(subsetsWithDup([]int{0})) // [[] [0]]
}

// Time: O(n * 2^n) | Space: O(n * 2^n)
```
