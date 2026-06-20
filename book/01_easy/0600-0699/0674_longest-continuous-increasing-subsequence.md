# 0674 — Longest Continuous Increasing Subsequence

## Deskripsi

**Soal:** [0674. Longest Continuous Increasing Subsequence](https://leetcode.com/problems/longest-continuous-increasing-subsequence/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #674: Longest Continuous Increasing Subsequence
// https://leetcode.com/problems/longest-continuous-increasing-subsequence/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 4, 7}))    // 3
	fmt.Println(findLengthOfLCIS([]int{2, 2, 2, 2, 2}))    // 1
	fmt.Println(findLengthOfLCIS([]int{1, 3, 5, 7}))       // 4
}

// findLengthOfLCIS finds the length of the longest continuous increasing subsequence.
// Time: O(n). Space: O(1).
func findLengthOfLCIS(nums []int) int {
  // Edge case: input kosong
	if len(nums) == 0 {
		return 0
	}
	maxLen, curr := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			curr++
			if curr > maxLen {
				maxLen = curr
			}
		} else {
			curr = 1
		}
	}
	return maxLen
}
```
