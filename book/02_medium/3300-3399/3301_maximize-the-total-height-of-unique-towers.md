# 3301 — Maximize The Total Height Of Unique Towers

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumTotalSum(maximumHeight []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n) Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3301: Maximize the Total Height of Unique Towers
// https://leetcode.com/problems/maximize-the-total-height-of-unique-towers/
// Difficulty: Medium
// Time: O(n log n) Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumTotalSum([]int{2, 3, 4, 3}))   // 10
	fmt.Println(maximumTotalSum([]int{2, 2, 1}))      // -1
	fmt.Println(maximumTotalSum([]int{5, 4, 3, 2, 1})) // 15
}

func maximumTotalSum(maximumHeight []int) int64 {
  // Custom sort
	sort.Slice(maximumHeight, func(i, j int) bool {
		return maximumHeight[i] > maximumHeight[j]
	})

	var total int64
	prev := maximumHeight[0]
	total += int64(prev)

	for i := 1; i < len(maximumHeight); i++ {
		if prev <= 1 {
			return -1
		}
		h := maximumHeight[i]
		if h >= prev {
			h = prev - 1
		}
		total += int64(h)
		prev = h
	}
	return total
}
```
