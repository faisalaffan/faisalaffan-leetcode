# 0154 — Find Minimum In Rotated Sorted Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func findMin(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #154: Find Minimum in Rotated Sorted Array II
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/
// Difficulty: Hard

import (
	"fmt"
)

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

  // Two-pointer loop
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// Minimum is in the right half
			left = mid + 1
		} else if nums[mid] < nums[right] {
			// Minimum is in the left half (including mid)
			right = mid
		} else {
			// nums[mid] == nums[right], cannot determine, shrink
			right--
		}
	}

	return nums[left]
}

func main() {
	nums := []int{2, 2, 2, 0, 1}
	result := findMin(nums)
	expected := 0

	fmt.Printf("findMin(%v) = %d\n", nums, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```
