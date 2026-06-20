# 1402 — Reducing Dishes

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxSatisfaction(satisfaction []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1402: Reducing Dishes
// https://leetcode.com/problems/reducing-dishes/
// Difficulty: Hard

import (
	"fmt"
	"sort"
)

func maxSatisfaction(satisfaction []int) int {
	// Sort descending
  // Custom sort
	sort.Slice(satisfaction, func(i, j int) bool {
		return satisfaction[i] > satisfaction[j]
	})

	curr := 0
	maxVal := 0

	for _, s := range satisfaction {
		curr += s
		if curr > 0 {
			maxVal += curr
		}
	}

	return maxVal
}

func main() {
	// Example 1
	fmt.Println(maxSatisfaction([]int{-1, -8, 0, 5, -9}))
	// Expected: 14

	// Example 2
	fmt.Println(maxSatisfaction([]int{4, 3, 2}))
	// Expected: 20

	// Example 3
	fmt.Println(maxSatisfaction([]int{-1, -4, -5}))
	// Expected: 0
}
```
