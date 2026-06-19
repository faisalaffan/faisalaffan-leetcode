package main

// LeetCode #3862: Find the Smallest Balanced Index
// https://leetcode.com/problems/find-the-smallest-balanced-index/
// Difficulty: Medium
// Time: O(N) | Space: O(N)
// Approach: Compute prefix sums and suffix products, check equality at each index.

import "fmt"

func FindTheSmallestBalancedIndex(nums []int) int {
	n := len(nums)

	prefixSum := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
	}

	suffixProd := make([]int, n+1)
	suffixProd[n] = 1
	for i := n - 1; i >= 0; i-- {
		suffixProd[i] = suffixProd[i+1] * nums[i]
	}

	for i := 0; i < n; i++ {
		leftSum := prefixSum[i]
		rightProd := suffixProd[i+1]
		if leftSum == rightProd {
			return i
		}
	}

	return -1
}

func main() {
	// Example 1
	fmt.Println(FindTheSmallestBalancedIndex([]int{2, 1, 2})) // Expected: 1

	// Example 2
	fmt.Println(FindTheSmallestBalancedIndex([]int{2, 8, 2, 2, 5})) // Expected: 2

	// Example 3
	fmt.Println(FindTheSmallestBalancedIndex([]int{1})) // Expected: -1
}
