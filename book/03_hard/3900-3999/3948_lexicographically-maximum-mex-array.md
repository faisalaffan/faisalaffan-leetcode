# 3948 — Lexicographically Maximum Mex Array

## Deskripsi

**Soal:** [3948. Lexicographically Maximum Mex Array](https://leetcode.com/problems/lexicographically-maximum-mex-array/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Track frequency of each value. For each i, find the

## Solusi Go

```go
package main

// LeetCode #3948: Lexicographically Maximum MEX Array
// https://leetcode.com/problems/lexicographically-maximum-mex-array/
// Difficulty: Hard
//
// Given array nums, construct a lexicographically maximum array
// result where result[i] is the MEX of a subsequence of nums
// ending at position i (or the MEX after certain operations).
//
// Approach: Track frequency of each value. For each i, find the
// MEX by checking the smallest non-negative integer not in the
// current multiset.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maximumMEXArray([]int{0, 1, 2, 3}))
	// Example 2
	fmt.Println(maximumMEXArray([]int{0, 0, 1, 1}))
	// Edge: empty
	fmt.Println(maximumMEXArray([]int{}))
}

func maximumMEXArray(nums []int) []int {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return []int{}
	}

  // Membuat map untuk pencarian O(1): key → value
	freq := make(map[int]int)
  // Membuat slice untuk menyimpan hasil
	result := make([]int, n)
	mex := 0

	for i, v := range nums {
		freq[v]++

		// Update mex: find smallest non-negative not in freq
		for freq[mex] > 0 {
			mex++
		}
		result[i] = mex
	}

	return result
}
```
