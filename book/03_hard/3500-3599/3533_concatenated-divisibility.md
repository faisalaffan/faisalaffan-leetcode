# 3533 — Concatenated Divisibility

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func concatenatedDivisibility(nums []int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** —  |  **Ruang:** —


## 💻 Solusi Go

```go
package main

// LeetCode #3533: Concatenated Divisibility
// https://leetcode.com/problems/concatenated-divisibility/
// Difficulty: Hard
//
// Given an array nums and integer k, find how many pairs (i,j) with i<j
// where concatenating nums[i] and nums[j] (as strings) is divisible by k.
//
// Approach: Process each number, track remainders of concatenated pairs.

import (
	"fmt"
	"math"
)

func main() {
	// Example 1
	fmt.Println(concatenatedDivisibility([]int{1, 2, 3, 4}, 5))
	// Example 2
	fmt.Println(concatenatedDivisibility([]int{10, 2, 3, 5}, 7))
	// Edge: single element
	fmt.Println(concatenatedDivisibility([]int{5}, 3))
	// Edge: all divisible
	fmt.Println(concatenatedDivisibility([]int{12, 34, 56}, 2))
}

func concatenatedDivisibility(nums []int, k int) []int {
	n := len(nums)
	var result []int

	for i := 0; i < n; i++ {
		count := 0
		for j := i + 1; j < n; j++ {
			// Concatenate nums[i] and nums[j]
			concat := concat(nums[i], nums[j])
			if concat%k == 0 {
				count++
			}
		}
		result = append(result, count)
	}

	return result
}

func concat(a, b int) int {
	if b == 0 {
		return a * 10
	}
	digits := int(math.Log10(float64(b))) + 1
	return a * int(math.Pow10(digits)) + b
}
```
