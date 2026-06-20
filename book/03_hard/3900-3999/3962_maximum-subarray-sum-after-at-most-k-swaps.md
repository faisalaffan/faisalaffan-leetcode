# 3962 — Maximum Subarray Sum After At Most K Swaps

## Deskripsi

**Soal:** [3962. Maximum Subarray Sum After At Most K Swaps](https://leetcode.com/problems/maximum-subarray-sum-after-at-most-k-swaps/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Greedy (pemilihan optimal lokal)

> **Ide Kunci:** Greedy. Sort array descending. Take the largest

## Solusi Go

```go
package main

// LeetCode #3962: Maximum Subarray Sum After at Most K Swaps
// https://leetcode.com/problems/maximum-subarray-sum-after-at-most-k-swaps/
// Difficulty: Hard
//
// At most K swaps between any two elements. Maximize the sum of a
// contiguous subarray after performing swaps.
//
// Approach: Greedy. Sort array descending. Take the largest
// elements that can be brought into a single subarray. Use
// Kadane's algorithm after allowing up to K elements to be
// swapped into optimal positions. For small K, try swapping
// large outside elements into the best subarray.

import (
	"fmt"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(maximumSubarraySum([]int{5, 1, 2, 3, 4}, 2))
	// Example 2
	fmt.Println(maximumSubarraySum([]int{1, 9, 2, 8, 3}, 1))
	// Edge: k = 0
	fmt.Println(maximumSubarraySum([]int{1, -2, 3, 4, -5}, 0))
}

func maximumSubarraySum(nums []int, k int) int {
	n := len(nums)
  // Edge case: input kosong
	if n == 0 {
		return 0
	}

	// Kadane for baseline (no swaps)
	best := nums[0]
	cur := nums[0]
	for i := 1; i < n; i++ {
		if cur < 0 {
			cur = nums[i]
		} else {
			cur += nums[i]
		}
		if cur > best {
			best = cur
		}
	}

	if k == 0 {
		return best
	}

	// With swaps: we can bring positive elements from outside
	// into the subarray. Sort descending and consider adding
	// top positive elements.
  // Membuat slice untuk menyimpan hasil
	sorted := make([]int, n)
	copy(sorted, nums)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i] > sorted[j] })

	extra := 0
	count := 0
	for _, v := range sorted {
		if v > 0 && count < k {
			extra += v
			count++
		}
	}

	// Also compute best subarray sum after swapping in K elements
	// Try every window as candidate and augment with swaps
	for start := 0; start < n; start++ {
		for end := start; end < n && end-start+1 <= n; end++ {
			sum := 0
			for i := start; i <= end; i++ {
				sum += nums[i]
			}
			// We can swap up to k elements outside [start, end] into it
			// The best we can do is replace negative values inside
			// with positive values from outside
  // Membuat slice untuk menyimpan hasil
			inside := make([]int, 0)
  // Membuat slice untuk menyimpan hasil
			outside := make([]int, 0)
			for i := 0; i < n; i++ {
				if i >= start && i <= end {
					inside = append(inside, nums[i])
				} else {
					outside = append(outside, nums[i])
				}
			}
			sort.Ints(inside)
			sort.Slice(outside, func(i, j int) bool { return outside[i] > outside[j] })

			curSum := sum
			swaps := 0
			pi, po := 0, 0
			for pi < len(inside) && po < len(outside) && swaps < k {
				if inside[pi] < outside[po] {
					curSum = curSum - inside[pi] + outside[po]
					pi++
					po++
					swaps++
				} else {
					break
				}
			}
			if curSum > best {
				best = curSum
			}
		}
	}

	return best
}
```
