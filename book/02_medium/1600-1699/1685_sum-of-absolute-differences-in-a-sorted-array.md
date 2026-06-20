# 1685 — Sum Of Absolute Differences In A Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func getSumAbsoluteDifferences(nums []int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n), Space: O(1) (excluding output)  |  **Ruang:** O(1) (excluding output)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1685: Sum of Absolute Differences in a Sorted Array
// https://leetcode.com/problems/sum-of-absolute-differences-in-a-sorted-array/
// Difficulty: Medium
// Time: O(n), Space: O(1) (excluding output)

import "fmt"

func getSumAbsoluteDifferences(nums []int) []int {
	n := len(nums)
	total := 0
	for _, v := range nums {
		total += v
	}

  // Alokasi slice
	result := make([]int, n)
	prefix := 0
	for i, v := range nums {
		// Left side: v * i - prefix_sum_left
		// Right side: (total - prefix_sum_left - v) - v * (n-1-i)
		left := v*i - prefix
		right := (total - prefix - v) - v*(n-1-i)
		result[i] = left + right
		prefix += v
	}
	return result
}

func main() {
	fmt.Println(getSumAbsoluteDifferences([]int{2, 3, 5}))    // Expected: [4, 3, 5]
	fmt.Println(getSumAbsoluteDifferences([]int{1, 4, 6, 8, 10})) // Expected: [24, 15, 13, 15, 21]
	fmt.Println(getSumAbsoluteDifferences([]int{1, 2}))       // Expected: [1, 1]
}
```
