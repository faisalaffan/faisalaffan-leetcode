# 0410 — Split Array Largest Sum

## Deskripsi

**Soal:** [0410. Split Array Largest Sum](https://leetcode.com/problems/split-array-largest-sum/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Binary Search (pencarian biner)

## Solusi Go

```go
package main

// LeetCode #410: Split Array Largest Sum
// https://leetcode.com/problems/split-array-largest-sum/
// Difficulty: Hard
//
// Binary search on the answer. The minimum possible largest sum is max(nums)
// and the maximum is sum(nums). For each candidate mid, greedily check if
// we can split the array into k contiguous subarrays each with sum <= mid.

import (
	"fmt"
)

func main() {
	// Example 1: [7,2,5,10,8], k=2 -> 18 (split at 10: [7,2,5] sum=14, [10,8] sum=18)
	fmt.Println(splitArray([]int{7, 2, 5, 10, 8}, 2))
	// Example 2: [1,2,3,4,5], k=2 -> 9
	fmt.Println(splitArray([]int{1, 2, 3, 4, 5}, 2))
	// Example 3: [1,4,4], k=3 -> 4
	fmt.Println(splitArray([]int{1, 4, 4}, 3))
	// Edge: single element
	fmt.Println(splitArray([]int{10}, 1))
	// Edge: all same
	fmt.Println(splitArray([]int{1, 1, 1, 1, 1}, 3))
}

func splitArray(nums []int, k int) int {
	left, right := 0, 0
	for _, v := range nums {
		if v > left {
			left = v
		}
		right += v
	}

  // Loop two-pointer: kiri vs kanan
	for left < right {
		mid := left + (right-left)/2
		if feasible(nums, k, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}

func feasible(nums []int, k, target int) bool {
	count, sum := 1, 0
	for _, v := range nums {
		if sum+v > target {
			count++
			sum = v
			if count > k {
				return false
			} else {
				continue
			}
		}
		sum += v
	}
	return true
}
```
