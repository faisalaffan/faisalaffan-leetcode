# 2279 — Maximum Bags With Full Capacity Of Rocks

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumBags(capacity []int, rocks []int, additionalRocks int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2279: Maximum Bags With Full Capacity of Rocks
// https://leetcode.com/problems/maximum-bags-with-full-capacity-of-rocks/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	n := len(capacity)
  // Alokasi slice
	needed := make([]int, n)
	for i := 0; i < n; i++ {
		needed[i] = capacity[i] - rocks[i]
	}
  // Sort O(n log n)
	sort.Ints(needed)

	full := 0
	for _, n := range needed {
  // Edge case: input kosong
		if n == 0 {
			full++
		} else if n <= additionalRocks {
			additionalRocks -= n
			full++
		} else {
			break
		}
	}
	return full
}

func main() {
	// Test case 1
	fmt.Println(maximumBags([]int{2, 3, 4, 5}, []int{1, 2, 4, 4}, 2))
	// Expected: 3

	// Test case 2
	fmt.Println(maximumBags([]int{10, 2, 2}, []int{2, 2, 0}, 100))
	// Expected: 3
}
```
