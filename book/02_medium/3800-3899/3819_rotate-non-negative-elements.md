# 3819 — Rotate Non Negative Elements

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func RotateNonNegativeElements(nums []int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(N)  |  **Ruang:** O(N)


## 💻 Solusi Go

```go
package main

// LeetCode #3819: Rotate Non Negative Elements
// https://leetcode.com/problems/rotate-non-negative-elements/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Collect non-negative elements, rotate left by k cyclically, place back.

import "fmt"

func RotateNonNegativeElements(nums []int, k int) []int {
	n := len(nums)
	type pair struct {
		idx int
		val int
	}
	var nonNeg []pair
	for i, v := range nums {
		if v >= 0 {
			nonNeg = append(nonNeg, pair{i, v})
		}
	}

	m := len(nonNeg)
	if m == 0 {
  // Alokasi slice
		res := make([]int, n)
		copy(res, nums)
		return res
	}

  // Alokasi slice
	result := make([]int, n)
	copy(result, nums)

	// For each position that had a non-negative, put the rotated value
	for i, p := range nonNeg {
		srcIdx := (i + k) % m
		result[p.idx] = nonNeg[srcIdx].val
	}

	return result
}

func main() {
	// Example 1
	fmt.Println(RotateNonNegativeElements([]int{1, -2, 3, -4}, 3)) // Expected: [3 -2 1 -4]

	// Example 2
	fmt.Println(RotateNonNegativeElements([]int{-3, -2, 7}, 1)) // Expected: [-3 -2 7]

	// Example 3
	fmt.Println(RotateNonNegativeElements([]int{5, 4, -9, 6}, 2)) // Expected: [6 5 -9 4]
}
```
