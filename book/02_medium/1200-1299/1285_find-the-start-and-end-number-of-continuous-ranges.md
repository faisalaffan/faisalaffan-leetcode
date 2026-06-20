# 1285 — Find The Start And End Number Of Continuous Ranges

## Deskripsi

**Soal:** [1285. Find The Start And End Number Of Continuous Ranges](https://leetcode.com/problems/find-the-start-and-end-number-of-continuous-ranges/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** LIS (Longest Increasing Subsequence)

**Fungsi Solusi:** `func findContinuousRanges(nums []int) [][]int`

## Solusi Go

```go
package main

import (
	"fmt"
	"sort"
)

// LeetCode #1285: Find the Start and End Number of Continuous Ranges
// https://leetcode.com/problems/find-the-start-and-end-number-of-continuous-ranges/
// Difficulty: Medium [Paid]

// Given a sorted list of unique integers, find continuous ranges.

// Time: O(n)
// Space: O(n)

func findContinuousRanges(nums []int) [][]int {
  // Edge case: input kosong
	if len(nums) == 0 {
		return [][]int{}
	}

  // Membuat slice 2D untuk DP/tabel
	result := make([][]int, 0)
	start := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			result = append(result, []int{start, nums[i-1]})
			start = nums[i]
		}
	}
	result = append(result, []int{start, nums[len(nums)-1]})

	sort.Slice(result, func(i, j int) bool {
		return result[i][0] < result[j][0]
	})

	return result
}

func main() {
	fmt.Printf("%v (expected: [[1 3] [6 7] [9 9]])\n",
		findContinuousRanges([]int{1, 2, 3, 6, 7, 9}))

	fmt.Printf("%v (expected: [[1 1]])\n",
		findContinuousRanges([]int{1}))

	fmt.Printf("%v (expected: [])\n",
		findContinuousRanges([]int{}))
}
```
