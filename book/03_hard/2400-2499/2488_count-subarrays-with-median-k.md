# 2488 — Count Subarrays With Median K

## Deskripsi

**Soal:** [2488. Count Subarrays With Median K](https://leetcode.com/problems/count-subarrays-with-median-k/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Prefix Sum (jumlah kumulatif)

## Solusi Go

```go
package main

// LeetCode #2488: Count Subarrays With Median K
// https://leetcode.com/problems/count-subarrays-with-median-k/
// Difficulty: Hard
//
// Transform array: >k => 1, <k => -1, =k => 0.
// A subarray has median k if it contains k and sum == 0 (odd length) or sum == 1 (even length).
// Find position p of k. Compute prefix balances left and right, use frequency maps.

import "fmt"

func main() {
	// Example 1: [3,2,1,4,5], 4 => 3
	fmt.Println(countSubarrays([]int{3, 2, 1, 4, 5}, 4))
	// Example 2: [2,3,1], 3 => 1
	fmt.Println(countSubarrays([]int{2, 3, 1}, 3))
	// Edge: single element
	fmt.Println(countSubarrays([]int{1}, 1))
	// Edge: k is first
	fmt.Println(countSubarrays([]int{5, 1, 2, 3, 4}, 5))
	// Edge: k at end
	fmt.Println(countSubarrays([]int{1, 2, 3, 4, 5}, 5))
}

func countSubarrays(nums []int, k int) int {
	n := len(nums)

	// Find position of k
	pos := -1
	for i, v := range nums {
		if v == k {
			pos = i
			break
		}
	}
	if pos == -1 {
		return 0
	}

	// Transform: >k => 1, <k => -1, =k => 0
	// Compute prefix sums from pos going right
  // Membuat map untuk pencarian O(1): key → value
	rightFreq := make(map[int]int)
	rightFreq[0] = 1 // empty suffix
	balance := 0
	for i := pos + 1; i < n; i++ {
		if nums[i] > k {
			balance++
		} else if nums[i] < k {
			balance--
		}
		rightFreq[balance]++
	}

	// Now go left, tracking left balance
  // Membuat map untuk pencarian O(1): key → value
	leftFreq := make(map[int]int)
	leftFreq[0] = 1 // empty prefix
	balance = 0
	for i := pos - 1; i >= 0; i-- {
		if nums[i] > k {
			balance++
		} else if nums[i] < k {
			balance--
		}
		leftFreq[balance]++
	}

	// Count subarrays where left_balance + right_balance == 0 (odd length, median at position)
	// or left_balance + right_balance == 1 (even length)
	result := 0
	for lb, lcnt := range leftFreq {
		// For odd length (1 + odd + even = odd), need sum == 0
		if rcnt, ok := rightFreq[-lb]; ok {
			result += lcnt * rcnt
		}
		// For even length, need sum == 1
		if rcnt, ok := rightFreq[1-lb]; ok {
			result += lcnt * rcnt
		}
	}

	return result
}
```
