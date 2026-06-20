# 1186 — Maximum Subarray Sum With One Deletion

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maximumSum(arr []int) int`

## 🔍 Petunjuk Penyelesaian

**Waktu:** O(n)  |  **Ruang:** O(1)


## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1186: Maximum Subarray Sum with One Deletion
// https://leetcode.com/problems/maximum-subarray-sum-with-one-deletion/
// Difficulty: Medium

// Kadane's algorithm variant with one deletion allowed.
// dp_no_del[i] = max subarray sum ending at i without deletion
// dp_del[i] = max subarray sum ending at i with one deletion

// Time: O(n)
// Space: O(1)

func maximumSum(arr []int) int {
	n := len(arr)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	noDel := arr[0]
	withDel := arr[0]
	result := arr[0]

	for i := 1; i < n; i++ {
		withDel = max(withDel+arr[i], noDel)
		noDel = max(noDel+arr[i], arr[i])
		result = max(result, max(noDel, withDel))
	}

	return result
}

func main() {
	fmt.Printf("%d (expected: 4)\n", maximumSum([]int{1, -2, 0, 3}))
	fmt.Printf("%d (expected: -1)\n", maximumSum([]int{-1, -1, -1, -1}))
	fmt.Printf("%d (expected: 7)\n", maximumSum([]int{1, -2, -2, 3, -1, 4}))
}
```
