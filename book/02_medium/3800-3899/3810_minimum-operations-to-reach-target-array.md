# 3810 — Minimum Operations To Reach Target Array

## Deskripsi

**Soal:** [3810. Minimum Operations To Reach Target Array](https://leetcode.com/problems/minimum-operations-to-reach-target-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(max(nums[i]))

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func MinimumOperationsToReachTargetArray(nums []int, target []int) int`

> **Ide Kunci:** Track which values need updating. Each operation picks a value x,

## Solusi Go

```go
package main

// LeetCode #3810: Minimum Operations to Reach Target Array
// https://leetcode.com/problems/minimum-operations-to-reach-target-array/
// Difficulty: Medium
// Time: O(N) | Space: O(max(nums[i]))
// Approach: Track which values need updating. Each operation picks a value x,
// finds all maximal contiguous segments where nums[i]==x, and simultaneously
// updates them to target values.

import "fmt"

func MinimumOperationsToReachTargetArray(nums []int, target []int) int {
	n := len(nums)
	// Map from value to list of positions
  // Membuat map untuk pencarian O(1): key → value
	valToPos := make(map[int][]int)
	for i, v := range nums {
		valToPos[v] = append(valToPos[v], i)
	}

	// Track positions that still need changes
  // Membuat map untuk pencarian O(1): key → value
	needChange := make(map[int]bool)
	for i := 0; i < n; i++ {
		if nums[i] != target[i] {
			needChange[i] = true
		}
	}

	ops := 0
	for len(needChange) > 0 {
		// Find a value that appears at positions needing change
		var chosenVal int
		found := false
		for v, pos := range valToPos {
			for _, p := range pos {
				if needChange[p] {
					chosenVal = v
					found = true
					break
				}
			}
			if found {
				break
			}
		}
		if !found {
			break
		}

		// Update all maximal contiguous segments of chosenVal
  // Membuat map untuk pencarian O(1): key → value
		changed := make(map[int]bool)
		for _, p := range valToPos[chosenVal] {
			if needChange[p] {
				changed[p] = true
			}
		}

		// Process each changed position
		for p := range changed {
			if !needChange[p] {
				continue
			}
			// Expand to full contiguous segment
			l, r := p, p
			for l-1 >= 0 && nums[l-1] == chosenVal {
				l--
			}
			for r+1 < n && nums[r+1] == chosenVal {
				r++
			}
			// Update the segment
			for i := l; i <= r; i++ {
				if needChange[i] {
					nums[i] = target[i]
					delete(needChange, i)
				}
			}
		}
		ops++
	}

	return ops
}

func main() {
	// Example 1
	fmt.Println(MinimumOperationsToReachTargetArray([]int{1, 2, 3}, []int{2, 1, 3})) // Expected: 2

	// Example 2
	fmt.Println(MinimumOperationsToReachTargetArray([]int{4, 1, 4}, []int{5, 1, 4})) // Expected: 1

	// Example 3
	fmt.Println(MinimumOperationsToReachTargetArray([]int{7, 3, 7}, []int{5, 5, 9})) // Expected: 2
}
```
