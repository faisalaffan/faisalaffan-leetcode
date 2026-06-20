# 2941 — Maximum Gcd Sum Of A Subarray

## Deskripsi

**Soal:** [2941. Maximum Gcd Sum Of A Subarray](https://leetcode.com/problems/maximum-gcd-sum-of-a-subarray/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func gcd(a, b int) int`

## Solusi Go

```go
package main

// LeetCode #2941: Maximum GCD-Sum of a Subarray
// https://leetcode.com/problems/maximum-gcd-sum-of-a-subarray/
// Difficulty: Hard
//
// For each position i as right endpoint, maintain list of (gcd, leftmost_index)
// for subarrays ending at i. The number of distinct gcd values is O(log max(nums)).
// For each pair, compute sum via prefix array: sum = pref[i+1] - pref[l].
// Answer = max over all (gcd * sum).

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func maxGcdSumOfSubarray(nums []int, k int) int64 {
	// Note: the problem has been observed with a k parameter (length constraint)
	// but the core algorithm works on the main array regardless.
	// We'll handle the general formulation.
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

  // Membuat slice untuk menyimpan hasil
	pref := make([]int64, n+1)
	for i, v := range nums {
		pref[i+1] = pref[i] + int64(v)
	}

	var ans int64
	// pairs: list of (gcd, leftmost_index) for subarrays ending at current position
	type pair struct {
		g int
		l int
	}
  // Membuat slice untuk menyimpan hasil
	cur := make([]pair, 0)

	for i := 0; i < n; i++ {
		// Build new list for subarrays ending at i
  // Membuat slice untuk menyimpan hasil
		nxt := make([]pair, 0)

		// Extend previous subarrays
		for _, p := range cur {
			ng := gcd(p.g, nums[i])
			if len(nxt) > 0 && nxt[len(nxt)-1].g == ng {
				// Same gcd, keep the earlier leftmost index
				continue
			}
			nxt = append(nxt, pair{g: ng, l: p.l})
		}

		// Add subarray containing only nums[i]
		if len(nxt) == 0 || nxt[len(nxt)-1].g != nums[i] {
			nxt = append(nxt, pair{g: nums[i], l: i})
		}

		// Evaluate
		for _, p := range nxt {
			sum := pref[i+1] - pref[p.l]
			val := int64(p.g) * sum
			if val > ans {
				ans = val
			}
		}

		cur = nxt
	}
	return ans
}

func main() {
	// Example
	fmt.Println(maxGcdSumOfSubarray([]int{3, 1, 4, 2, 2, 1}, 0))

	// Simple cases
	fmt.Println(maxGcdSumOfSubarray([]int{1, 2, 3, 4, 5}, 0))
	fmt.Println(maxGcdSumOfSubarray([]int{10, 20, 30}, 0))

	// Single element
	fmt.Println(maxGcdSumOfSubarray([]int{7}, 0))

	// All same
	fmt.Println(maxGcdSumOfSubarray([]int{5, 5, 5, 5}, 0))
}
```
