# 3830 — Longest Alternating Subarray After Removing At Most One Element

## Deskripsi

**Soal:** [3830. Longest Alternating Subarray After Removing At Most One Element](https://leetcode.com/problems/longest-alternating-subarray-after-removing-at-most-one-element/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Precompute longest alternating ending and starting at

## Solusi Go

```go
package main

// LeetCode #3830: Longest Alternating Subarray After Removing
// At Most One Element
// https://leetcode.com/problems/longest-alternating-subarray-after-removing-at-most-one-element/
// Difficulty: Hard
//
// Find the longest subarray where adjacent elements alternate
// (nums[i] != nums[i-1]) after removing at most one element.
//
// Approach: Precompute longest alternating ending and starting at
// each position. For each possible removal, combine left+right
// alternating segments.

import "fmt"

func main() {
	// Example 1
	fmt.Println(longestAlternating([]int{1, 2, 3, 4}))
	// Example 2
	fmt.Println(longestAlternating([]int{1, 2, 1, 2, 3}))
	// Edge: all same
	fmt.Println(longestAlternating([]int{3, 3, 3}))
	// Edge: single element
	fmt.Println(longestAlternating([]int{5}))
}

func longestAlternating(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	// pref[i] = longest alternating subarray ending at i
  // Membuat slice untuk menyimpan hasil
	pref := make([]int, n)
	pref[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] != nums[i-1] {
			pref[i] = pref[i-1] + 1
		} else {
			pref[i] = 1
		}
	}

	// suff[i] = longest alternating subarray starting at i
  // Membuat slice untuk menyimpan hasil
	suff := make([]int, n)
	suff[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		if nums[i] != nums[i+1] {
			suff[i] = suff[i+1] + 1
		} else {
			suff[i] = 1
		}
	}

	// Best without any removal
	best := 0
	for _, v := range pref {
		if v > best {
			best = v
		}
	}

	// Try removing each element i
	for i := 0; i < n; i++ {
		// Remove nums[i], combine pref[i-1] and suff[i+1]
		// They can be combined only if nums[i-1] != nums[i+1]
		left := 0
		if i > 0 {
			left = pref[i-1]
		}
		right := 0
		if i < n-1 {
			right = suff[i+1]
		}

		combined := 0
		if left > 0 && right > 0 {
			if nums[i-1] != nums[i+1] {
				combined = left + right
			} else {
				combined = max(left, right)
			}
		} else {
			combined = left + right
		}
		if combined > best {
			best = combined
		}
	}

	return best
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```
