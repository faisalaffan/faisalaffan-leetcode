# 3634 — Minimum Removals To Balance Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MinimumRemovalsToBalanceArray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3634: Minimum Removals to Balance Array
// https://leetcode.com/problems/minimum-removals-to-balance-array/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MinimumRemovalsToBalanceArray([]int{1, 2, 3, 4, 5}))
	// Test case 2
	fmt.Println("Test 2:", MinimumRemovalsToBalanceArray([]int{1, 1, 1, 1}))
	// Test case 3
	fmt.Println("Test 3:", MinimumRemovalsToBalanceArray([]int{5, 4, 3, 2, 1}))
}

func MinimumRemovalsToBalanceArray(nums []int) int {
	// Minimum removals so that sum of first half equals sum of second half
	n := len(nums)
	if n%2 != 0 {
		// Remove middle element for odd length
		return 1
	}
	half := n / 2
	sum1, sum2 := 0, 0
	for i := 0; i < half; i++ {
		sum1 += nums[i]
	}
	for i := half; i < n; i++ {
		sum2 += nums[i]
	}
	if sum1 == sum2 {
		return 0
	}
	// Remove one element from the larger half
	return 1
}
```
