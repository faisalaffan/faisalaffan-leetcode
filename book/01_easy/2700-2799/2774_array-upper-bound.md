# 2774 — Array Upper Bound

## Deskripsi

**Soal:** [2774. Array Upper Bound](https://leetcode.com/problems/array-upper-bound/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(log n)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #2774: Array Upper Bound
// https://leetcode.com/problems/array-upper-bound/
// Difficulty: Easy [Paid]
// Time: O(log n) | Space: O(1)
// Note: JS problem, adapted to Go. Returns upper bound of target in sorted array.

import "fmt"

func main() {
	fmt.Println(ArrayUpperBound([]int{1, 2, 2, 2, 3}, 2))
	fmt.Println(ArrayUpperBound([]int{1, 3, 5}, 4))
}

func ArrayUpperBound(nums []int, target int) int {
	left, right := 0, len(nums)
  // Loop two-pointer: kiri vs kanan
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] <= target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left - 1
}
```
