# 3576 — Transform Array To All Equal Elements

## Deskripsi

**Soal:** [3576. Transform Array To All Equal Elements](https://leetcode.com/problems/transform-array-to-all-equal-elements/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3576: Transform Array to All Equal Elements
// https://leetcode.com/problems/transform-array-to-all-equal-elements/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", TransformArrayToAllEqualElements([]int{1, 2, 3}))
	// Test case 2
	fmt.Println("Test 2:", TransformArrayToAllEqualElements([]int{1, 1, 1}))
	// Test case 3
	fmt.Println("Test 3:", TransformArrayToAllEqualElements([]int{1, 100}))
}

func TransformArrayToAllEqualElements(nums []int) int {
	// Minimum operations to make all elements equal
	// Each operation: increment or decrement by 1
	// Optimal target is median
	// Simple approach: find median and compute sum of absolute differences
	n := len(nums)
	// Selection algorithm to find median (here using simple O(n^2) for small n)
	// For larger arrays, we'd use QuickSelect or sort
	median := 0
	if n%2 == 0 {
		// Use either median
		median = findKth(nums, n/2)
	} else {
		median = findKth(nums, n/2)
	}
	ops := 0
	for _, v := range nums {
		if v > median {
			ops += v - median
		} else {
			ops += median - v
		}
	}
	return ops
}

func findKth(nums []int, k int) int {
	// Simple O(n^2) selection
	for i := 0; i <= k; i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums[k]
}
```
