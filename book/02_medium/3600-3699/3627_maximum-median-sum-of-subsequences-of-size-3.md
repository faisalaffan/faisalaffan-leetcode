# 3627 — Maximum Median Sum Of Subsequences Of Size 3

## Deskripsi

**Soal:** [3627. Maximum Median Sum Of Subsequences Of Size 3](https://leetcode.com/problems/maximum-median-sum-of-subsequences-of-size-3/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3627: Maximum Median Sum of Subsequences of Size 3
// https://leetcode.com/problems/maximum-median-sum-of-subsequences-of-size-3/
// Difficulty: Medium
// Complexity: O(n^3) time, O(1) space

import (
	"fmt"
	"sort"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", MaximumMedianSumOfSubsequencesOfSizeThree([]int{1, 2, 3, 4, 5}))
	// Test case 2
	fmt.Println("Test 2:", MaximumMedianSumOfSubsequencesOfSizeThree([]int{5, 1, 5, 1, 5}))
	// Test case 3
	fmt.Println("Test 3:", MaximumMedianSumOfSubsequencesOfSizeThree([]int{1, 2, 3}))
}

func MaximumMedianSumOfSubsequencesOfSizeThree(nums []int) int {
	n := len(nums)
	maxSum := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				sub := []int{nums[i], nums[j], nums[k]}
				sort.Ints(sub)
				sum := sub[0] + sub[1] + sub[2]
				if sum > maxSum {
					maxSum = sum
				}
			}
		}
	}
	return maxSum
}
```
