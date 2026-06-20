# 0918 — Maximum Sum Circular Subarray

## Deskripsi

**Soal:** [0918. Maximum Sum Circular Subarray](https://leetcode.com/problems/maximum-sum-circular-subarray/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #918: Maximum Sum Circular Subarray
// https://leetcode.com/problems/maximum-sum-circular-subarray/
// Difficulty: Medium

import "fmt"

func main() {
	fmt.Println(MaximumSumCircularSubarray([]int{1, -2, 3, -2}))
	fmt.Println(MaximumSumCircularSubarray([]int{5, -3, 5}))
	fmt.Println(MaximumSumCircularSubarray([]int{-3, -2, -3}))
}

// Time: O(n) | Space: O(1)
func MaximumSumCircularSubarray(nums []int) int {
	total := 0
	maxSum, curMax := nums[0], 0
	minSum, curMin := nums[0], 0

	for _, v := range nums {
		total += v
		curMax = max(curMax+v, v)
		maxSum = max(maxSum, curMax)
		curMin = min(curMin+v, v)
		minSum = min(minSum, curMin)
	}

	// If all numbers are negative, return the max (non-circular)
	if maxSum < 0 {
		return maxSum
	}

	return max(maxSum, total-minSum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
```
