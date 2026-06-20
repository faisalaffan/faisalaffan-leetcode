# 0325 — Maximum Size Subarray Sum Equals K

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func maxSubArrayLen(nums []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(n)

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #325: Maximum Size Subarray Sum Equals k
// https://leetcode.com/problems/maximum-size-subarray-sum-equals-k/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(n)

import "fmt"

func maxSubArrayLen(nums []int, k int) int {
	prefixSum := 0
	maxLen := 0
	// Map prefix sum -> earliest index
	sumMap := map[int]int{0: -1}

	for i, num := range nums {
		prefixSum += num

		if idx, ok := sumMap[prefixSum-k]; ok {
			if i-idx > maxLen {
				maxLen = i - idx
			}
		}

		// Only store first occurrence (earliest index)
		if _, ok := sumMap[prefixSum]; !ok {
			sumMap[prefixSum] = i
		}
	}
	return maxLen
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxSubArrayLen([]int{1, -1, 5, -2, 3}, 3))
	// Expected: 4 ([1,-1,5,-2])

	// Test case 2
	fmt.Println("Test 2:", maxSubArrayLen([]int{-2, -1, 2, 1}, 1))
	// Expected: 2 ([-1,2])

	// Test case 3
	fmt.Println("Test 3:", maxSubArrayLen([]int{1, 2, 3}, 6))
	// Expected: 3 ([1,2,3])
}
```
