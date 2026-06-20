# 0611 — Valid Triangle Number

## Deskripsi

**Soal:** [0611. Valid Triangle Number](https://leetcode.com/problems/valid-triangle-number/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n^2)  
**Kompleksitas Ruang:** O(log n) for sorting

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #611: Valid Triangle Number
// https://leetcode.com/problems/valid-triangle-number/
// Difficulty: Medium
// Time: O(n^2)
// Space: O(log n) for sorting

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(TriangleNumber([]int{2, 2, 3, 4}))
	fmt.Println(TriangleNumber([]int{4, 2, 3, 4}))
}

func TriangleNumber(nums []int) int {
	sort.Ints(nums)
	count := 0
	n := len(nums)

	for i := n - 1; i >= 2; i-- {
		left, right := 0, i-1
  // Loop two-pointer: kiri vs kanan
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				count += right - left
				right--
			} else {
				left++
			}
		}
	}

	return count
}
```
