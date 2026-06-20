# 2091 — Removing Minimum And Maximum From Array

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minimumDeletions(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2091: Removing Minimum and Maximum From Array
// https://leetcode.com/problems/removing-minimum-and-maximum-from-array/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func minimumDeletions(nums []int) int {
	n := len(nums)
	minIdx, maxIdx := 0, 0
	minVal, maxVal := nums[0], nums[0]

	for i, v := range nums {
		if v < minVal {
			minVal = v
			minIdx = i
		}
		if v > maxVal {
			maxVal = v
			maxIdx = i
		}
	}

	// Possible strategies:
	// 1. Delete both from front: max(minIdx, maxIdx) + 1
	// 2. Delete both from back: n - min(minIdx, maxIdx)
	// 3. Delete one from front, one from back: min(minIdx, maxIdx) + 1 + n - max(minIdx, maxIdx)
	front := max(minIdx, maxIdx) + 1
	back := n - min(minIdx, maxIdx)
	mixed := min(minIdx, maxIdx) + 1 + n - max(minIdx, maxIdx)

	return min(front, min(back, mixed))
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minimumDeletions([]int{2, 10, 7, 5, 4, 1, 8, 6}))
	// Expected: 5

	// Test case 2
	fmt.Println("Test 2:", minimumDeletions([]int{0, -4, 19, 1, 8, -2, -3, 5}))
	// Expected: 3

	// Test case 3
	fmt.Println("Test 3:", minimumDeletions([]int{101}))
	// Expected: 1
}
```
