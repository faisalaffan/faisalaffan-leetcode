# 1493 — Longest Subarray Of 1S After Deleting One Element

## Deskripsi

**Soal:** [1493. Longest Subarray Of 1S After Deleting One Element](https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(N), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Sliding Window (jendela geser)

## Solusi Go

```go
package main

// LeetCode #1493: Longest Subarray of 1's After Deleting One Element
// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(LongestSubarray([]int{1, 1, 0, 1}))
	fmt.Println(LongestSubarray([]int{0, 1, 1, 1, 0, 1, 1, 0, 1}))
	fmt.Println(LongestSubarray([]int{1, 1, 1}))
}

func LongestSubarray(nums []int) int {
	// Time: O(N), Space: O(1)
	// Sliding window with at most one zero
	left := 0
	zeroCount := 0
	maxLen := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}

		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}

		// Window length minus the one element we must delete
		currLen := right - left
		if currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen
}
```
