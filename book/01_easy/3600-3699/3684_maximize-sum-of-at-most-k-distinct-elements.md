# 3684 — Maximize Sum Of At Most K Distinct Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MaximizeSumOfAtMostKDistinctElements(nums []int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3684: Maximize Sum of At Most K Distinct Elements
// https://leetcode.com/problems/maximize-sum-of-at-most-k-distinct-elements/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MaximizeSumOfAtMostKDistinctElements([]int{84, 93, 100, 77, 90}, 3))
	fmt.Println(MaximizeSumOfAtMostKDistinctElements([]int{84, 93, 100, 77, 93}, 3))
	fmt.Println(MaximizeSumOfAtMostKDistinctElements([]int{1, 1, 1, 2, 2, 2}, 6))
}

// Time: O(n log n)
// Space: O(n)
func MaximizeSumOfAtMostKDistinctElements(nums []int, k int) []int {
  // HashMap: O(1) lookup
	seen := make(map[int]bool)
  // Alokasi slice
	unique := make([]int, 0)
	for _, v := range nums {
		if !seen[v] {
			seen[v] = true
			unique = append(unique, v)
		}
	}

  // Custom sort
	sort.Slice(unique, func(i, j int) bool {
		return unique[i] > unique[j]
	})

	size := k
	if size > len(unique) {
		size = len(unique)
	}
	return unique[:size]
}
```
