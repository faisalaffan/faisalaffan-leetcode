# 1191 — K Concatenation Maximum Sum

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sedang

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func kConcatenationMaxSum(arr []int, k int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** Prefix Sum

**Waktu:** O(n)  |  **Ruang:** O(1)

> 🎓 **Fresh Grad Tips:** Kuasai **Prefix Sum** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

import (
	"fmt"
)

// LeetCode #1191: K-Concatenation Maximum Sum
// https://leetcode.com/problems/k-concatenation-maximum-sum/
// Difficulty: Medium

// If k == 1, just Kadane. If k >= 2, compute max suffix + max prefix +
// (k-2)*total if total > 0.

// Time: O(n)
// Space: O(1)

const mod = 1_000_000_007

func kConcatenationMaxSum(arr []int, k int) int {
	n := len(arr)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// Kadane on single array
	kadane := func(nums []int) int64 {
		var maxEnd, maxSoFar int64
		for _, v := range nums {
			maxEnd = max64(maxEnd+int64(v), int64(v))
			if maxEnd > maxSoFar {
				maxSoFar = maxEnd
			}
		}
		return maxSoFar
	}

	singleMax := kadane(arr)

	if k == 1 {
		return int(singleMax % mod)
	}

	// Total sum
	var total int64
	for _, v := range arr {
		total += int64(v)
	}

	// Max prefix sum
	var prefixSum, maxPrefix int64
	for _, v := range arr {
		prefixSum += int64(v)
		if prefixSum > maxPrefix {
			maxPrefix = prefixSum
		}
	}

	// Max suffix sum
	var suffixSum, maxSuffix int64
	for i := n - 1; i >= 0; i-- {
		suffixSum += int64(arr[i])
		if suffixSum > maxSuffix {
			maxSuffix = suffixSum
		}
	}

	doubleMax := kadane(append(arr, arr...))
	if k == 2 {
		if doubleMax > singleMax {
			return int(doubleMax % mod)
		}
		return int(singleMax % mod)
	}

	// k >= 3
	var result int64
	result = max64(singleMax, max64(doubleMax, maxPrefix+maxSuffix+total*int64(k-2)))
	// Also consider just maxPrefix + maxSuffix + total * (k-2), but may overflow
	result = max64(result, maxPrefix+maxSuffix)

	if result < 0 {
		return 0
	}
	return int(result % mod)
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Printf("%d (expected: 9)\n", kConcatenationMaxSum([]int{1, -2, 1}, 5))
	fmt.Printf("%d (expected: 2)\n", kConcatenationMaxSum([]int{-1, -2}, 7))
	fmt.Printf("%d (expected: 4)\n", kConcatenationMaxSum([]int{1, 2}, 1))
}
```
