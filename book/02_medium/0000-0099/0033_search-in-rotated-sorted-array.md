# 0033 — Search In Rotated Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func search(nums []int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(log n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #33: Search in Rotated Sorted Array
// https://leetcode.com/problems/search-in-rotated-sorted-array/
// Difficulty: Medium

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

  // Binary search loop
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}

		if nums[left] <= nums[mid] {
			// Left half is sorted
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Right half is sorted
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}

func main() {
	// Test case 1
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0)) // 4

	// Test case 2
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3)) // -1

	// Test case 3
	fmt.Println(search([]int{1}, 0)) // -1
}

// Time: O(log n) | Space: O(1)
```
