# 2576 — Find The Maximum Number Of Marked Indices

## Deskripsi

**Soal:** [2576. Find The Maximum Number Of Marked Indices](https://leetcode.com/problems/find-the-maximum-number-of-marked-indices/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maxNumOfMarkedIndices(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #2576: Find the Maximum Number of Marked Indices
// https://leetcode.com/problems/find-the-maximum-number-of-marked-indices/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func maxNumOfMarkedIndices(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	left, right := 0, n/2
	count := 0

	for left < n/2 && right < n {
		if 2*nums[left] <= nums[right] {
			count += 2
			left++
			right++
		} else {
			right++
		}
	}
	return count
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maxNumOfMarkedIndices([]int{3, 5, 2, 4}))
	// Expected: 2

	// Test case 2
	fmt.Println("Test 2:", maxNumOfMarkedIndices([]int{9, 2, 5, 4}))
	// Expected: 4

	// Test case 3
	fmt.Println("Test 3:", maxNumOfMarkedIndices([]int{7, 6, 8}))
	// Expected: 0
}
```
