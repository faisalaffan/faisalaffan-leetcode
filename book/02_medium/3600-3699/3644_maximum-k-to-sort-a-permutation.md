# 3644 — Maximum K To Sort A Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumKToSortAPermutation(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3644: Maximum K to Sort a Permutation
// https://leetcode.com/problems/maximum-k-to-sort-a-permutation/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumKToSortAPermutation(nums []int) int {
	ans := -1 // all bits set to 1 (identity for AND)
	for i, x := range nums {
		if i != x {
			ans &= x
		}
	}
	if ans < 0 {
		return 0
	}
	return ans
}

func main() {
	fmt.Println(maximumKToSortAPermutation([]int{0, 3, 2, 1}))
	fmt.Println(maximumKToSortAPermutation([]int{3, 2, 1, 0}))
	fmt.Println(maximumKToSortAPermutation([]int{0, 1, 2, 3}))
}
```
