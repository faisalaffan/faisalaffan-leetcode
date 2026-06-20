# 3480 — Maximize Subarrays After Removing One Conflicting Pair

## Deskripsi

**Soal:** [3480. Maximize Subarrays After Removing One Conflicting Pair](https://leetcode.com/problems/maximize-subarrays-after-removing-one-conflicting-pair/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), LIS (Longest Increasing Subsequence)

> **Ide Kunci:** For each conflicting pair, compute how many subarrays are blocked

## Solusi Go

```go
package main

// LeetCode #3480: Maximize Subarrays After Removing One Conflicting Pair
// https://leetcode.com/problems/maximize-subarrays-after-removing-one-conflicting-pair/
// Difficulty: Hard
//
// Given n and a list of conflicting pairs, remove exactly one conflicting pair
// to maximize the number of subarrays that can be formed.
//
// Approach: For each conflicting pair, compute how many subarrays are blocked
// by it, pick the pair whose removal unblocks the most subarrays.

import "fmt"

func main() {
	// Example 1
	fmt.Println(maxSubarrays(4, [][]int{{2, 3}, {0, 1}}))
	// Example 2
	fmt.Println(maxSubarrays(5, [][]int{{0, 2}, {1, 3}, {2, 4}}))
	// Edge: single pair
	fmt.Println(maxSubarrays(3, [][]int{{0, 1}}))
	// Edge: no pairs
	fmt.Println(maxSubarrays(3, [][]int{}))
}

func maxSubarrays(n int, conflictingPairs [][]int) int64 {
	if len(conflictingPairs) == 0 {
		// Total subarrays = n*(n+1)/2
		return int64(n) * int64(n+1) / 2
	}

	// Each conflicting pair blocks subarrays that include both endpoints
	// A pair (i, j) blocks subarrays where left <= i and right >= j
	// For each pair, count blocked subarrays = (i+1) * (n-j)

	// Track per-pair blocked count and find the pair whose removal
	// maximizes the total unblocked

	totalSubarrays := int64(n) * int64(n+1) / 2

	// Compute blocked subarrays for each pair
	maxBlocked := int64(0)
	totalBlocked := int64(0)
	for _, pair := range conflictingPairs {
		i, j := pair[0], pair[1]
		if i > j {
			i, j = j, i
		}
		blocked := int64(i+1) * int64(n-j)
		if blocked > maxBlocked {
			maxBlocked = blocked
		}
		totalBlocked += blocked
	}

	// Remove the pair that blocks the most subarrays
	ans := totalSubarrays - totalBlocked + maxBlocked
	return ans
}
```
