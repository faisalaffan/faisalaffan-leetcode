# 0075 — Sort Colors

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func sortColors(nums []int) `

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #75: Sort Colors
// https://leetcode.com/problems/sort-colors/
// Difficulty: Medium

import "fmt"

func sortColors(nums []int) {
	left, right := 0, len(nums)-1
	curr := 0

	for curr <= right {
		if nums[curr] == 0 {
			nums[left], nums[curr] = nums[curr], nums[left]
			left++
			curr++
		} else if nums[curr] == 2 {
			nums[right], nums[curr] = nums[curr], nums[right]
			right--
		} else {
			curr++
		}
	}
}

func main() {
	// Test case 1
	nums1 := []int{2, 0, 2, 1, 1, 0}
	sortColors(nums1)
	fmt.Println(nums1) // [0 0 1 1 2 2]

	// Test case 2
	nums2 := []int{2, 0, 1}
	sortColors(nums2)
	fmt.Println(nums2) // [0 1 2]

	// Test case 3
	nums3 := []int{0}
	sortColors(nums3)
	fmt.Println(nums3) // [0]
}

// Time: O(n) | Space: O(1)
```
