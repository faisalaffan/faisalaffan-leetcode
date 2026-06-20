# 3065 — Minimum Operations To Exceed Threshold Value I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumOperationsToExceedThresholdValueI(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3065: Minimum Operations to Exceed Threshold Value I
// https://leetcode.com/problems/minimum-operations-to-exceed-threshold-value-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: minOperations
	fmt.Println(MinimumOperationsToExceedThresholdValueI([]int{2, 11, 10, 1, 3}, 10)) // 3
	fmt.Println(MinimumOperationsToExceedThresholdValueI([]int{1, 1, 2, 4, 9}, 9))    // 4
}

// Time: O(n) | Space: O(1)
// LeetCode submission name: minOperations
func MinimumOperationsToExceedThresholdValueI(nums []int, k int) int {
	count := 0
	for _, v := range nums {
		if v < k {
			count++
		}
	}
	return count
}
```
