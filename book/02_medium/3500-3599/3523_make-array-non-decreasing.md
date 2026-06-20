# 3523 — Make Array Non Decreasing

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MakeArrayNonDecreasing(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3523: Make Array Non-decreasing
// https://leetcode.com/problems/make-array-non-decreasing/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MakeArrayNonDecreasing([]int{4, 2, 3}))
	// Test case 2
	fmt.Println("Test 2:", MakeArrayNonDecreasing([]int{4, 2, 1}))
	// Test case 3
	fmt.Println("Test 3:", MakeArrayNonDecreasing([]int{1, 2, 3}))
}

func MakeArrayNonDecreasing(nums []int) int {
	ops := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			ops += nums[i-1] - nums[i]
			nums[i] = nums[i-1]
		}
	}
	return ops
}
```
