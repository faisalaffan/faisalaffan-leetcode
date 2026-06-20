# 0215 — Kth Largest Element In An Array

## Deskripsi

**Soal:** [0215. Kth Largest Element In An Array](https://leetcode.com/problems/kth-largest-element-in-an-array/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n) average with quickselect, O(n log n) worst case, Space: O(log n)  
**Kompleksitas Ruang:** O(log n)

**Algoritma:** —

**Fungsi Solusi:** `func findKthLargest(nums []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #215: Kth Largest Element in an Array
// https://leetcode.com/problems/kth-largest-element-in-an-array/
// Difficulty: Medium
// Time: O(n) average with quickselect, O(n log n) worst case, Space: O(log n)

import "fmt"

func findKthLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

func quickSelect(nums []int, left, right, kSmallest int) int {
	if left == right {
		return nums[left]
	}

	pivotIdx := partition(nums, left, right)

	if kSmallest == pivotIdx {
		return nums[kSmallest]
	} else if kSmallest < pivotIdx {
		return quickSelect(nums, left, pivotIdx-1, kSmallest)
	}
	return quickSelect(nums, pivotIdx+1, right, kSmallest)
}

func partition(nums []int, left, right int) int {
	pivot := nums[right]
	storeIdx := left

	for i := left; i < right; i++ {
		if nums[i] < pivot {
			nums[storeIdx], nums[i] = nums[i], nums[storeIdx]
			storeIdx++
		}
	}

	nums[storeIdx], nums[right] = nums[right], nums[storeIdx]
	return storeIdx
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2))
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
	fmt.Println(findKthLargest([]int{1}, 1))
}
```
