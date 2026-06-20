# 2563 — Count The Number Of Fair Pairs

## Deskripsi

**Soal:** [2563. Count The Number Of Fair Pairs](https://leetcode.com/problems/count-the-number-of-fair-pairs/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func countFairPairs(nums []int, lower int, upper int) int64`

## Solusi Go

```go
package main

// LeetCode #2563: Count the Number of Fair Pairs
// https://leetcode.com/problems/count-the-number-of-fair-pairs/
// Difficulty: Medium
// Time: O(n log n) | Space: O(1)

import (
	"fmt"
	"sort"
)

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)
	var ans int64
	n := len(nums)

	for i := 0; i < n; i++ {
		// Find lower bound for j > i such that nums[j] >= lower - nums[i]
		left := sort.Search(n, func(j int) bool {
			return j > i && nums[j] >= lower-nums[i]
		})
		// Find upper bound for j > i such that nums[j] <= upper - nums[i]
		right := sort.Search(n, func(j int) bool {
			return j > i && nums[j] > upper-nums[i]
		})
		if right > left {
			ans += int64(right - left)
		}
	}
	return ans
}

func main() {
	// Test case 1
	fmt.Println("Test 1:", countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6))
	// Expected: 6

	// Test case 2
	fmt.Println("Test 2:", countFairPairs([]int{1, 7, 9, 2, 5}, 11, 11))
	// Expected: 1

	// Test case 3
	fmt.Println("Test 3:", countFairPairs([]int{0, 0, 0, 0, 0, 0}, 0, 0))
	// Expected: 15
}
```
