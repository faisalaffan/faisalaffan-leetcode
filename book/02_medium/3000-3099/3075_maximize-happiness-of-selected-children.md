# 3075 — Maximize Happiness Of Selected Children

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumHappinessSum(happiness []int, k int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3075: Maximize Happiness of Selected Children
// https://leetcode.com/problems/maximize-happiness-of-selected-children/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumHappinessSum([]int{1, 2, 3}, 2))
	fmt.Println(maximumHappinessSum([]int{1, 1, 1, 1}, 2))
	fmt.Println(maximumHappinessSum([]int{2, 3, 4, 5}, 1))
}

func maximumHappinessSum(happiness []int, k int) int64 {
  // Custom sort
	sort.Slice(happiness, func(i, j int) bool {
		return happiness[i] > happiness[j]
	})
	ans := int64(0)
	for i := 0; i < k; i++ {
		val := happiness[i] - i
		if val > 0 {
			ans += int64(val)
		}
	}
	return ans
}
```
