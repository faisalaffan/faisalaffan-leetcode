# 3523 — Make Array Non Decreasing

## Deskripsi

**Soal:** [3523. Make Array Non Decreasing](https://leetcode.com/problems/make-array-non-decreasing/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3523: Make Array Non-decreasing
// https://leetcode.com/problems/make-array-non-decreasing/
// Difficulty: Medium
// Complexity: O(n) time, O(1) space

import "fmt"

func main() {
	// Test case 1
	fmt.Println("Test 1:", MakeArrayNonDecreasing([]int{4, 2, 3}))
	// Test case 2
	fmt.Println("Test 2:", MakeArrayNonDecreasing([]int{4, 2, 1}))
	// Test case 3
	fmt.Println("Test 3:", MakeArrayNonDecreasing([]int{1, 2, 3}))
}

func MakeArrayNonDecreasing(nums []int) int {
	ops := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			ops += nums[i-1] - nums[i]
			nums[i] = nums[i-1]
		}
	}
	return ops
}
```
