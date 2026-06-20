# 2913 — Subarrays Distinct Element Sum Of Squares I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SubarraysDistinctElementSumOfSquaresI(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap

**Waktu:** O(n^2)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2913: Subarrays Distinct Element Sum of Squares I
// https://leetcode.com/problems/subarrays-distinct-element-sum-of-squares-i/
// Difficulty: Easy

import "fmt"

func main() {
	// LeetCode name: sumCounts
	fmt.Println(SubarraysDistinctElementSumOfSquaresI([]int{1, 2, 1})) // 15
	fmt.Println(SubarraysDistinctElementSumOfSquaresI([]int{2, 2}))    // 3
}

// Time: O(n^2) | Space: O(n)
// LeetCode submission name: sumCounts
func SubarraysDistinctElementSumOfSquaresI(nums []int) int {
	n := len(nums)
	total := 0
	for i := 0; i < n; i++ {
  // HashMap: O(1) lookup
		seen := make(map[int]bool)
		distinct := 0
		for j := i; j < n; j++ {
			if !seen[nums[j]] {
				seen[nums[j]] = true
				distinct++
			}
			total += distinct * distinct
		}
	}
	return total
}
```
