# 1877 — Minimize Maximum Pair Sum In Array

## Deskripsi

**Soal:** [1877. Minimize Maximum Pair Sum In Array](https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #1877: Minimize Maximum Pair Sum in Array
// https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/
// Difficulty: Medium

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(MinPairSum([]int{3, 5, 2, 3}))
	fmt.Println(MinPairSum([]int{3, 5, 4, 2, 4, 6}))
}

// Time: O(n log n), Space: O(1)
func MinPairSum(nums []int) int {
	sort.Ints(nums)
	maxSum := 0
	n := len(nums)
	for i := 0; i < n/2; i++ {
		pairSum := nums[i] + nums[n-1-i]
		if pairSum > maxSum {
			maxSum = pairSum
		}
	}
	return maxSum
}
```
