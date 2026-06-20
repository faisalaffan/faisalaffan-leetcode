# 3312 — Sorted Gcd Pair Queries

## Deskripsi

**Soal:** [3312. Sorted Gcd Pair Queries](https://leetcode.com/problems/sorted-gcd-pair-queries/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner), Prefix Sum (jumlah kumulatif)

> **Ide Kunci:** Use divisor enumeration + inclusion-exclusion to count pairs by

## Solusi Go

```go
package main

// LeetCode #3312: Sorted GCD Pair Queries
// https://leetcode.com/problems/sorted-gcd-pair-queries/
// Difficulty: Hard
//
// Given an array nums and queries, for each query index i, return the i-th
// smallest value among all gcd(nums[a], nums[b]) for a < b.
//
// Approach: Use divisor enumeration + inclusion-exclusion to count pairs by
// GCD value. Build prefix sum, then binary search for each query.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(gcdValues([]int{2, 3, 4}, []int64{0, 2, 3}))
	// Example 2
	fmt.Println(gcdValues([]int{4, 4, 2, 1}, []int64{5, 3, 1, 0}))
	// Example 3
	fmt.Println(gcdValues([]int{2, 2}, []int64{0, 0}))
	// Edge case
	fmt.Println(gcdValues([]int{6, 10, 15}, []int64{0, 1, 2}))
}

func gcdValues(nums []int, queries []int64) []int {
	mx := 0
	for _, v := range nums {
		if v > mx {
			mx = v
		}
	}

	// Count frequency of each value
  // Membuat slice untuk menyimpan hasil
	cnt := make([]int, mx+1)
	for _, v := range nums {
		cnt[v]++
	}

	// Count divisor frequencies
  // Membuat slice untuk menyimpan hasil
	divCnt := make([]int, mx+1)
	for _, v := range nums {
		for d := 1; d*d <= v; d++ {
			if v%d == 0 {
				divCnt[d]++
				if d != v/d {
					divCnt[v/d]++
				}
			}
		}
	}

	// Inclusion-exclusion: count pairs with exact GCD = g
  // Membuat slice untuk menyimpan hasil
	pairCnt := make([]int, mx+1)
	for g := mx; g >= 1; g-- {
		c := divCnt[g]
		pairCnt[g] = c * (c - 1) / 2
		for multiple := 2 * g; multiple <= mx; multiple += g {
			pairCnt[g] -= pairCnt[multiple]
		}
	}

	// Prefix sum
  // Membuat slice untuk menyimpan hasil
	prefix := make([]int, mx+1)
	for i := 1; i <= mx; i++ {
		prefix[i] = prefix[i-1] + pairCnt[i]
	}

  // Membuat slice untuk menyimpan hasil
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = sort.Search(mx+1, func(j int) bool {
			return prefix[j] > int(q)
		})
	}
	return ans
}
```
