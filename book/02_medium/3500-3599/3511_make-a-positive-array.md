# 3511 — Make A Positive Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func MakeAPositiveArray(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3511: Make a Positive Array
// https://leetcode.com/problems/make-a-positive-array/
// Difficulty: Medium [Paid]
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MakeAPositiveArray([]int{-1, 2, -3, 4}))
	// Test case 2
	fmt.Println("Test 2:", MakeAPositiveArray([]int{1, 2, 3}))
	// Test case 3
	fmt.Println("Test 3:", MakeAPositiveArray([]int{-5, -10}))
}

func MakeAPositiveArray(nums []int) int {
	// Minimum operations to make all elements positive
	// Each operation can increment an element by 1
	ops := 0
	for _, v := range nums {
		if v <= 0 {
			ops += -v + 1
		}
	}
	return ops
}
```
