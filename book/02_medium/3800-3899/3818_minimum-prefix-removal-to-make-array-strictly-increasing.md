# 3818 — Minimum Prefix Removal To Make Array Strictly Increasing

## Deskripsi

**Soal:** [3818. Minimum Prefix Removal To Make Array Strictly Increasing](https://leetcode.com/problems/minimum-prefix-removal-to-make-array-strictly-increasing/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func MinimumPrefixRemovalToMakeArrayStrictlyIncreasing(nums []int) int`

> **Ide Kunci:** Scan from right to left to find the longest strictly increasing suffix.

## Solusi Go

```go
package main

// LeetCode #3818: Minimum Prefix Removal to Make Array Strictly Increasing
// https://leetcode.com/problems/minimum-prefix-removal-to-make-array-strictly-increasing/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Scan from right to left to find the longest strictly increasing suffix.

import "fmt"

func MinimumPrefixRemovalToMakeArrayStrictlyIncreasing(nums []int) int {
	n := len(nums)
	for i := n - 1; i > 0; i-- {
		if nums[i-1] >= nums[i] {
			return i
		}
	}
	return 0
}

func main() {
	// Example 1
	fmt.Println(MinimumPrefixRemovalToMakeArrayStrictlyIncreasing([]int{1, -1, 2, 3, 3, 4, 5})) // Expected: 4

	// Example 2
	fmt.Println(MinimumPrefixRemovalToMakeArrayStrictlyIncreasing([]int{4, 3, -2, -5})) // Expected: 3

	// Example 3
	fmt.Println(MinimumPrefixRemovalToMakeArrayStrictlyIncreasing([]int{1, 2, 3, 4})) // Expected: 0
}
```
