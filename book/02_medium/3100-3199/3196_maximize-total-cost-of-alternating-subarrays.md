# 3196 — Maximize Total Cost Of Alternating Subarrays

## Deskripsi

**Soal:** [3196. Maximize Total Cost Of Alternating Subarrays](https://leetcode.com/problems/maximize-total-cost-of-alternating-subarrays/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func maximumTotalCost(nums []int) int64`

## Solusi Go

```go
package main

// LeetCode #3196: Maximize Total Cost of Alternating Subarrays
// https://leetcode.com/problems/maximize-total-cost-of-alternating-subarrays/
// Difficulty: Medium
// Time: O(n) | Space: O(1)

import "fmt"

func maximumTotalCost(nums []int) int64 {
  // Edge case: input kosong
	if len(nums) == 0 {
		return 0
	}

	add := int64(nums[0])
	sub := int64(nums[0])

	for i := 1; i < len(nums); i++ {
		v := int64(nums[i])
		newAdd := maxInt64(add, sub) + v
		newSub := add - v
		add, sub = newAdd, newSub
	}

	return maxInt64(add, sub)
}

func maxInt64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maximumTotalCost([]int{1, -2, 3, 4}))    // Expected: 10
	fmt.Println(maximumTotalCost([]int{1, -1, 1, -1}))   // Expected: 4
}
```
