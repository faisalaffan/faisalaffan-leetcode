# 3698 — Split Array With Minimum Difference

## Deskripsi

**Soal:** [3698. Split Array With Minimum Difference](https://leetcode.com/problems/split-array-with-minimum-difference/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(n)

**Algoritma:** —

**Fungsi Solusi:** `func splitArrayWithMinimumDifference(nums []int) int64`

## Solusi Go

```go
package main

// LeetCode #3698: Split Array With Minimum Difference
// https://leetcode.com/problems/split-array-with-minimum-difference/
// Difficulty: Medium
// Time: O(n) | Space: O(n)

import (
	"fmt"
	"math"
)

func splitArrayWithMinimumDifference(nums []int) int64 {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	prefix := make([]int64, n)
	prefix[0] = int64(nums[0])
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + int64(nums[i])
	}
	total := prefix[n-1]

  // Membuat slice untuk menyimpan hasil
	inc := make([]bool, n)
	inc[0] = true
	for i := 1; i < n; i++ {
		inc[i] = inc[i-1] && nums[i] > nums[i-1]
	}

  // Membuat slice untuk menyimpan hasil
	dec := make([]bool, n)
	dec[n-1] = true
	for i := n - 2; i >= 0; i-- {
		dec[i] = dec[i+1] && nums[i] > nums[i+1]
	}

	minDiff := int64(math.MaxInt64)
	for i := 0; i < n-1; i++ {
		if inc[i] && dec[i+1] {
			leftSum := prefix[i]
			rightSum := total - prefix[i]
			diff := leftSum - rightSum
			if diff < 0 {
				diff = -diff
			}
			if diff < minDiff {
				minDiff = diff
			}
		}
	}

	if minDiff == int64(math.MaxInt64) {
		return -1
	}
	return minDiff
}

func main() {
	fmt.Println(splitArrayWithMinimumDifference([]int{1, 3, 2}))
	fmt.Println(splitArrayWithMinimumDifference([]int{1, 2, 4, 3}))
	fmt.Println(splitArrayWithMinimumDifference([]int{3, 1, 2}))
}
```
