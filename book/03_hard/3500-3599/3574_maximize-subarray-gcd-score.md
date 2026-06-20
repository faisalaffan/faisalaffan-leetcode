# 3574 — Maximize Subarray Gcd Score

## Deskripsi

**Soal:** [3574. Maximize Subarray Gcd Score](https://leetcode.com/problems/maximize-subarray-gcd-score/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** For each subarray, compute GCD and count of elements with minimum

## Solusi Go

```go
package main

// LeetCode #3574: Maximize Subarray GCD Score
// https://leetcode.com/problems/maximize-subarray-gcd-score/
// Difficulty: Hard
//
// Given array nums and integer k, maximize length * GCD of a contiguous subarray
// after optionally doubling up to k elements (each at most once).
//
// Approach: For each subarray, compute GCD and count of elements with minimum
// trailing-zero count of factor 2. If that count <= k, GCD can be doubled.
// O(n^2) with GCD caching.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxGCDScore([]int{2, 4, 8, 16}, 2))
	// Example 2
	fmt.Println(maxGCDScore([]int{1, 2, 3, 4}, 1))
	// Edge: single element
	fmt.Println(maxGCDScore([]int{10}, 0))
	// Edge: all ones with k=0
	fmt.Println(maxGCDScore([]int{1, 1, 1}, 0))
}

func maxGCDScore(nums []int, k int) int64 {
	n := len(nums)
	var result int64

	// Precompute v2 count (trailing zeros / power of 2 factor)
  // Membuat slice untuk menyimpan hasil
	v2 := make([]int, n)
	for i, v := range nums {
		v2[i] = trailingZeros(v)
	}

	for i := 0; i < n; i++ {
		g := nums[i]
		minV2 := v2[i]
		cntMinV2 := 1
		length := 1

		// Update result for subarray starting at i
		score := int64(length) * int64(g)
		if cntMinV2 <= k {
			score = int64(length) * int64(g*2)
		}
		if score > result {
			result = score
		}

		for j := i + 1; j < n; j++ {
			g = gcd(g, nums[j])
			if v2[j] < minV2 {
				minV2 = v2[j]
				cntMinV2 = 1
			} else if v2[j] == minV2 {
				cntMinV2++
			}
			length++

			score := int64(length) * int64(g)
			if cntMinV2 <= k {
				score = int64(length) * int64(g*2)
			}
			if score > result {
				result = score
			}
		}
	}

	return result
}

func trailingZeros(x int) int {
	if x == 0 {
		return 0
	}
	c := 0
	for x%2 == 0 {
		x /= 2
		c++
	}
	return c
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
```
