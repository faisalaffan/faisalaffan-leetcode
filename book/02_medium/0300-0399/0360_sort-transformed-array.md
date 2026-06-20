# 0360 — Sort Transformed Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func sortTransformedArray(nums []int, a int, b int, c int) []int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Two Pointer

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Two Pointer** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #360: Sort Transformed Array
// https://leetcode.com/problems/sort-transformed-array/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func sortTransformedArray(nums []int, a int, b int, c int) []int {
	n := len(nums)
  // Alokasi slice
	result := make([]int, n)
	left, right := 0, n-1

	// Parabola opens upward → fill from right; downward → fill from left
	idx := n - 1
	if a < 0 {
		idx = 0
	}

	f := func(x int) int {
		return a*x*x + b*x + c
	}

  // Binary search loop
	for left <= right {
		lv, rv := f(nums[left]), f(nums[right])
		if a >= 0 {
			if lv > rv {
				result[idx] = lv
				left++
			} else {
				result[idx] = rv
				right--
			}
			idx--
		} else {
			if lv < rv {
				result[idx] = lv
				left++
			} else {
				result[idx] = rv
				right--
			}
			idx++
		}
	}
	return result
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", sortTransformedArray([]int{-4, -2, 2, 4}, 1, 3, 5))
	// Expected: [3, 9, 15, 33]

	// Test case 2
	fmt.Println("Test 2:", sortTransformedArray([]int{-4, -2, 2, 4}, -1, 3, 5))
	// Expected: [-23, -5, 1, 7]

	// Test case 3: Single element
	fmt.Println("Test 3:", sortTransformedArray([]int{0}, 1, 0, 0))
	// Expected: [0]
}
```
