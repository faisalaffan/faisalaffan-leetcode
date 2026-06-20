# 1723 — Find Minimum Time To Finish All Jobs

## Deskripsi

**Soal:** [1723. Find Minimum Time To Finish All Jobs](https://leetcode.com/problems/find-minimum-time-to-finish-all-jobs/)

**Tingkat Kesulitan:** Sulit

**Kompleksitas Waktu:** —  
**Kompleksitas Ruang:** —

**Algoritma:** Dynamic Programming (DP), Bitmask (representasi himpunan dengan bit)

**Fungsi Solusi:** `func minimumTimeRequired(jobs []int, k int) int`

## Solusi Go

```go
package main

// LeetCode #1723: Find Minimum Time to Finish All Jobs
// https://leetcode.com/problems/find-minimum-time-to-finish-all-jobs/
// Difficulty: Hard
// Strategy: DP over bitmask. dp[mask] = min possible max time.
// Iterate w from 1 to k, each iteration adds one more worker.
// For each mask, try all subset splits: dp_prev[remaining] vs sum[sub].

import (
	"fmt"
	"math"
)

func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)

	// Precompute sum of each subset
  // Membuat slice untuk menyimpan hasil
	sum := make([]int, 1<<n)
	for mask := 1; mask < 1<<n; mask++ {
		lsb := mask & -mask
		idx := 0
		temp := lsb
		for temp > 1 {
			temp >>= 1
			idx++
		}
		sum[mask] = sum[mask^lsb] + jobs[idx]
	}

	// dp[mask] after w workers = min possible max time for jobs in mask
  // Membuat slice untuk menyimpan hasil
	dp := make([]int, 1<<n)
	for mask := range dp {
		dp[mask] = sum[mask] // 1 worker = sum of all jobs in mask
	}

	// For 2nd through kth worker
	for w := 2; w <= k; w++ {
  // Membuat slice untuk menyimpan hasil
		next := make([]int, 1<<n)
		for mask := range next {
			next[mask] = dp[mask] // start with previous value (one fewer worker)
		}

		for mask := 0; mask < 1<<n; mask++ {
			// Try splitting mask: subset 'sub' goes to the new worker,
			// remaining = mask ^ sub goes to the existing (w-1) workers
			sub := mask
			for sub > 0 {
				// Skip full mask (sub == mask): this would mean ALL jobs go to new worker,
				// which means the previous workers do nothing — handled by initialization.
				remaining := mask ^ sub
				if dp[remaining] != math.MaxInt32 {
					candidate := max(dp[remaining], sum[sub])
					if candidate < next[mask] {
						next[mask] = candidate
					}
				}
				sub = (sub - 1) & mask
			}
		}
		dp = next
	}

	return dp[(1<<n)-1]
}

func main() {
	// LeetCode Example 1: jobs=[3,2,3], k=3 -> 3
	// (each worker gets one job: max(3,2,3)=3)
	fmt.Printf("minimumTimeRequired(%v, %d) = %d (expected 3)\n",
		[]int{3, 2, 3}, 3, minimumTimeRequired([]int{3, 2, 3}, 3))

	// LeetCode Example 2: jobs=[1,2,4,7,8], k=2 -> 11
	// split: [8,2,1]=11, [7,4]=11
	fmt.Printf("minimumTimeRequired(%v, %d) = %d (expected 11)\n",
		[]int{1, 2, 4, 7, 8}, 2, minimumTimeRequired([]int{1, 2, 4, 7, 8}, 2))

	// LeetCode Example 3: jobs=[11,2,7,4,8,10,3,1], k=3
	fmt.Printf("minimumTimeRequired(%v, %d) = %d\n",
		[]int{11, 2, 7, 4, 8, 10, 3, 1}, 3, minimumTimeRequired([]int{11, 2, 7, 4, 8, 10, 3, 1}, 3))

	// Single worker
	fmt.Printf("minimumTimeRequired(%v, %d) = %d (expected 15)\n",
		[]int{5, 5, 5}, 1, minimumTimeRequired([]int{5, 5, 5}, 1))

	// All jobs to each worker (k == n)
	fmt.Printf("minimumTimeRequired(%v, %d) = %d (expected 5)\n",
		[]int{1, 2, 3, 4, 5}, 5, minimumTimeRequired([]int{1, 2, 3, 4, 5}, 5))
}
```
