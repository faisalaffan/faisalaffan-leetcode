# 3674 — Minimum Operations To Equalize Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumOperationsToEqualizeArray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3674: Minimum Operations to Equalize Array
// https://leetcode.com/problems/minimum-operations-to-equalize-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(MinimumOperationsToEqualizeArray([]int{1, 2, 3}))
	fmt.Println(MinimumOperationsToEqualizeArray([]int{5, 5, 5}))
}

// Time: O(n)
// Space: O(1)
func MinimumOperationsToEqualizeArray(nums []int) int {
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[0] {
			return 1
		}
	}
	return 0
}
```
