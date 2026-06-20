# 1051 — Height Checker

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func heightChecker(heights []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1051: Height Checker
// https://leetcode.com/problems/height-checker/
// Difficulty: Easy
// Time: O(n log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(heightChecker([]int{1, 1, 4, 2, 1, 3})) // 3
	fmt.Println(heightChecker([]int{5, 1, 2, 3, 4}))    // 5
	fmt.Println(heightChecker([]int{1, 2, 3, 4, 5}))    // 0
}

// LeetCode submission: heightChecker
func heightChecker(heights []int) int {
  // Alokasi slice
	expected := make([]int, len(heights))
	copy(expected, heights)
  // Sort O(n log n)
	sort.Ints(expected)
	count := 0
  // Range loop
	for i := range heights {
		if heights[i] != expected[i] {
			count++
		}
	}
	return count
}
```
