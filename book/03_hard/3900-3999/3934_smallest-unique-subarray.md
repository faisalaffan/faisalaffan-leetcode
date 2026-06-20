# 3934 — Smallest Unique Subarray

## Deskripsi

**Soal:** [3934. Smallest Unique Subarray](https://leetcode.com/problems/smallest-unique-subarray/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

> **Ide Kunci:** Use suffix array or rolling hash to detect duplicates.

## Solusi Go

```go
package main

// LeetCode #3934: Smallest Unique Subarray
// https://leetcode.com/problems/smallest-unique-subarray/
// Difficulty: Hard
//
// Find the smallest subarray (by length, then lexicographically)
// that appears exactly once in the array (unique subarray).
//
// Approach: Use suffix array or rolling hash to detect duplicates.
// For each starting position i, find the longest prefix that
// appears elsewhere. The shortest unique subarray starting at i
// has length = longestCommonPrefix + 1. Track minimum.

import "fmt"

func main() {
	// Example 1
	fmt.Println(smallestUniqueSubarray([]int{1, 2, 1, 2}))
	// Example 2
	fmt.Println(smallestUniqueSubarray([]int{1, 2, 3, 4}))
	// Edge: all same
	fmt.Println(smallestUniqueSubarray([]int{1, 1, 1}))
}

func smallestUniqueSubarray(nums []int) int {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	bestLen := n + 1

	for start := 0; start < n; start++ {
		// Try growing subarray from start
		for end := start; end < n; end++ {
			length := end - start + 1
			if length >= bestLen {
				break
			}
			// Check if subarray nums[start:end+1] is unique
			unique := true
			for other := 0; other <= n-length; other++ {
				if other == start {
					continue
				}
				match := true
				for k := 0; k < length; k++ {
					if nums[other+k] != nums[start+k] {
						match = false
						break
					}
				}
				if match {
					unique = false
					break
				}
			}
			if unique {
				bestLen = length
				break
			}
		}
	}

	if bestLen > n {
		return 0
	}
	return bestLen
}
```
