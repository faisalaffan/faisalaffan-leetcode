package main

// LeetCode #2963: Count the Number of Good Partitions
// https://leetcode.com/problems/count-the-number-of-good-partitions/
// Difficulty: Hard
//
// Approach: Greedy merging of intervals + fast exponentiation.
// For each value, find its first and last occurrence (forming an interval).
// Merge overlapping intervals. Each merged segment is a "good" partition
// because no value appears across segments.
// Number of ways to partition k segments: 2^(k-1) mod (10^9+7).

import "fmt"

func numberOfGoodPartitions(nums []int) int {
	const mod = 1_000_000_007

	// Find last occurrence of each value
	last := make(map[int]int)
	for i, x := range nums {
		last[x] = i
	}

	// Merge overlapping intervals
	maxEnd := -1
	parts := 0
	for i, x := range nums {
		if last[x] > maxEnd {
			maxEnd = last[x]
		}
		if i == maxEnd {
			parts++
		}
	}

	// Compute 2^(parts-1) mod mod
	ans := 1
	for i := 1; i < parts; i++ {
		ans = (ans * 2) % mod
	}
	return ans
}

func main() {
	// Example: [1,2,3,4] -> 8 (4 segments, 2^3 = 8 ways)
	fmt.Println(numberOfGoodPartitions([]int{1, 2, 3, 4}))

	// All same value
	fmt.Println(numberOfGoodPartitions([]int{1, 1, 1, 1}))

	// Overlapping
	fmt.Println(numberOfGoodPartitions([]int{1, 2, 1, 3}))
}
