package main

// LeetCode #1723: Find Minimum Time to Finish All Jobs
// https://leetcode.com/problems/find-minimum-time-to-finish-all-jobs/
// Difficulty: Hard
// Strategy: DP over bitmask. dp[mask][k] = min possible max time
// assigning jobs in mask to k workers.

import (
	"fmt"
	"math"
)

func minimumTimeRequired(jobs []int, k int) int {
	n := len(jobs)

	// Precompute sum of each subset
	sum := make([]int, 1<<n)
	for mask := 1; mask < 1<<n; mask++ {
		lsb := mask & -mask
		// Find index of LSB
		idx := 0
		temp := lsb
		for temp > 1 {
			temp >>= 1
			idx++
		}
		sum[mask] = sum[mask^lsb] + jobs[idx]
	}

	// dp[mask] = minimum possible maximum time for this mask,
	// assigned to some number of workers
	// We'll do DP iteratively for each worker
	dp := make([]int, 1<<n)
	for mask := range dp {
		dp[mask] = math.MaxInt32
	}
	dp[0] = 0

	// For each worker, update dp
	for w := 0; w < k; w++ {
		next := make([]int, 1<<n)
		for mask := range next {
			next[mask] = math.MaxInt32
		}

		for mask := 0; mask < 1<<n; mask++ {
			if dp[mask] == math.MaxInt32 {
				continue
			}
			remaining := ((1 << n) - 1) ^ mask
			// Try all subsets of remaining for this worker
			sub := remaining
			for sub > 0 {
				candidate := max(dp[mask], sum[sub])
				newMask := mask | sub
				if candidate < next[newMask] {
					next[newMask] = candidate
				}
				sub = (sub - 1) & remaining
			}
		}
		dp = next
	}

	return dp[(1<<n)-1]
}

func main() {
	// Example 1: jobs=[3,2,3], k=2 -> 3
	jobs1 := []int{3, 2, 3}
	k1 := 2
	fmt.Printf("minimumTimeRequired(%v, %d) = %d (expected 3)\n", jobs1, k1, minimumTimeRequired(jobs1, k1))

	// Example 2: jobs=[1,2,4,7,8], k=2 -> 11
	jobs2 := []int{1, 2, 4, 7, 8}
	k2 := 2
	fmt.Printf("minimumTimeRequired(%v, %d) = %d (expected 11)\n", jobs2, k2, minimumTimeRequired(jobs2, k2))

	// Example 3: jobs=[11,2,7,4,8,10,3,1], k=3 -> 18
	jobs3 := []int{11, 2, 7, 4, 8, 10, 3, 1}
	k3 := 3
	fmt.Printf("minimumTimeRequired(%v, %d) = %d (expected 18)\n", jobs3, k3, minimumTimeRequired(jobs3, k3))

	// Single worker
	jobs4 := []int{5, 5, 5}
	k4 := 1
	fmt.Printf("minimumTimeRequired(%v, %d) = %d (expected 15)\n", jobs4, k4, minimumTimeRequired(jobs4, k4))

	// All jobs to each worker
	jobs5 := []int{1, 2, 3, 4, 5}
	k5 := 5
	fmt.Printf("minimumTimeRequired(%v, %d) = %d (expected 5)\n", jobs5, k5, minimumTimeRequired(jobs5, k5))
}
