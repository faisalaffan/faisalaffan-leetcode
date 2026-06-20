# 0154 — Find Minimum In Rotated Sorted Array Ii

## Deskripsi

**Soal:** [0154. Find Minimum In Rotated Sorted Array Ii](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** —

**Fungsi Solusi:** `func findMin(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #154: Find Minimum in Rotated Sorted Array II
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/
// Difficulty: Hard

import (
	"fmt"
)

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

  // Loop two-pointer: kiri vs kanan
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			// Minimum is in the right half
			left = mid + 1
		} else if nums[mid] < nums[right] {
			// Minimum is in the left half (including mid)
			right = mid
		} else {
			// nums[mid] == nums[right], cannot determine, shrink
			right--
		}
	}

	return nums[left]
}

func main() {
	nums := []int{2, 2, 2, 0, 1}
	result := findMin(nums)
	expected := 0

	fmt.Printf("findMin(%v) = %d\n", nums, result)
	if result == expected {
		fmt.Println("PASS")
	} else {
		fmt.Printf("FAIL: expected %d\n", expected)
	}
}
```
