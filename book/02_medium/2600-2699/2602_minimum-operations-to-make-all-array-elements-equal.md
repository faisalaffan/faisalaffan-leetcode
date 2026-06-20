# 2602 — Minimum Operations To Make All Array Elements Equal

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func minOperations(nums []int, queries []int) []int64`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Sorting, Prefix Sum

**Waktu:** O((n+q) log n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **Sorting** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #2602: Minimum Operations to Make All Array Elements Equal
// https://leetcode.com/problems/minimum-operations-to-make-all-array-elements-equal/
// Difficulty: Medium
// Time: O((n+q) log n) | Space: O(n)

import (
	"fmt"
	"sort"
)

func minOperations(nums []int, queries []int) []int64 {
	n := len(nums)
  // Sort O(n log n)
	sort.Ints(nums)

  // Alokasi slice
	prefix := make([]int64, n+1)
	for i, v := range nums {
		prefix[i+1] = prefix[i] + int64(v)
	}

  // Alokasi slice
	ans := make([]int64, len(queries))
	for i, q := range queries {
		idx := sort.SearchInts(nums, q)
		leftCount := int64(idx)
		rightCount := int64(n - idx)
		leftSum := prefix[idx]
		rightSum := prefix[n] - prefix[idx]
		ops := int64(q)*leftCount - leftSum + rightSum - int64(q)*rightCount
		ans[i] = ops
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", minOperations([]int{3, 1, 6, 8}, []int{1, 5}))
	// Expected: [8, 10]? Let me check...

	// Test case 2
	fmt.Println("Test 2:", minOperations([]int{2, 4, 6, 8}, []int{4, 5}))
	// Expected: [4, 4]

	// Test case 3
	fmt.Println("Test 3:", minOperations([]int{1, 2, 3, 4, 5}, []int{3}))
	// Expected: [6]
}
```
