# 3940 — Limit Occurrences In Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Mudah

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func LimitOccurrencesInSortedArray(nums []int, k int) []int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #3940: Limit Occurrences in Sorted Array
// https://leetcode.com/problems/limit-occurrences-in-sorted-array/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(LimitOccurrencesInSortedArray([]int{1, 1, 1, 2, 2, 3}, 2))
	fmt.Println(LimitOccurrencesInSortedArray([]int{1, 2, 3}, 1))
}

// Time: O(n)
// Space: O(1)
func LimitOccurrencesInSortedArray(nums []int, k int) []int {
  // Edge case: input kosong
	if len(nums) == 0 {
		return nums
	}
	write := 0
	for _, v := range nums {
		if write < k || v != nums[write-k] {
			nums[write] = v
			write++
		}
	}
	return nums[:write]
}
```
