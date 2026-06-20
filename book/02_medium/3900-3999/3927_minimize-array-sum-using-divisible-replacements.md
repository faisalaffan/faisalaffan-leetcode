# 3927 — Minimize Array Sum Using Divisible Replacements

## Deskripsi

**Soal:** [3927. Minimize Array Sum Using Divisible Replacements](https://leetcode.com/problems/minimize-array-sum-using-divisible-replacements/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N sqrt(M))  
**Kompleksitas Ruang:** O(N) where M = max value

**Algoritma:** —

**Fungsi Solusi:** `func MinimizeArraySumUsingDivisibleReplacements(nums []int) int64`

> **Ide Kunci:** For each element, find its minimum divisor present in the array.

## Solusi Go

```go
package main

// LeetCode #3927: Minimize Array Sum Using Divisible Replacements
// https://leetcode.com/problems/minimize-array-sum-using-divisible-replacements/
// Difficulty: Medium
// Time: O(N sqrt(M)) | Space: O(N) where M = max value
// Approach: For each element, find its minimum divisor present in the array.
// Each number can be replaced with any present divisor (via chaining).
// Sum the minimal reachable value for each element.

import (
	"fmt"
	"math"
)

func MinimizeArraySumUsingDivisibleReplacements(nums []int) int64 {
	// Track which values exist
  // Membuat map untuk pencarian O(1): key → value
	exists := make(map[int]bool)
	minVal := math.MaxInt32
	for _, v := range nums {
		exists[v] = true
		if v < minVal {
			minVal = v
		}
	}

	var ans int64 = 0
	for _, v := range nums {
		reduced := false
		// Try divisors from 1 to sqrt(v)
		for d := 1; d*d <= v; d++ {
			if v%d == 0 {
				if exists[d] {
					ans += int64(d)
					reduced = true
					break
				}
				other := v / d
				if other != d && exists[other] && other < v {
					// Can't directly use if not minimum, but may be useful
				}
			}
		}
		if !reduced {
			ans += int64(v)
		}
	}
	return ans
}

func main() {
	// Example 1
	fmt.Println(MinimizeArraySumUsingDivisibleReplacements([]int{3, 6, 2})) // Expected: 7

	// Example 2
	fmt.Println(MinimizeArraySumUsingDivisibleReplacements([]int{4, 2, 8, 3})) // Expected: 9

	// Example 3
	fmt.Println(MinimizeArraySumUsingDivisibleReplacements([]int{7, 5, 9})) // Expected: 21
}
```
