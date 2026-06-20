# 3774 — Absolute Difference Between Maximum And Minimum K Elements

## Deskripsi

**Soal:** [3774. Absolute Difference Between Maximum And Minimum K Elements](https://leetcode.com/problems/absolute-difference-between-maximum-and-minimum-k-elements/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3774: Absolute Difference Between Maximum and Minimum K Elements
// https://leetcode.com/problems/absolute-difference-between-maximum-and-minimum-k-elements/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(AbsoluteDifferenceBetweenMaximumAndMinimumKElements([]int{5, 2, 2, 4}, 2))
	fmt.Println(AbsoluteDifferenceBetweenMaximumAndMinimumKElements([]int{100}, 1))
}

// Time: O(n log n)
// Space: O(1)
func AbsoluteDifferenceBetweenMaximumAndMinimumKElements(nums []int, k int) int {
	sort.Ints(nums)
	n := len(nums)
	maxSum := 0
	minSum := 0
	for i := 0; i < k; i++ {
		minSum += nums[i]
		maxSum += nums[n-1-i]
	}
	return maxSum - minSum
}
```
