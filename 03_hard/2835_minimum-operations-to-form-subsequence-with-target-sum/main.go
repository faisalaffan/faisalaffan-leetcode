package main

// LeetCode #2835: Minimum Operations to Form Subsequence With Target Sum
// https://leetcode.com/problems/minimum-operations-to-form-subsequence-with-target-sum/
// Difficulty: Hard
//
// Given array nums containing powers of 2. Operation: split a number > 1 into
// two equal halves. Find minimum operations to form a subsequence summing to
// target. If total sum < target, return -1.
// Greedy: count occurrences of each power of 2, process target bits from low
// to high, splitting higher powers when needed. Carry extra counts upward.
// O(N + log target) time, O(log MAX) space.

import "fmt"

func minOperations(nums []int, target int) int {
	// Count total sum for early impossibility check
	totalSum := 0
	for _, v := range nums {
		totalSum += v
	}
	if totalSum < target {
		return -1
	}

	// Count occurrences of each power of 2
	cnt := make([]int, 31)
	for _, v := range nums {
		bit := 0
		for v > 1 {
			v >>= 1
			bit++
		}
		cnt[bit]++
	}

	ops := 0
	for i := 0; i < 31; i++ {
		if target&(1<<i) != 0 {
			if cnt[i] > 0 {
				cnt[i]--
			} else {
				// Find a higher power to split down to bit i
				j := i + 1
				for j < 31 && cnt[j] == 0 {
					j++
				}
				if j == 31 {
					return -1
				}
				// Split cnt[j] down to bit i
				for k := j; k > i; k-- {
					cnt[k]--
					cnt[k-1] += 2
					ops++
				}
				cnt[i]--
			}
		}
		// Carry remaining counts upward (pair up)
		cnt[i+1] += cnt[i] / 2
	}

	return ops
}

func main() {
	// Example 1: [1,2,8], target=7 => 1
	fmt.Println(minOperations([]int{1, 2, 8}, 7))

	// Example 2: [1,32,1,2], target=12 => 2
	fmt.Println(minOperations([]int{1, 32, 1, 2}, 12))

	// Example 3: [1,32,1], target=35 => -1
	fmt.Println(minOperations([]int{1, 32, 1}, 35))

	// Single element already matches
	fmt.Println(minOperations([]int{8}, 8))

	// Need to split down
	fmt.Println(minOperations([]int{16}, 8))

	// Multiple splits needed
	fmt.Println(minOperations([]int{32}, 10))

	// Already have the right combination
	fmt.Println(minOperations([]int{1, 2, 4, 8}, 15))

	// Duplicates available, combine to form target
	fmt.Println(minOperations([]int{1, 1, 1, 1, 1, 1, 1, 1}, 8))

	// Target is 0 (subset sum = 0 = pick nothing)
	fmt.Println(minOperations([]int{1, 2, 4}, 0))
}
