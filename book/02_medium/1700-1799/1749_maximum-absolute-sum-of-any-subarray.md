# 1749 — Maximum Absolute Sum Of Any Subarray

## Deskripsi

**Soal:** [1749. Maximum Absolute Sum Of Any Subarray](https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maxAbsoluteSum(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #1749: Maximum Absolute Sum of Any Subarray
// https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/
// Difficulty: Medium
// Time: O(n), Space: O(1)

import "fmt"

func maxAbsoluteSum(nums []int) int {
	maxEnding := 0
	minEnding := 0
	maxSoFar := 0

	for _, v := range nums {
		maxEnding = max(0, maxEnding+v)
		minEnding = min(0, minEnding+v)
		maxSoFar = max(maxSoFar, maxEnding, -minEnding)
	}
	return maxSoFar
}

func max(nums ...int) int {
	r := nums[0]
	for _, v := range nums[1:] {
		if v > r {
			r = v
		}
	}
	return r
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxAbsoluteSum([]int{1, -3, 2, 3, -4})) // Expected: 5
	fmt.Println(maxAbsoluteSum([]int{2, -5, 1, -4, 3, -2})) // Expected: 8
	fmt.Println(maxAbsoluteSum([]int{-1})) // Expected: 1
}
```
