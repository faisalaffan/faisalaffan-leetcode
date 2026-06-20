# 3375 — Minimum Operations To Make Array Values Equal To K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumOperationsToMakeArrayValuesEqualToK(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n). Space: O(n).  |  **Ruang:** O(n).

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3375: Minimum Operations to Make Array Values Equal to K
// https://leetcode.com/problems/minimum-operations-to-make-array-values-equal-to-k/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToMakeArrayValuesEqualToK([]int{5, 2, 5, 4, 5}, 2))
	fmt.Println(MinimumOperationsToMakeArrayValuesEqualToK([]int{2, 1, 2}, 2))
	fmt.Println(MinimumOperationsToMakeArrayValuesEqualToK([]int{9, 7, 5, 3}, 1))
}

// MinimumOperationsToMakeArrayValuesEqualToK returns the minimum operations to reduce all numbers to k.
// In one operation, you can change any number > x to x for some x.
// Time: O(n). Space: O(n).
func MinimumOperationsToMakeArrayValuesEqualToK(nums []int, k int) int {
	minVal := nums[0]
	for _, v := range nums {
		if v < minVal {
			minVal = v
		}
	}
	if minVal < k {
		return -1
	}

  // HashMap: O(1) lookup
	seen := make(map[int]bool)
	for _, v := range nums {
		if v > k {
			seen[v] = true
		}
	}
	return len(seen)
}
```
