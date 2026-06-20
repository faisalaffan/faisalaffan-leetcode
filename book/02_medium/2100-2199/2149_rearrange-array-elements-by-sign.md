# 2149 — Rearrange Array Elements By Sign

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func rearrangeArray(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(n)


## 💻 Solusi Go

```go
package main

// LeetCode #2149: Rearrange Array Elements by Sign
// https://leetcode.com/problems/rearrange-array-elements-by-sign/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import "fmt"

func rearrangeArray(nums []int) []int {
	n := len(nums)
  // Alokasi slice
	result := make([]int, n)
	posIdx, negIdx := 0, 1

	for _, v := range nums {
		if v > 0 {
			result[posIdx] = v
			posIdx += 2
		} else {
			result[negIdx] = v
			negIdx += 2
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
	// Expected: [3,-2,1,-5,2,-4]

	// Test case 2
	fmt.Println("Test 2:", rearrangeArray([]int{-1, 1}))
	// Expected: [1,-1]

	// Test case 3
	fmt.Println("Test 3:", rearrangeArray([]int{28, -41, 22, -8, -37, 46, 35, -9}))
	// Expected: [28,-41,22,-8,46,-37,35,-9]
}
```
