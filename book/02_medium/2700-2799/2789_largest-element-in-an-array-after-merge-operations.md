# 2789 — Largest Element In An Array After Merge Operations

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func LargestElementInAnArrayAfterMergeOperations(nums []int) int64`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

// LeetCode #2789: Largest Element in an Array after Merge Operations
// https://leetcode.com/problems/largest-element-in-an-array-after-merge-operations/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func LargestElementInAnArrayAfterMergeOperations(nums []int) int64 {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	result := int64(nums[n-1])
	for i := n - 2; i >= 0; i-- {
		if int64(nums[i]) <= result {
			result += int64(nums[i])
		} else {
			result = int64(nums[i])
		}
	}

	return result
}

func main() {
	fmt.Println(LargestElementInAnArrayAfterMergeOperations([]int{2, 3, 7, 9, 3}))
	fmt.Println(LargestElementInAnArrayAfterMergeOperations([]int{5, 3, 3}))
}
```
