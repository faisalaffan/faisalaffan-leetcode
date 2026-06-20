# 0413 — Arithmetic Slices

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func numberOfArithmeticSlices(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #413: Arithmetic Slices
// https://leetcode.com/problems/arithmetic-slices/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func numberOfArithmeticSlices(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	total := 0
	curr := 0

	for i := 2; i < len(nums); i++ {
		if nums[i]-nums[i-1] == nums[i-1]-nums[i-2] {
			curr++
			total += curr
		} else {
			curr = 0
		}
	}
	return total
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", numberOfArithmeticSlices([]int{1, 2, 3, 4}))
	// Expected: 3

	// Test case 2
	fmt.Println("Test 2:", numberOfArithmeticSlices([]int{1}))
	// Expected: 0

	// Test case 3
	fmt.Println("Test 3:", numberOfArithmeticSlices([]int{1, 2, 3, 8, 9, 10}))
	// Expected: 2
}
```
