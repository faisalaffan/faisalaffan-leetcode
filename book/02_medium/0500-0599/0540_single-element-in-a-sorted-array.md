# 0540 — Single Element In A Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func SingleNonDuplicate(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #540: Single Element in a Sorted Array
// https://leetcode.com/problems/single-element-in-a-sorted-array/
// Difficulty: Medium
// Time: O(log n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(SingleNonDuplicate([]int{1, 1, 2, 3, 3, 4, 4, 8, 8}))
	fmt.Println(SingleNonDuplicate([]int{3, 3, 7, 7, 10, 11, 11}))
}

func SingleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		mid := lo + (hi-lo)/2
		// Ensure mid is even to compare with mid+1
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			lo = mid + 2
		} else {
			hi = mid
		}
	}
	return nums[lo]
}
```
