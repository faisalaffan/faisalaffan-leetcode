package main

// LeetCode #3231: Minimum Number of Increasing Subsequence to Be Removed
// https://leetcode.com/problems/minimum-number-of-increasing-subsequence-to-be-removed/
// Difficulty: Hard [Paid]
//
// By Dilworth's theorem, the minimum number of strictly increasing subsequences
// needed to partition the array equals the length of the longest non-increasing
// subsequence (LNDS). Compute LNDS via patience sorting (greedy + binary search)
// in O(n log n) time.

import "fmt"

func main() {
	// Example 1: [5,4,3,2,1] => 5 (each element alone)
	fmt.Println(minOperations([]int{5, 4, 3, 2, 1}))
	// Example 2: [1,2,3,4,5] => 1 (whole array)
	fmt.Println(minOperations([]int{1, 2, 3, 4, 5}))
	// Example 3: [5,3,1,4,2] => 3
	fmt.Println(minOperations([]int{5, 3, 1, 4, 2}))
	// Example 4: single element
	fmt.Println(minOperations([]int{1}))
	// Example 5: with duplicates
	fmt.Println(minOperations([]int{3, 3, 3}))
}

func minOperations(nums []int) int {
	// tails[k] = largest possible last element of a non-increasing subsequence of length k+1
	tails := make([]int, 0)

	for _, x := range nums {
		// Binary search for the first index where tails[i] < x
		lo, hi := 0, len(tails)
		for lo < hi {
			mid := (lo + hi) / 2
			if tails[mid] < x {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		if lo == len(tails) {
			tails = append(tails, x)
		} else {
			tails[lo] = x
		}
	}

	return len(tails)
}
