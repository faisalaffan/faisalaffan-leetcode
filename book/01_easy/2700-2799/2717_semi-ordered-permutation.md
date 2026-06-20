# 2717 — Semi Ordered Permutation

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SemiOrderedPermutation(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2717: Semi-Ordered Permutation
// https://leetcode.com/problems/semi-ordered-permutation/
// Difficulty: Easy
// Time: O(n) | Space: O(1)

import "fmt"

func main() {
	fmt.Println(SemiOrderedPermutation([]int{2, 1, 4, 3}))
	fmt.Println(SemiOrderedPermutation([]int{2, 4, 1, 3}))
}

func SemiOrderedPermutation(nums []int) int {
	n := len(nums)
	pos1, posN := 0, 0
	for i, v := range nums {
		if v == 1 {
			pos1 = i
		}
		if v == n {
			posN = i
		}
	}

	swaps := pos1 + (n - 1 - posN)
	if pos1 > posN {
		swaps--
	}
	return swaps
}
```
