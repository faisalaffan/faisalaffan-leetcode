# 0081 — Search In Rotated Sorted Array Ii

## Deskripsi

**Soal:** [0081. Search In Rotated Sorted Array Ii](https://leetcode.com/problems/search-in-rotated-sorted-array-ii/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(log n) average, O(n) worst  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func search(nums []int, target int) bool`

## Solusi Go

```go
package main

// LeetCode #81: Search in Rotated Sorted Array II
// https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
// Difficulty: Medium

import "fmt"

func search(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return true
		}

		// Handle duplicates: shrink window
		if nums[left] == nums[mid] && nums[mid] == nums[right] {
			left++
			right--
		} else if nums[left] <= nums[mid] {
			// Left half is sorted
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// Right half is sorted
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}

func main() {
	// Test case 1
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 0)) // true

	// Test case 2
	fmt.Println(search([]int{2, 5, 6, 0, 0, 1, 2}, 3)) // false

	// Test case 3
	fmt.Println(search([]int{1, 0, 1, 1, 1}, 0)) // true
}

// Time: O(log n) average, O(n) worst | Space: O(1)
```
