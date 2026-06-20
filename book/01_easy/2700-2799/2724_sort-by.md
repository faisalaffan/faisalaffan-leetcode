# 2724 — Sort By

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SortBy(arr []int, fn func(int) int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2724: Sort By
// https://leetcode.com/problems/sort-by/
// Difficulty: Easy
// Time: O(n log n) | Space: O(1)
// Note: JavaScript problem, adapted to Go.

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fn := func(n int) int { return n % 2 }
	fmt.Println(SortBy(arr, fn))

	arr2 := []int{1, 2, 3, 4, 5}
	fn2 := func(n int) int { return n }
	fmt.Println(SortBy(arr2, fn2))
}

func SortBy(arr []int, fn func(int) int) []int {
	sort.SliceStable(arr, func(i, j int) bool {
		return fn(arr[i]) < fn(arr[j])
	})
	return arr
}
```
