# 1671 — Minimum Number Of Removals To Make Mountain Array

## Deskripsi

**Soal:** [1671. Minimum Number Of Removals To Make Mountain Array](https://leetcode.com/problems/minimum-number-of-removals-to-make-mountain-array/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** LIS (Longest Increasing Subsequence)

## Solusi Go

```go
package main

// LeetCode #1671: Minimum Number of Removals to Make Mountain Array
// https://leetcode.com/problems/minimum-number-of-removals-to-make-mountain-array/
// Difficulty: Hard

import "fmt"

func main() {
	nums1 := []int{1, 3, 1}
	fmt.Printf("Test 1 - Input: %v\nExpected: 0\nGot: %d\n\n", nums1, minimumMountainRemovals(nums1))

	nums2 := []int{2, 1, 1, 5, 6, 2, 3, 1}
	fmt.Printf("Test 2 - Input: %v\nExpected: 3\nGot: %d\n\n", nums2, minimumMountainRemovals(nums2))

	nums3 := []int{4, 3, 2, 1, 1, 2, 3, 1}
	fmt.Printf("Test 3 - Input: %v\nExpected: 4\nGot: %d\n", nums3, minimumMountainRemovals(nums3))
}

func minimumMountainRemovals(nums []int) int {
	n := len(nums)
  // Membuat slice untuk menyimpan hasil
	lis := make([]int, n) // longest increasing subsequence ending at i
  // Membuat slice untuk menyimpan hasil
	lds := make([]int, n) // longest decreasing subsequence starting at i

	// Compute LIS from left
	for i := 0; i < n; i++ {
		lis[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && lis[j]+1 > lis[i] {
				lis[i] = lis[j] + 1
			}
		}
	}

	// Compute LDS from right (reverse LIS)
	for i := n - 1; i >= 0; i-- {
		lds[i] = 1
		for j := n - 1; j > i; j-- {
			if nums[j] < nums[i] && lds[j]+1 > lds[i] {
				lds[i] = lds[j] + 1
			}
		}
	}

	// Find the longest bitonic subsequence (mountain)
	maxMountain := 0
	for i := 1; i < n-1; i++ {
		if lis[i] > 1 && lds[i] > 1 {
			mountainLen := lis[i] + lds[i] - 1
			if mountainLen > maxMountain {
				maxMountain = mountainLen
			}
		}
	}

	return n - maxMountain
}
```
