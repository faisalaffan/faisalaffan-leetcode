# 3427 — Sum Of Variable Length Subarrays

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SumOfVariableLengthSubarrays(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n^2). Space: O(1).  |  **Ruang:** O(1).


## 💻 Solusi Go

```go
package main

// LeetCode #3427: Sum of Variable Length Subarrays
// https://leetcode.com/problems/sum-of-variable-length-subarrays/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(SumOfVariableLengthSubarrays([]int{2, 3, 1}))
	fmt.Println(SumOfVariableLengthSubarrays([]int{3, 1, 1, 2}))
}

// SumOfVariableLengthSubarrays computes sum of subarrays where each subarray starts at i-(nums[i]%something) and ends at i.
// Time: O(n^2). Space: O(1).
func SumOfVariableLengthSubarrays(nums []int) int {
	n := len(nums)
	total := 0
	for i := 0; i < n; i++ {
		start := i - nums[i]
		if start < 0 {
			start = 0
		}
		for j := start; j <= i; j++ {
			total += nums[j]
		}
	}
	return total
}
```
