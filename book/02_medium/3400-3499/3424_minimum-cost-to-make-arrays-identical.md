# 3424 — Minimum Cost To Make Arrays Identical

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minCost(arr []int, brr []int, k int64) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n) Space: O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3424: Minimum Cost to Make Arrays Identical
// https://leetcode.com/problems/minimum-cost-to-make-arrays-identical/
// Difficulty: Medium
// Time: O(n log n) Space: O(n)

import (
	"fmt"
	"sort"
)

func minCost(arr []int, brr []int, k int64) int64 {
	var cost1 int64
  // Linear scan O(n)
	for i := 0; i < len(arr); i++ {
		diff := arr[i] - brr[i]
		if diff < 0 {
			diff = -diff
		}
		cost1 += int64(diff)
	}

  // Alokasi slice
	sortedArr := make([]int, len(arr))
  // Alokasi slice
	sortedBrr := make([]int, len(brr))
	copy(sortedArr, arr)
	copy(sortedBrr, brr)
  // Sort O(n log n)
	sort.Ints(sortedArr)
  // Sort O(n log n)
	sort.Ints(sortedBrr)

	var cost2 int64
  // Linear scan O(n)
	for i := 0; i < len(sortedArr); i++ {
		diff := sortedArr[i] - sortedBrr[i]
		if diff < 0 {
			diff = -diff
		}
		cost2 += int64(diff)
	}
	cost2 += k

	if cost1 < cost2 {
		return cost1
	}
	return cost2
}

func main() {
	fmt.Println(minCost([]int{4, 2, 5}, []int{6, 3, 1}, 2)) // 7
	fmt.Println(minCost([]int{1, 2, 3}, []int{4, 5, 6}, 1)) // 9
}
```
