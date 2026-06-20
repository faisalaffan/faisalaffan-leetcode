# 0080 — Remove Duplicates From Sorted Array Ii

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func removeDuplicates(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #80: Remove Duplicates from Sorted Array II
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/
// Difficulty: Medium

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}

	write := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[write-2] {
			nums[write] = nums[i]
			write++
		}
	}

	return write
}

func main() {
	// Test case 1
	nums1 := []int{1, 1, 1, 2, 2, 3}
	k1 := removeDuplicates(nums1)
	fmt.Println(nums1[:k1]) // [1 1 2 2 3]

	// Test case 2
	nums2 := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	k2 := removeDuplicates(nums2)
	fmt.Println(nums2[:k2]) // [0 0 1 1 2 3 3]
}

// Time: O(n) | Space: O(1)
```
