# 3828 — Final Element After Subarray Deletions

## Deskripsi

**Soal:** [3828. Final Element After Subarray Deletions](https://leetcode.com/problems/final-element-after-subarray-deletions/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func FinalElementAfterSubarrayDeletions(nums []int) int`

> **Ide Kunci:** Alice (first) can always keep first or last. Optimal play yields max(first, last).

## Solusi Go

```go
package main

// LeetCode #3828: Final Element After Subarray Deletions
// https://leetcode.com/problems/final-element-after-subarray-deletions/
// Difficulty: Medium
// Time: O(1) | Space: O(1)
// Approach: Alice (first) can always keep first or last. Optimal play yields max(first, last).

import "fmt"

func FinalElementAfterSubarrayDeletions(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if nums[0] > nums[len(nums)-1] {
		return nums[0]
	}
	return nums[len(nums)-1]
}

func main() {
	// Example 1
	fmt.Println(FinalElementAfterSubarrayDeletions([]int{1, 5, 2})) // Expected: 2

	// Example 2
	fmt.Println(FinalElementAfterSubarrayDeletions([]int{3, 7})) // Expected: 7

	// Example 3
	fmt.Println(FinalElementAfterSubarrayDeletions([]int{5})) // Expected: 5
}
```
