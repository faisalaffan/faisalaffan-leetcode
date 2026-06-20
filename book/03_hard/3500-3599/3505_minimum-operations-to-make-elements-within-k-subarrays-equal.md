# 3505 — Minimum Operations To Make Elements Within K Subarrays Equal

## Deskripsi

**Soal:** [3505. Minimum Operations To Make Elements Within K Subarrays Equal](https://leetcode.com/problems/minimum-operations-to-make-elements-within-k-subarrays-equal/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Sliding Window (jendela geser), Dynamic Programming (DP)

> **Ide Kunci:** Compute cost for each size-x window (optimal = sum of absolute

## Solusi Go

```go
package main

// LeetCode #3505: Minimum Operations to Make Elements Within K Subarrays Equal
// https://leetcode.com/problems/minimum-operations-to-make-elements-within-k-subarrays-equal/
// Difficulty: Hard
//
// Given nums, x, and k, find minimum operations to make k non-overlapping
// subarrays of length x all equal. Each operation increments or decrements
// any element by 1.
//
// Approach: Compute cost for each size-x window (optimal = sum of absolute
// differences to median). Then DP to select k non-overlapping windows.

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	// Example 1
	fmt.Println(minOperations([]int{1, 2, 3, 4, 5}, 2, 2))
	// Example 2
	fmt.Println(minOperations([]int{1, 2, 3, 4}, 2, 1))
	// Edge: k=1, single window
	fmt.Println(minOperations([]int{1, 10, 100}, 3, 1))
	// Edge: all same
	fmt.Println(minOperations([]int{5, 5, 5, 5, 5}, 2, 2))
}

func minOperations(nums []int, x int, k int) int64 {
	n := len(nums)
	if n < k*x {
		return 0
	}

	// Compute cost for each sliding window of size x
  // Membuat slice untuk menyimpan hasil
	costs := make([]int64, n-x+1)
	for i := 0; i <= n-x; i++ {
		// Extract window
  // Membuat slice untuk menyimpan hasil
		window := make([]int, x)
		copy(window, nums[i:i+x])
		sort.Ints(window)
		median := window[x/2]
		var total int64
		for _, v := range window {
			diff := v - median
			if diff < 0 {
				diff = -diff
			}
			total += int64(diff)
		}
		costs[i] = total
	}

	// DP: dp[j][i] = min cost with j subarrays using first i elements
  // Membuat slice 2D untuk DP/tabel
	dp := make([][]int64, k+1)
	for j := 0; j <= k; j++ {
		dp[j] = make([]int64, n+1)
		for i := 0; i <= n; i++ {
			dp[j][i] = math.MaxInt64 / 2
		}
	}
	for i := 0; i <= n; i++ {
		dp[0][i] = 0
	}

	for j := 1; j <= k; j++ {
		for i := 1; i <= n; i++ {
			// Skip element i-1
			dp[j][i] = dp[j][i-1]
			// Take subarray ending at i-1
			if i-x >= 0 {
				if dp[j-1][i-x]+costs[i-x] < dp[j][i] {
					dp[j][i] = dp[j-1][i-x] + costs[i-x]
				}
			}
		}
	}

	return dp[k][n]
}
```
