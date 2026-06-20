# 2967 — Minimum Cost To Make Array Equalindromic

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumCostEqualindromic(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting

**Waktu:** O(n log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2967: Minimum Cost to Make Array Equalindromic
// https://leetcode.com/problems/minimum-cost-to-make-array-equalindromic/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(minimumCostEqualindromic([]int{1, 2, 3, 4, 5}))
	fmt.Println(minimumCostEqualindromic([]int{10, 12, 13, 14, 15}))
}

func minimumCostEqualindromic(nums []int) int64 {
  // Sort O(n log n)
	sort.Ints(nums)
	median := nums[len(nums)/2]

	isPal := func(x int) bool {
		s := fmt.Sprintf("%d", x)
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			if s[i] != s[j] {
				return false
			}
		}
		return true
	}

	nextPal := median
	for !isPal(nextPal) {
		nextPal++
	}
	prevPal := median
	for !isPal(prevPal) {
		prevPal--
	}

	cost1, cost2 := int64(0), int64(0)
	for _, num := range nums {
		cost1 += int64(abs(num - prevPal))
		cost2 += int64(abs(num - nextPal))
	}
	if cost1 < cost2 {
		return cost1
	}
	return cost2
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```
