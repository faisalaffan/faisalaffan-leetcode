# 1664 — Ways To Make A Fair Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func WaysToMakeFair(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(N), Space: O(1)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1664: Ways to Make a Fair Array
// https://leetcode.com/problems/ways-to-make-a-fair-array/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(WaysToMakeFair([]int{2, 1, 6, 4}))
	fmt.Println(WaysToMakeFair([]int{1, 1, 1}))
	fmt.Println(WaysToMakeFair([]int{1, 2, 3, 4, 5}))
}

func WaysToMakeFair(nums []int) int {
	// Time: O(N), Space: O(1)

	// Calculate total sum at even and odd indices
	totalEven := 0
	totalOdd := 0
	for i, num := range nums {
		if i%2 == 0 {
			totalEven += num
		} else {
			totalOdd += num
		}
	}

	result := 0
	prefixEven := 0
	prefixOdd := 0

	for i, num := range nums {
		if i%2 == 0 {
			totalEven -= num
		} else {
			totalOdd -= num
		}

		// After removing nums[i], all indices shift:
		// Elements to the right of i swap parity
		// Even sum = prefixEven + totalOdd
		// Odd sum = prefixOdd + totalEven
		if prefixEven+totalOdd == prefixOdd+totalEven {
			result++
		}

		if i%2 == 0 {
			prefixEven += num
		} else {
			prefixOdd += num
		}
	}

	return result
}
```
