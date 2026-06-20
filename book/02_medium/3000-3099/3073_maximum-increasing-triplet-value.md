# 3073 — Maximum Increasing Triplet Value

## Deskripsi

**Soal:** [3073. Maximum Increasing Triplet Value](https://leetcode.com/problems/maximum-increasing-triplet-value/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3073: Maximum Increasing Triplet Value (PAID)
// https://leetcode.com/problems/maximum-increasing-triplet-value/
// Difficulty: Medium [Paid]
// Time: O(n^2) | Space: O(1)

// Find the maximum value of nums[i] - nums[j] + nums[k] where
// i < j < k and nums[i] < nums[j] < nums[k]. At least one valid triplet
// is guaranteed to exist.

import "fmt"

func main() {
	// Test 1: Strictly increasing
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(maximumIncreasingTripletValue(nums)) // 4  (3-4+5 or 2-3+5 or 1-2+5)

	// Test 2: Decreasing then increasing
	nums2 := []int{3, 2, 1, 4, 5}
	fmt.Println(maximumIncreasingTripletValue(nums2)) // 4 (3-4+5)

	// Test 3: Small array
	nums3 := []int{10, 20, 30}
	fmt.Println(maximumIncreasingTripletValue(nums3)) // 20 (10-20+30)

	// Test 4: Duplicates — must be strictly increasing
	nums4 := []int{5, 5, 5, 5}
	fmt.Println(maximumIncreasingTripletValue(nums4)) // 0 (no strictly increasing triplet)
}

func maximumIncreasingTripletValue(nums []int) int {
	n := len(nums)
	ans := 0
	for j := 1; j < n-1; j++ {
		leftBest := -1
		for i := 0; i < j; i++ {
			if nums[i] < nums[j] && nums[i] > leftBest {
				leftBest = nums[i]
			}
		}
		rightBest := -1
		for k := j + 1; k < n; k++ {
			if nums[k] > nums[j] && nums[k] > rightBest {
				rightBest = nums[k]
			}
		}
		if leftBest != -1 && rightBest != -1 {
			val := leftBest - nums[j] + rightBest
			if val > ans {
				ans = val
			}
		}
	}
	return ans
}
```
