# 0724 — Find Pivot Index

## Deskripsi

**Soal:** [0724. Find Pivot Index](https://leetcode.com/problems/find-pivot-index/)

**Tingkat Kesulitan:** Mudah

**Kompleksitas Waktu:** O(n). Space: O(1).  
**Kompleksitas Ruang:** O(1).

**Algoritma:** —

## Solusi Go

```go
package main

// LeetCode #724: Find Pivot Index
// https://leetcode.com/problems/find-pivot-index/
// Difficulty: Easy

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))   // 3
	fmt.Println(pivotIndex([]int{1, 2, 3}))             // -1
	fmt.Println(pivotIndex([]int{2, 1, -1}))            // 0
}

// pivotIndex finds the index where sum of left elements equals sum of right elements.
// Time: O(n). Space: O(1).
func pivotIndex(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	leftSum := 0
	for i, v := range nums {
		if leftSum == total-leftSum-v {
			return i
		}
		leftSum += v
	}
	return -1
}
```
