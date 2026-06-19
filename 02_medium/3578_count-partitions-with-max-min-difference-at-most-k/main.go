package main

// LeetCode #3578: Count Partitions With Max-Min Difference at Most K
// https://leetcode.com/problems/count-partitions-with-max-min-difference-at-most-k/
// Difficulty: Medium
// Complexity: O(2^n) time, O(n) space

import (
	"fmt"
	"math"
)

func main() {
	// Test case 1
	fmt.Println("Test 1:", CountPartitionsWithMaxMinDifferenceAtMostK([]int{1, 2, 3}, 1))
	// Test case 2
	fmt.Println("Test 2:", CountPartitionsWithMaxMinDifferenceAtMostK([]int{1, 1, 1}, 0))
	// Test case 3
	fmt.Println("Test 3:", CountPartitionsWithMaxMinDifferenceAtMostK([]int{5, 1, 2, 6}, 2))
}

func CountPartitionsWithMaxMinDifferenceAtMostK(nums []int, k int) int {
	n := len(nums)
	count := 0
	// Try all possible subsets
	for mask := 1; mask < (1<<n)-1; mask++ {
		minA, maxA := math.MaxInt32, 0
		minB, maxB := math.MaxInt32, 0
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				if nums[i] < minA {
					minA = nums[i]
				}
				if nums[i] > maxA {
					maxA = nums[i]
				}
			} else {
				if nums[i] < minB {
					minB = nums[i]
				}
				if nums[i] > maxB {
					maxB = nums[i]
				}
			}
		}
		if maxA-minA <= k && maxB-minB <= k {
			count++
		}
	}
	return count
}
