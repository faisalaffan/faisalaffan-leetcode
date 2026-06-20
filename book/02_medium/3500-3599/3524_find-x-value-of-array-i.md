# 3524 — Find X Value Of Array I

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func FindXValueOfArrayI(nums []int, target int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer, Binary Search, Sorting, Prefix Sum

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3524: Find X Value of Array I
// https://leetcode.com/problems/find-x-value-of-array-i/
// Difficulty: Medium
// Complexity: O(n log n) time, O(1) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", FindXValueOfArrayI([]int{1, 2, 3, 4}, 4))
	// Test case 2
	fmt.Println("Test 2:", FindXValueOfArrayI([]int{1, 1, 2, 3}, 3))
	// Test case 3
	fmt.Println("Test 3:", FindXValueOfArrayI([]int{5, 1, 3}, 2))
}

func FindXValueOfArrayI(nums []int, target int) int {
  // Sort O(n log n)
	sort.Ints(nums)
	// Find x such that sum of (nums[i] > x ? x : nums[i]) equals target
	// Using prefix sums and binary search
  // Alokasi slice
	prefix := make([]int, len(nums)+1)
  // Linear scan O(n)
	for i := 0; i < len(nums); i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	// Binary search on x
	left, right := 0, nums[len(nums)-1]
  // Two-pointer loop
	for left < right {
		mid := left + (right-left)/2
		// Find first index > mid
		idx := sort.Search(len(nums), func(i int) bool { return nums[i] > mid })
		sum := prefix[idx] + mid*(len(nums)-idx)
		if sum >= target {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
```
