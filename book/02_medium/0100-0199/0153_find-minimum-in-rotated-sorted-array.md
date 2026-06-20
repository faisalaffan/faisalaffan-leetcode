# 0153 — Find Minimum In Rotated Sorted Array

## Deskripsi

**Soal:** [0153. Find Minimum In Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log n), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func findMin(nums []int) int`

## Solusi Go

```go
package main

// LeetCode #153: Find Minimum in Rotated Sorted Array
// https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
// Difficulty: Medium
// Time: O(log n), Space: O(1)

import "fmt"

func findMin(nums []int) int {
	left, right := 0, len(nums)-1

  // Loop two-pointer: kiri vs kanan
	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}

func main() {
	fmt.Println(findMin([]int{3, 4, 5, 1, 2}))
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
}
```
