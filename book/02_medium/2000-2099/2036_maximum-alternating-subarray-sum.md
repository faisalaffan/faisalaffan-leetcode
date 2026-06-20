# 2036 — Maximum Alternating Subarray Sum

## Deskripsi

**Soal:** [2036. Maximum Alternating Subarray Sum](https://leetcode.com/problems/maximum-alternating-subarray-sum/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** Dynamic Programming (DP)

**Fungsi Solusi:** `func maximumAlternatingSubarraySum(nums []int) int64`

## Solusi Go

```go
package main

// LeetCode #2036: Maximum Alternating Subarray Sum
// https://leetcode.com/problems/maximum-alternating-subarray-sum/
// Difficulty: Medium [Paid]
// Time: O(n) | Space: O(1)

import "fmt"

func maximumAlternatingSubarraySum(nums []int) int64 {
	// dp0: max sum ending at i with the current element as positive (even index in subarray)
	// dp1: max sum ending at i with the current element as negative (odd index in subarray)
	dp0 := int64(nums[0])
	dp1 := int64(-1 << 62) // very negative
	result := dp0

	for i := 1; i < len(nums); i++ {
		// nums[i] as positive: either start new, or continue from dp1
		newDp0 := max64(int64(nums[i]), dp1+int64(nums[i]))
		// nums[i] as negative: must continue from dp0
		newDp1 := dp0 - int64(nums[i])

		dp0, dp1 = newDp0, newDp1
		result = max64(result, dp0)
	}

	return result
}

func max64(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", maximumAlternatingSubarraySum([]int{4, 2, 5, 3}))
	// Expected: 7

	// Test case 2
	fmt.Println("Test 2:", maximumAlternatingSubarraySum([]int{5, 6, 7, 8}))
	// Expected: 8

	// Test case 3
	fmt.Println("Test 3:", maximumAlternatingSubarraySum([]int{-1, -2, -3}))
	// Expected: -1
}
```
