# 3914 — Minimum Operations To Make Array Non Decreasing

## Deskripsi

**Soal:** [3914. Minimum Operations To Make Array Non Decreasing](https://leetcode.com/problems/minimum-operations-to-make-array-non-decreasing/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Greedy (pemilihan optimal lokal)

**Fungsi Solusi:** `func MinimumOperationsToMakeArrayNonDecreasing(nums []int) int64`

> **Ide Kunci:** Greedy. When nums[i] < nums[i-1], need to increase a suffix.

## Solusi Go

```go
package main

// LeetCode #3914: Minimum Operations to Make Array Non Decreasing
// https://leetcode.com/problems/minimum-operations-to-make-array-non-decreasing/
// Difficulty: Medium
// Time: O(N) | Space: O(1)
// Approach: Greedy. When nums[i] < nums[i-1], need to increase a suffix.
// Track operations with running max.

import "fmt"

func MinimumOperationsToMakeArrayNonDecreasing(nums []int) int64 {
	var ans int64 = 0
	mx := 0
	for _, v := range nums {
		if v >= mx {
			mx = v
		} else {
			ans += int64(mx - v)
		}
	}
	return ans
}

func main() {
	// Example
	fmt.Println(MinimumOperationsToMakeArrayNonDecreasing([]int{1, 2, 3}))        // Expected: 0
	fmt.Println(MinimumOperationsToMakeArrayNonDecreasing([]int{3, 2, 1}))        // Expected: 3
	fmt.Println(MinimumOperationsToMakeArrayNonDecreasing([]int{1, 2, 1, 2, 1})) // Expected: 2
}
```
