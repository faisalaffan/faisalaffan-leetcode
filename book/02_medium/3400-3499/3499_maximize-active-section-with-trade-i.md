# 3499 — Maximize Active Section With Trade I

## Deskripsi

**Soal:** [3499. Maximize Active Section With Trade I](https://leetcode.com/problems/maximize-active-section-with-trade-i/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3499: Maximize Active Section with Trade I
// https://leetcode.com/problems/maximize-active-section-with-trade-i/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MaximizeActiveSectionWithTradeI([]int{0, 1, 1, 0, 1}))
	// Test case 2
	fmt.Println("Test 2:", MaximizeActiveSectionWithTradeI([]int{1, 0, 1, 0, 1}))
	// Test case 3
	fmt.Println("Test 3:", MaximizeActiveSectionWithTradeI([]int{0, 0, 1, 1}))
}

func MaximizeActiveSectionWithTradeI(arr []int) int {
	// Find longest contiguous segment of 1s
	maxLen := 0
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 1 {
			j := i
			for j < n && arr[j] == 1 {
				j++
			}
			if j-i > maxLen {
				maxLen = j - i
			}
			i = j
		}
	}
	// Allow one 0 to be flipped to connect two segments
	// Find two 1-segments separated by a single 0
	for i := 1; i < n-1; i++ {
		if arr[i] == 0 && arr[i-1] == 1 && arr[i+1] == 1 {
			left := i - 1
			for left >= 0 && arr[left] == 1 {
				left--
			}
			right := i + 1
			for right < n && arr[right] == 1 {
				right++
			}
			merged := (i - 1 - left) + (right - i - 1) + 1
			if merged > maxLen {
				maxLen = merged
			}
		}
	}
	return maxLen
}
```
