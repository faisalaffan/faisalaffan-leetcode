# 3299 — Sum Of Consecutive Subsequences

## 📖 Deskripsi Soal

**Tingkat Kesulitan:** Sulit

Kamu diberikan array. Tugasmu mencari, menghitung, atau memanipulasi elemen.

**Cara berpikir:** Struktur data paling dasar. Akses O(1). Gunakan HashMap untuk lookup cepat, Two Pointer untuk pencarian pasangan.

**Fungsi Solusi:** `func sumOfConsecutiveSubsequences(nums []int) int`

## 🔍 Petunjuk Penyelesaian

**Teknik:** HashMap, DP

**Waktu:** —  |  **Ruang:** —

> 🎓 **Fresh Grad Tips:** Kuasai **HashMap** — sering muncul di interview!

## 💻 Solusi Go

```go
package main

// LeetCode #3299: Sum of Consecutive Subsequences
// https://leetcode.com/problems/sum-of-consecutive-subsequences/
// Difficulty: Hard [Paid]
//
// Given an array nums, consider all subsequences (not necessarily contiguous)
// that form a consecutive sequence of integers (e.g., 3,4,5 or 7,8,9).
// For each such subsequence, compute the sum of its elements. Return the
// total sum across all consecutive subsequences modulo 1e9+7.
//
// Approach:
//   For each element nums[i], maintain DP for subsequences ending at this
//   element that form a consecutive chain.
//   Let dp[v] = (count, sum) for subsequences ending with value v.
//   For each element x = nums[i]:
//     1. Starting a new subsequence: count=1, sum=x.
//     2. Extending from x-1: add dp[x-1].count subsequences,
//        each with additional sum contribution of x.
//     3. Extending from x (same value): add dp[x].count subsequences.
//   Accumulate all sums.

import (
	"fmt"
)

const mod = 1_000_000_007

func main() {
	// Example 1
	fmt.Println(sumOfConsecutiveSubsequences([]int{1, 2, 3}))
	// Example 2
	fmt.Println(sumOfConsecutiveSubsequences([]int{1, 1, 2, 3}))
	// Example 3
	fmt.Println(sumOfConsecutiveSubsequences([]int{5, 6, 7, 8}))
	// Example 4: single element
	fmt.Println(sumOfConsecutiveSubsequences([]int{10}))
	// Example 5: non-consecutive elements
	fmt.Println(sumOfConsecutiveSubsequences([]int{1, 3, 5, 7}))
	// Example 6: descending
	fmt.Println(sumOfConsecutiveSubsequences([]int{3, 2, 1}))
}

type pair struct {
	count int64
	sum   int64
}

func sumOfConsecutiveSubsequences(nums []int) int {
	// dp maps value -> (count, sum) for subsequences ending with that value.
  // HashMap: O(1) lookup
	dp := make(map[int]*pair)

	var total int64

	for _, x := range nums {
		// Count and sum for new subsequences ending at x.
		var cnt int64 = 1 // subsequence [x] alone
		var s int64 = int64(x)

		// Extend from x-1 (consecutive increasing chain).
		if p, ok := dp[x-1]; ok {
			cnt = (cnt + p.count) % mod
			// Each of p.count subsequences gets x added to its sum.
			s = (s + p.sum + p.count*int64(x)) % mod
		}

		// Extend from x (same value, allows repeated values in subsequence).
		if p, ok := dp[x]; ok {
			cnt = (cnt + p.count) % mod
			s = (s + p.sum + p.count*int64(x)) % mod
		}

		total = (total + s) % mod

		if dp[x] == nil {
			dp[x] = &pair{}
		}
		dp[x].count = (dp[x].count + cnt) % mod
		dp[x].sum = (dp[x].sum + s) % mod
	}

	return int(total)
}
```
