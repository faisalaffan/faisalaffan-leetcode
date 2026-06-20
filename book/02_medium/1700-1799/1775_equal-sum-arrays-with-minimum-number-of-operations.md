# 1775 — Equal Sum Arrays With Minimum Number Of Operations

## Deskripsi

**Soal:** [1775. Equal Sum Arrays With Minimum Number Of Operations](https://leetcode.com/problems/equal-sum-arrays-with-minimum-number-of-operations/)

**Tingkat Kesulitan:** Sedang

**Kompleksitas Waktu:** O(n + m), Space: O(1)  
**Kompleksitas Ruang:** O(1)

**Algoritma:** —

**Fungsi Solusi:** `func minOperations(nums1 []int, nums2 []int) int`

## Solusi Go

```go
package main

// LeetCode #1775: Equal Sum Arrays With Minimum Number of Operations
// https://leetcode.com/problems/equal-sum-arrays-with-minimum-number-of-operations/
// Difficulty: Medium
// Time: O(n + m), Space: O(1)

import "fmt"

func minOperations(nums1 []int, nums2 []int) int {
	sum1, sum2 := 0, 0
	for _, v := range nums1 {
		sum1 += v
	}
	for _, v := range nums2 {
		sum2 += v
	}

	if sum1 == sum2 {
		return 0
	}

	// Make nums1 the one with smaller sum
	if sum1 > sum2 {
		nums1, nums2 = nums2, nums1
		sum1, sum2 = sum2, sum1
	}

	diff := sum2 - sum1
  // Membuat slice untuk menyimpan hasil
	count := make([]int, 7) // possible increments/decrements

	// nums1: smaller sum, we want to increase values (change to max 6)
	for _, v := range nums1 {
		count[6-v]++ // max possible increase
	}
	// nums2: larger sum, we want to decrease values (change to min 1)
	for _, v := range nums2 {
		count[v-1]++ // max possible decrease
	}

	ops := 0
	for i := 6; i >= 1; i-- {
		for count[i] > 0 && diff > 0 {
			diff -= i
			count[i]--
			ops++
		}
	}

	if diff > 0 {
		return -1
	}
	return ops
}

func main() {
	fmt.Println(minOperations([]int{1, 2, 3, 4, 5, 6}, []int{1, 1, 2, 2, 2, 2})) // Expected: 3
	fmt.Println(minOperations([]int{1, 1, 1, 1}, []int{6, 6, 6, 6})) // Expected: 4
	fmt.Println(minOperations([]int{6, 6}, []int{1})) // Expected: 3
}
```
