# 1060 — Missing Element In Sorted Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func missingElement(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Binary Search

**Waktu:** O(log n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Binary Search** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #1060: Missing Element in Sorted Array
// https://leetcode.com/problems/missing-element-in-sorted-array/
// Difficulty: Medium
//
// Approach: Binary search. Number of missing elements up to index i = nums[i] - nums[0] - i.
// Time: O(log n)
// Space: O(1)

import "fmt"

func main() {
	fmt.Println(missingElement([]int{4, 7, 9, 10}, 1))  // 5
	fmt.Println(missingElement([]int{4, 7, 9, 10}, 3))  // 8
	fmt.Println(missingElement([]int{1, 2, 4}, 3))      // 6
}

func missingElement(nums []int, k int) int {
	n := len(nums)

	missingCount := func(idx int) int {
		return nums[idx] - nums[0] - idx
	}

	if missingCount(n-1) < k {
		return nums[n-1] + (k - missingCount(n-1))
	}

	left, right := 0, n-1
  // Two-pointer loop
	for left < right {
		mid := left + (right-left)/2
		if missingCount(mid) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left-1] + (k - missingCount(left-1))
}
```
