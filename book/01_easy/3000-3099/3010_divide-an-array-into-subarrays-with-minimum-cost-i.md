# 3010 — Divide An Array Into Subarrays With Minimum Cost I

## Deskripsi

**Soal:** [3010. Divide An Array Into Subarrays With Minimum Cost I](https://leetcode.com/problems/divide-an-array-into-subarrays-with-minimum-cost-i/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #3010: Divide an Array Into Subarrays With Minimum Cost I
// https://leetcode.com/problems/divide-an-array-into-subarrays-with-minimum-cost-i/
// Difficulty: Easy

import (
	"fmt"
	"sort"
)

func main() {
	// LeetCode name: minimumCost
	fmt.Println(DivideAnArrayIntoSubarraysWithMinimumCostI([]int{1, 2, 3, 12})) // 6
	fmt.Println(DivideAnArrayIntoSubarraysWithMinimumCostI([]int{5, 4, 3, 2, 1})) // 8
}

// Time: O(n log n) | Space: O(1)
// LeetCode submission name: minimumCost
// Cost = nums[0] + sum of two smallest elements from nums[1:]
func DivideAnArrayIntoSubarraysWithMinimumCostI(nums []int) int {
	// First subarray starts at nums[0], so nums[0] is always included.
	// For remaining subarrays, we pick the two smallest elements.
	rest := nums[1:]
	sort.Ints(rest)
	return nums[0] + rest[0] + rest[1]
}
```
